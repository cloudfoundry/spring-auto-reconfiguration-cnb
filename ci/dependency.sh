#!/usr/bin/env bash

set -euo pipefail

if [[ -d $PWD/go-module-cache && ! -d ${GOPATH}/pkg/mod ]]; then
  mkdir -p ${GOPATH}/pkg
  ln -s $PWD/go-module-cache ${GOPATH}/pkg/mod
fi

commit() {
  git commit -a -m "Dependency Upgrade: $1 $2" || true
}

version() {
  local PATTERN="([0-9]+)\.([0-9]+)\.([0-9]+)-(.*)"

  for VERSION in $(cat $1); do
      if [[ ${VERSION} =~ ${PATTERN} ]]; then
        echo "${BASH_REMATCH[1]}.${BASH_REMATCH[2]}.${BASH_REMATCH[3]}"
        return
      else
        >2 echo "version is not semver"
        exit 1
      fi
    done
}

cd "$(dirname "${BASH_SOURCE[0]}")/.."
git config --local user.name "$GIT_USER_NAME"
git config --local user.email ${GIT_USER_EMAIL}

go build -ldflags='-s -w' -o bin/dependency github.com/cloudfoundry/libcfbuildpack/dependency

bin/dependency auto-reconfiguration "[\d]+\.[\d]+\.[\d]+" $(version ../auto-reconfiguration/version) $(cat ../auto-reconfiguration/uri)  $(cat ../auto-reconfiguration/sha256)
commit auto-reconfiguration $(version ../auto-reconfiguration/version)
