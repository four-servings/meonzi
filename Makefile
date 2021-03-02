ifndef binary
	binary=debug
endif

test:
	go test -v -cover -covermode=atomic ./...

mod-download:
	go mod download

generate:
	bash ./generate.sh

build:
	go build -o ${binary} .

unittest:
	go test -short  ./...

clean:
	if [ -f ${binary} ] ; then rm ${binary} ; fi

lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run \
		--exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		./...