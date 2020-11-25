.PHONY: test vet release

# would be great to make the bash location portable but not sure how
SHELL = /bin/bash -o pipefail

DIFFER := $(GOPATH)/bin/differ
WRITE_MAILMAP := $(GOPATH)/bin/write_mailmap
BUMP_VERSION := $(GOPATH)/bin/bump_version
STATICCHECK := $(GOPATH)/bin/staticcheck

test: lint
	go test ./...

$(BUMP_VERSION):
	go get github.com/kevinburke/bump_version

$(DIFFER):
	go get github.com/kevinburke/differ

$(STATICCHECK):
	go get honnef.co/go/tools/cmd/staticcheck

$(WRITE_MAILMAP):
	go get github.com/kevinburke/write_mailmap

lint: fmt | $(STATICCHECK)
	go vet ./...
	$(STATICCHECK) ./...

race-test: lint
	go test -race ./...

race-test-short: lint
	go test -short -race ./...

fmt:
	go fmt ./...

ci: | $(DIFFER)
	# would love to run differ make authors here, but Github doesn't check out
	# the full history
	$(DIFFER) $(MAKE) fmt
	$(MAKE) lint race-test-short

release: race-test | $(DIFFER) $(BUMP_VERSION)
	$(DIFFER) $(MAKE) authors
	$(DIFFER) $(MAKE) fmt
	$(BUMP_VERSION) minor http.go

force: ;

AUTHORS.txt: .mailmap force | $(WRITE_MAILMAP)
	$(WRITE_MAILMAP) > AUTHORS.txt

authors: AUTHORS.txt
