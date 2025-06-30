#!/bin/bash

cd $(dirname $0)/../http3
for f in $(find . -name "*.go"); do
  sed -i -r -e 's:[*]quic[.]((Conn|Stream|SendStream|ReceiveStream|Transport)([,\) ]|$)):quicapi.\1:' $f
done
