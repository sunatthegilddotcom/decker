#!/bin/bash
rm -rf $GOPATH/bin/decker
go install && $GOPATH/bin/decker $@
