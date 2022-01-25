// copy from https://github.com/helm/helm/blob/main/pkg/chartutil/create.go
package action

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Stderr is an io.Writer to which error messages can be written
//
// In Helm 4, this will be replaced. It is needed in Helm 3 to preserve API backward
// compatibility.
var Stderr io.Writer = os.Stderr

// Create creates a new chart in a directory.
//
// Inside of dir, this will create a directory based on the name of
// chartfile.Name. It will then write the Chart.yaml into this directory and
// create the (empty) appropriate directories.
//
// The returned string will point to the newly created directory. It will be
// an absolute path, even if the provided base directory was relative.
//
// If dir does not exist, this will return an error.
// If Chart.yaml or any directories cannot be created, this will return an
// error. In such a case, this will attempt to clean up by removing the
// new chart directory.
func Create(name, dir string) (string, error) {

	// Sanity-check the name of a chart so user doesn't create one that causes problems.
	if err := validateChartName(name); err != nil {
		return "", err
	}

	path, err := filepath.Abs(dir)
	if err != nil {
		return path, err
	}

	if fi, err := os.Stat(path); err != nil {
		return path, err
	} else if !fi.IsDir() {
		return path, errors.Errorf("no such directory %s", path)
	}

	cdir := filepath.Join(path, name)
	if fi, err := os.Stat(cdir); err == nil && !fi.IsDir() {
		return cdir, errors.Errorf("file %s already exists and is not a directory", cdir)
	}

	files := []struct {
		path    string
		content []byte
	}{
		{
			// Chart.yaml
			path:    filepath.Join(cdir, ChartfileName),
			content: []byte(fmt.Sprintf(defaultChartfile, name)),
		},
		{
			// values.yaml
			path:    filepath.Join(cdir, ValuesfileName),
			content: []byte(fmt.Sprintf(defaultValues, name)),
		},
		{
			// .helmignore
			path:    filepath.Join(cdir, IgnorefileName),
			content: []byte(defaultIgnore),
		},
		{
			// ingress.yaml
			path:    filepath.Join(cdir, IngressFileName),
			content: transform(defaultIngress, name),
		},
		{
			// deployment.yaml
			path:    filepath.Join(cdir, DeploymentName),
			content: transform(defaultDeployment, name),
		},
		{
			// service.yaml
			path:    filepath.Join(cdir, ServiceName),
			content: transform(defaultService, name),
		},
		{
			// serviceaccount.yaml
			path:    filepath.Join(cdir, ServiceAccountName),
			content: transform(defaultServiceAccount, name),
		},
		{
			// hpa.yaml
			path:    filepath.Join(cdir, HorizontalPodAutoscalerName),
			content: transform(defaultHorizontalPodAutoscaler, name),
		},
		{
			// NOTES.txt
			path:    filepath.Join(cdir, NotesName),
			content: transform(defaultNotes, name),
		},
		{
			// _helpers.tpl
			path:    filepath.Join(cdir, HelpersName),
			content: transform(defaultHelpers, name),
		},
		{
			// test-connection.yaml
			path:    filepath.Join(cdir, TestConnectionName),
			content: transform(defaultTestConnection, name),
		},
	}

	for _, file := range files {
		if _, err := os.Stat(file.path); err == nil {
			// There is no handle to a preferred output stream here.
			fmt.Fprintf(Stderr, "WARNING: File %q already exists. Overwriting.\n", file.path)
		}
		if err := writeFile(file.path, file.content); err != nil {
			return cdir, err
		}
	}
	// Need to add the ChartsDir explicitly as it does not contain any file OOTB
	if err := os.MkdirAll(filepath.Join(cdir, ChartsDir), 0755); err != nil {
		return cdir, err
	}
	return cdir, nil
}

// transform performs a string replacement of the specified source for
// a given key with the replacement string
func transform(src, replacement string) []byte {
	return []byte(strings.ReplaceAll(src, "<CHARTNAME>", replacement))
}

func writeFile(name string, content []byte) error {
	if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(name, content, 0644)
}

func validateChartName(name string) error {
	if name == "" || len(name) > maxChartNameLength {
		return fmt.Errorf("chart name must be between 1 and %d characters", maxChartNameLength)
	}
	if !chartName.MatchString(name) {
		return fmt.Errorf("chart name must match the regular expression %q", chartName.String())
	}
	return nil
}
