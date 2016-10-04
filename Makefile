name = terraform-provider-awsplus
package = github.com/bwits/$(name)

.PHONY: release
release:
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/$(name)-linux-amd64 $(package)
	GOOS=linux GOARCH=386 go build -o release/$(name)-linux-386 $(package)
	GOOS=linux GOARCH=arm go build -o release/$(name)-linux-arm $(package)
	GOOS=darwin GOARCH=amd64 go build -o release/$(name)-darwin-amd64 $(package)
