RELEASES=bin/webgo-darwin-amd64 \
	 bin/webgo-linux-amd64 \
	 bin/webgo-windows-amd64.exe 

all: $(RELEASES)

bin/webgo-%: GOOS=$(firstword $(subst -, ,$*))
bin/webgo-%: GOARCH=$(subst .exe,,$(word 2,$(subst -, ,$*)))
bin/webgo-%: $(wildcard *.go)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build \
	     -ldflags "-X main.osarch=$(GOOS)/$(GOARCH) -s -w" \
	     -buildmode=exe \
	     -tags release \
	     -o $@

clean:
	rm -rf bin