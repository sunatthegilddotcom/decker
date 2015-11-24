#!/bin/bash
rm -rf $GOPATH/bin/decker-cli
go install && $GOPATH/bin/decker-cli $@
