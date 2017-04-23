name = terraform-provider-awsplus
package = github.com/bwits/$(name)
modules = $(shell find . -type f -name "*.tf" -exec dirname {} \;|sort -u)

.PHONY: release

get-deps:
	go get -v github.com/bwits/$(name)

release: get-deps
	mkdir -p $(HOME)/release
	GOOS=linux GOARCH=amd64 go build -o $(HOME)/release/$(name)-linux-amd64 $(package)
	GOOS=linux GOARCH=386 go build -o $(HOME)/release/$(name)-linux-386 $(package)
	GOOS=linux GOARCH=arm go build -o $(HOME)/release/$(name)-linux-arm $(package)
	GOOS=darwin GOARCH=amd64 go build -o $(HOME)/release/$(name)-darwin-amd64 $(package)

plan: release
	{ \
	cp .terraformrc $(HOME)/.terraformrc ;\
	terraform plan ;\
	}

test:
	@for m in $(modules); do (terraform validate "$$m" && echo "√ $$m") || exit 1 ; done

fmt:
	@if [ `terraform fmt | wc -c` -ne 0 ]; then echo "terraform files need be formatted"; exit 1; fi
