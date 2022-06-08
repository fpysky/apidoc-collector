#!/bin/sh

# shellcheck disable=SC2164
cd /opt/from-host/project/apidoc-generator

./main

apidoc -c ./app-doc/apidoc.json -i ./app-doc -o ./app-apidoc
rm -r /opt/from-host/apidoc/app-apidoc
cp -r ./app-apidoc /opt/from-host/apidoc/app-apidoc

apidoc -c ./admin-doc/apidoc.json -i ./admin-doc -o ./admin-apidoc
rm -r /opt/from-host/apidoc/admin-apidoc
cp -r ./admin-apidoc /opt/from-host/apidoc/admin-apidoc

apidoc -c ./mp-doc/apidoc.json -i ./mp-doc -o ./mp-apidoc
rm -r /opt/from-host/apidoc/mp-apidoc
cp -r ./mp-apidoc /opt/from-host/apidoc/mp-apidoc