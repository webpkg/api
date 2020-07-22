RELEASES=bin/api-darwin-amd64 \
	 bin/api-linux-amd64

all: $(RELEASES)

bin/api-%: GOOS=$(firstword $(subst -, ,$*))
bin/api-%: GOARCH=$(subst .exe,,$(word 2,$(subst -, ,$*)))
bin/api-%: $(wildcard *.go)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build \
	     -ldflags "-X main.osarch=$(GOOS)/$(GOARCH) -s -w" \
	     -buildmode=exe \
	     -tags release \
	     -o $@

clean:
	rm -rf bin
