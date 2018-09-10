#!/usr/bin/env bash
echo "Running go-bindata.
If go-bindata is not installed, install it via 'go get -u github.com/jteeuwen/go-bindata/...' "ba

echo "======== Srart ========"
echo "pwd: $(pwd)"
go-bindata -prefix "./templates"  -o "./bind/bindata.go" -pkg "bind" "./templates"

echo "======== End ========"
