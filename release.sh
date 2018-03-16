#!/bin/bash
export CI_BUILD_TAG=goxc-0.1.12
export CI_BUILD_REF=`git log --decorate | head -n 1 | awk -F' ' '{print $2}'`
export BUILD_STAMP=`date -u +%Y/%m/%d\_%H:%M:%S\_%Z` # -ldflags -X doesn't like spaces
export BUILD_NUMBER=`echo $BUILD_STAMP | sed -e 's/^..//g' | sed -e 's/....$//g' | sed -e 's/\//./g' | sed -e 's/_/-/g' | sed -e 's/://g'`

go build -installsuffix netgo -tags netgo -a -ldflags "-s -X github.com/openyard/goxc.version=${CI_BUILD_TAG} -X github.com/openyard/goxc.rev=${CI_BUILD_REF} -X github.com/openyard/goxc.built=${BUILD_STAMP}" -o goxc main/goxc.go
