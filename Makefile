.PHONY: clean dist-clean

SOURCE := mustache.go
BINARY=mustache
VERSION=0.2
BUILD_TIME=`date +%FT%T%z`
DEPS = $(go list -f '{{range .Imports}}{{.}} {{end}}' ./...)

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCE)
	$(MAKE) get-deps
	go build -o ${BINARY} ${SOURCE}
	strip --strip-unneeded mustache

get-deps:
	echo $(DEPS) | xargs -n1 go get -d
	go get -d "github.com/onsi/gomega"
	go get -d "github.com/onsi/ginkgo"

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

dist-clean:
	rm -r mustache lib
