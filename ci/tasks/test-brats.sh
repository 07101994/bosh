#!/usr/bin/env bash

set -eu

script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
src_dir="${script_dir}/../../.."

"${src_dir}/bosh-src/ci/docker/main-bosh-docker/start-bosh.sh"

source /tmp/local-bosh/director/env

bosh int /tmp/local-bosh/director/creds.yml --path /jumpbox_ssh/private_key > /tmp/jumpbox_ssh_key.pem
chmod 400 /tmp/jumpbox_ssh_key.pem

export BOSH_SSH_PRIVATE_KEY_PATH="/tmp/jumpbox_ssh_key.pem"
export BOSH_BINARY_PATH=$(which bosh)
export BOSH_RELEASE="${PWD}/bosh-src/src/spec/assets/dummy-release.tgz"
export BOSH_DIRECTOR_IP="10.245.0.3"
export BOSH_DIRECTOR_RELEASE_PATH="${PWD}/bosh-release"
export DNS_RELEASE_PATH="$(realpath $(find ${PWD}/bosh-dns-release -maxdepth 1 -path '*.tgz'))"
export CANDIDATE_STEMCELL_TARBALL_PATH="$(realpath ${src_dir}/stemcell/*.tgz)"
export BOSH_DNS_ADDON_OPS_FILE_PATH="/usr/local/bosh-deployment/experimental/dns-addon-with-api-certificates.yml"


mkdir -p bbr-binary
export BBR_VERSION=1.2.2
curl -L -o bbr-binary/bbr https://s3.amazonaws.com/bosh-dependencies/bbr-$BBR_VERSION

export BBR_SHA256=829160a61a44629a2626b578668777074c7badd75a9b5dab536defdbdd84b17a
export BBR_BINARY_PATH="${PWD}/bbr-binary/bbr"

echo "${BBR_SHA256} ${BBR_BINARY_PATH}" | sha256sum -c -

chmod +x ${BBR_BINARY_PATH}

pushd bosh-src/src/go
  export GOPATH=$(pwd)
  export PATH="${GOPATH}/bin":$PATH

  pushd src/github.com/cloudfoundry/bosh-release-acceptance-tests
    go install ./vendor/github.com/onsi/ginkgo/ginkgo
    ginkgo -v -r -race -randomizeSuites -randomizeAllSpecs .
  popd
popd
