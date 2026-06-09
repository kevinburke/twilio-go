.PHONY: test lint race-test race-test-short fmt ci release force authors

# would be great to make the bash location portable but not sure how
SHELL = /bin/bash -o pipefail

STATICCHECK ?= staticcheck
version ?= minor

test: lint
	go test ./...

lint: fmt
	go vet ./...
	$(STATICCHECK) -version
	@output="$$( $(STATICCHECK) ./... 2>&1 )"; status="$$?"; \
	if [ -n "$$output" ]; then printf '%s\n' "$$output"; fi; \
	if [ "$$status" -ne 0 ]; then exit "$$status"; fi; \
	case "$$output" in \
		*"matched no packages"*) echo "$(STATICCHECK) matched no packages" >&2; exit 1 ;; \
	esac

race-test: lint
	go test -race ./...

race-test-short: lint
	go test -short -race ./...

fmt:
	go fmt ./...

ci: | $(DIFFER)
	# would love to run differ make authors here, but Github doesn't check out
	# the full history
	#
	# each differ invocation runs the command and fails the build if it leaves
	# a git diff, i.e. if the checked-in code is not already fixed/formatted.
	go run -v github.com/kevinburke/differ@latest go fix ./...
	go run -v github.com/kevinburke/differ@latest go fmt ./...
	go run -v github.com/kevinburke/differ@latest go run golang.org/x/tools/cmd/goimports@latest -w .
	$(MAKE) lint race-test-short

release: race-test-short
	go run -v github.com/kevinburke/differ@latest $(MAKE) authors
	go run -v github.com/kevinburke/differ@latest $(MAKE) fmt
	go run -v github.com/kevinburke/bump_version@latest --tag-prefix=v $(version) http.go

force: ;

AUTHORS.txt: .mailmap force
	go run -v github.com/kevinburke/write_mailmap@latest > AUTHORS.txt

authors: AUTHORS.txt
