#!/bin/bash

cd $(dirname $0)/..

OLD="github.com/quic-go/quic-go"
NEW="github.com/c2FmZQ/quic-go-api"

for f in $(find . -name "*.go"); do
  if [[ $1 == "-undo" ]]; then
    sed -i -r -e 's:"'${NEW}':"'${OLD}':g' $f
  else
    sed -i -r -e 's:"'${OLD}':"'${NEW}':g' $f
  fi
done
