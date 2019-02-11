#!/bin/sh

rm -f release.txt

for OS in windows linux darwin; do
  mkdir -p ${OS}
  GOOS=${OS} GOARCH=amd64 go build -o ${OS}/kubectl-match_name
  if [ "${OS}" == "windows" ]; then
    mv windows/kubectl-match_name windows/kubectl-match_name.exe
  fi
  zip kubectl-match_name-${OS}-amd64.zip ${OS}/kubectl-match_name*
  rm -rf ${OS}/
  shasum -a 256 kubectl-match_name-${OS}-amd64.zip >>release.txt
done
