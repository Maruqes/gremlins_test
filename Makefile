GREMLINS_VERSION := v0.6.0
GREMLINS := $(CURDIR)/bin/gremlins

.PHONY: run test install-gremlins mutation

run:
	go run .

test:
	go test -v -cover ./...

install-gremlins:
	GOBIN="$(CURDIR)/bin" go install github.com/go-gremlins/gremlins/cmd/gremlins@$(GREMLINS_VERSION)

$(GREMLINS):
	$(MAKE) install-gremlins

mutation: $(GREMLINS)
	$(GREMLINS) unleash --workers 2 --timeout-coefficient 50 --threshold-efficacy 100
