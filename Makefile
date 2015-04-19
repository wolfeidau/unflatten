GOPATH := ${PWD}/.gopath
STAGING_PATH = ${PWD}/.gopath/src/github.com/wolfeidau
DEPS = $(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)

all: deps test bench

deps:
	@mkdir -p ${STAGING_PATH}
	@ln -s ${PWD} ${STAGING_PATH} || true
	@cd ${STAGING_PATH}/unflatten
	go get -d -v ./...

test: deps
	@cd ${STAGING_PATH}/unflatten
	go test -timeout=3s -v ./...

bench: deps
	@cd ${STAGING_PATH}/unflatten
	go test --bench .

clean:
	rm -rf ${PWD}/.gopath || true

.PHONY: all deps test bench clean
