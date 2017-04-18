name = terraform-provider-awsplus
package = github.com/bwits/$(name)
TERRAFORM_IMAGE = hashicorp/terraform:0.9.1
TERRAFORM_CMD = docker run -ti --rm -w /app -v $(TRAVIS_BUILD_DIR):/app -v $(HOME):/root $(TERRAFORM_IMAGE)

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
	GOOS=linux GOARCH=amd64 go build -o $(HOME)/release/$(name)-linux-amd64 $(package) ;\
	cp .terraformrc $(HOME)/.terraformrc ;\
	$(TERRAFORM_CMD) plan ;\
	}
