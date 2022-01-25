// copy from https://github.com/helm/helm/blob/main/pkg/chartutil/create_test.go

package action

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"helm.sh/helm/v3/pkg/chart/loader"
)

func TestCreate(t *testing.T) {
	tdir, err := ioutil.TempDir("", "helm-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tdir)

	c, err := Create("foo", tdir)
	if err != nil {
		t.Fatal(err)
	}

	dir := filepath.Join(tdir, "foo")

	mychart, err := loader.LoadDir(c)
	if err != nil {
		t.Fatalf("Failed to load newly created chart %q: %s", c, err)
	}

	if mychart.Name() != "foo" {
		t.Errorf("Expected name to be 'foo', got %q", mychart.Name())
	}

	for _, f := range []string{
		ChartfileName,
		DeploymentName,
		HelpersName,
		IgnorefileName,
		NotesName,
		ServiceAccountName,
		ServiceName,
		TemplatesDir,
		TemplatesTestsDir,
		TestConnectionName,
		ValuesfileName,
	} {
		if _, err := os.Stat(filepath.Join(dir, f)); err != nil {
			t.Errorf("Expected %s file: %s", f, err)
		}
	}
}

// TestCreate_Overwrite is a regression test for making sure that files are overwritten.
func TestCreateOverwrite(t *testing.T) {
	tdir, err := ioutil.TempDir("", "helm-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tdir)

	var errlog bytes.Buffer

	if _, err := Create("foo", tdir); err != nil {
		t.Fatal(err)
	}

	dir := filepath.Join(tdir, "foo")

	tplname := filepath.Join(dir, "templates/hpa.yaml")
	writeFile(tplname, []byte("FOO"))

	// Now re-run the create
	Stderr = &errlog
	if _, err := Create("foo", tdir); err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadFile(tplname)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) == "FOO" {
		t.Fatal("File that should have been modified was not.")
	}

	if errlog.Len() == 0 {
		t.Errorf("Expected warnings about overwriting files.")
	}
}

func TestValidateChartName(t *testing.T) {
	for name, shouldPass := range map[string]bool{
		"":                              false,
		"abcdefghijklmnopqrstuvwxyz-_.": true,
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ-_.": true,
		"$hello":                        false,
		"Hellô":                         false,
		"he%%o":                         false,
		"he\nllo":                       false,

		"abcdefghijklmnopqrstuvwxyz-_." +
			"abcdefghijklmnopqrstuvwxyz-_." +
			"abcdefghijklmnopqrstuvwxyz-_." +
			"abcdefghijklmnopqrstuvwxyz-_." +
			"abcdefghijklmnopqrstuvwxyz-_." +
			"abcdefghijklmnopqrstuvwxyz-_." +
			"abcdefghijklmnopqrstuvwxyz-_." +
			"abcdefghijklmnopqrstuvwxyz-_." +
			"abcdefghijklmnopqrstuvwxyz-_." +
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ-_.": false,
	} {
		if err := validateChartName(name); (err != nil) == shouldPass {
			t.Errorf("test for %q failed", name)
		}
	}
}
