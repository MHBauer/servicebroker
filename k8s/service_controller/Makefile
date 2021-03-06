.PHONY: all binary test-unit test clean server validate

EXE=service_controller
IMAGE=service-controller
SRC_FILES=$(shell find . -name \*.go -print | grep -v _test.go | grep -v \\./tests )
TEST_UNIT_DIRS=$(shell find . -name \*_test.go -exec dirname {} \; | grep -v \\./tests | sort -u)

all: validate binary test-unit test

#####################

binary: $(EXE)

$(EXE): $(SRC_FILES)
ifeq (,$(shell which go))
	@$(MAKE) --no-print-directory binary-docker
else
	@$(MAKE) --no-print-directory binary-local
endif
	docker build -t $(IMAGE):latest .

binary-docker:
	docker run -ti \
	  -w /src/github.com/servicebroker/servicebroker/k8s/service_controller \
	  -v ${CURDIR}:/src/github.com/servicebroker/servicebroker/k8s/service_controller \
	  -e GOPATH=/ \
	  golang \
	  sh -c "go get -d ./... && make binary-local"

binary-local:
	go get -d ./...
	go build -o $(EXE) .

#####################

test-unit: binary
	@echo "\n### Running the unit tests"
ifeq (,$(shell which go))
	@$(MAKE) --no-print-directory test-unit-docker
else
	@$(MAKE) --no-print-directory test-unit-local
endif

test-unit-docker:
	docker run -ti \
	  -w /src/github.com/servicebroker/servicebroker/k8s/service_controller \
	  -v ${CURDIR}:/src/github.com/servicebroker/servicebroker/k8s/service_controller \
	  -e GOPATH=/ \
	  golang \
	  sh -c "go get -d ./... && make --no-print-directory test-unit-local"

test-unit-local:
	@go test -v -x $(TEST_UNIT_DIRS)

#####################

test: binary
	@echo "\n### Running the integration tests"
ifeq (,$(shell which go))
	@$(MAKE) --no-print-directory test-docker
else
	@$(MAKE) --no-print-directory test-local
endif

test-docker:
	docker run -ti \
	  -w /src/github.com/servicebroker/servicebroker/k8s/service_controller \
	  -v ${CURDIR}:/src/github.com/servicebroker/servicebroker/k8s/service_controller \
	  -e GOPATH=/ \
	  --net host \
	  golang \
	  sh -c "go get -d ./... && make --no-print-directory test-local"

test-local: binary
	@go test -v ./tests

#####################

clean:
	-rm -f $(EXE)
	-docker rmi -f $(IMAGE) 2> /dev/null

server: binary
	docker run --net=host -ti -p 10000:10000 $(IMAGE)

broker:
	docker run --net=host -ti -p 9090:9090 gobroker


#####################

validate: 
ifeq (,$(shell which go))
	@$(MAKE) --no-print-directory validate-docker
else
	@$(MAKE) --no-print-directory validate-local
endif

validate-docker:
	docker run -ti \
	  -w /src/github.com/servicebroker/servicebroker/k8s/service_controller \
	  -v ${CURDIR}:/src/github.com/servicebroker/servicebroker/k8s/service_controller \
	  -e GOPATH=/ \
	  golang \
	  sh -c "go get -d ./... && make --no-print-directory validate-local"

validate-local:
	@if ! [ "`find . -name \*.go -exec gofmt -s -l {} \;`" = "" ]; then \
		echo Please fix the following gofmt issues: ; \
		find . -name \*.go -exec gofmt -s -d {} \; ; \
		exit 1 ; \
	fi
