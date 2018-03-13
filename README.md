### build templates
```sh
go-bindata -o templates.go -pkg goxc templates
```

### build executable
```sh
export CI_BUILD_REF_NAME=
export CI_BUILD_REF=`git log --decorate | head -n 1 | awk -F' ' '{print $2}'`
export BUILD_STAMP=`date -u +%Y/%m/%d\_%H:%M:%S\_%Z` # -ldflags -X doesn't like spaces
export BUILD_NUMBER=`echo $BUILD_STAMP | sed -e 's/^..//g' | sed -e 's/....$//g' | sed -e 's/\//./g' | sed -e 's/_/-/g' | sed -e 's/://g'`

go build -installsuffix netgo -tags netgo -a -ldflags "-s -X github.com/openyard/goxc.version=${CI_BUILD_REF_NAME}-${BUILD_NUMBER} -X github.com/openyard/goxc.rev=${CI_BUILD_REF} -X github.com/openyard/goxc.built=${BUILD_STAMP}" -o goxc ../main
```
