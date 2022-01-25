#!/bin/sh -e

# copied from https://github.com/chartmuseum/helm-push/blob/main/scripts/install_plugin.sh

if [ -n "${HELM_LINTER_PLUGIN_NO_INSTALL_HOOK}" ]; then
    echo "Development mode: not downloading versioned release."
    exit 0
fi

version="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"
echo "Downloading and installing bcs-create ${version} ..."

url=""
if [ "$(uname)" = "Darwin" ]; then
    url="https://github.com/dellkeji/bcs-create-chart/releases/download/v${version}/bcs-create_${version}_darwin_amd64.tar.gz"
elif [ "$(uname)" = "Linux" ] ; then
    url="https://github.com/dellkeji/bcs-create-chart/releases/download/v${version}/bcs-create_${version}_linux_amd64.tar.gz"
else
    url="https://github.com/dellkeji/bcs-create-chart/releases/download/v${version}/bcs-create_${version}_windows_amd64.tar.gz"
fi

echo "$url"

mkdir -p "bin"
mkdir -p "releases/${version}"

# Download with curl if possible.
if [ -x "$(which curl 2>/dev/null)" ]; then
    curl -sSL "${url}" -o "releases/${version}.tar.gz"
else
    wget -q "${url}" -O "releases/${version}.tar.gz"
fi
tar xzf "releases/${version}.tar.gz" -C "releases/${version}"
mv "releases/${version}/bcs-create" "bin/bcs-create"
