#!/bin/bash
go install -a && $GOPATH/bin/decker-cli $@
