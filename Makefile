.PHONY: clean dist-clean

SOURCE := mustache.go
BINARY=mustache
VERSION=0.1
BUILD_TIME=`date +%FT%T%z`
DEPS = $(go list -f '{{range .Imports}}{{.}} {{end}}' ./...)

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCE)
	echo $(DEPS) | xargs -n1 go get -d
	go build -o ${BINARY} ${SOURCE}
	strip --strip-unneeded mustache

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

dist-clean:
	rm -r mustache lib
