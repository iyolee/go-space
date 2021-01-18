#!/bin/bash
if [[ $GOPATH =~ .*$PWD.* ]]
then
  echo "currnet dir is already in GOPATH"
else
  export GOPATH=$GOPATH:$PWD
  export GOBIN=$GOPATH/bin
echo "fininsh setting $PWD in GOPATH"
fi