.PHONY: build release install test testacc vet salt-master-up salt-master-down fmt 

TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=$(shell jq .hostname package.json | xargs)
NAMESPACE=$(shell jq .namespace package.json | xargs)
NAME=$(shell jq .name package.json | xargs)
BINARY=terraform-provider-${NAME}
VERSION=$(shell jq .version package.json | xargs)
OS_ARCH=darwin_amd64

ifdef SALTSTACK_VERSION
export SALTSTACK_VER=${SALTSTACK_VERSION}
else
export SALTSTACK_VER=3004.2
endif

export SALTSTACK_HOST=localhost
export SALTSTACK_PORT=8000
export SALTSTACK_SCHEME=http
export SALTSTACK_EAUTH=sharedsecret
export SALTSTACK_USERNAME=username
export SALTSTACK_PASSWORD=password

default: install

build:
	go build -o ${BINARY}

release:
	goreleaser release --rm-dist --snapshot --skip-publish  --skip-sign

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test: fmtcheck
	go test $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

testacc: fmtcheck salt-master-up
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

salt-master-up:
	@echo "Starting up Salt-Master $(SALTSTACK_VER)..."
	docker-compose -f docker-compose/docker-compose.yaml up --remove-orphans -d
	docker-compose -f docker-compose/docker-compose.yaml logs
	scripts/wait-for-salt-master localhost:8000
	@echo "Salt-Master $(SALTSTACK_VER) is up and running."

salt-master-down:
	@echo "Shutting down Salt-Master $(SALTSTACK_VER)..."
	docker-compose -f docker-compose/docker-compose.yaml down --volumes
	@echo "Salt-Master $(SALTSTACK_VER) is down."

#! Development 
# The following make goals are only for local usage 
fmt:
	go fmt
	go fmt saltstack/*.go
