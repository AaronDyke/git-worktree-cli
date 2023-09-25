build:
	go build -o ${GOPATH}/bin/wt

watch:
	ls **/*.go | entr go build -o ${GOPATH}/bin/wt