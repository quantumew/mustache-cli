.PHONY: clean dist-clean

SOURCE := mustache.go
BINARY=mustache
VERSION=0.2
BUILD_TIME=`date +%FT%T%z`

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCE)
	$(MAKE) get-deps
	go build -o ${BINARY} ${SOURCE}
	strip mustache

get-deps:
	go get -d
	go get -d "github.com/onsi/gomega"
	go get -d "github.com/onsi/ginkgo"

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

dist-clean:
	rm -r mustache lib
