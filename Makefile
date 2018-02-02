.PHONY: .FORCE

GO=go
GOOS=linux
GOPATH := $(GOPATH)

PROGS = learnToWork 

SRCDIR = $(GOPATH)/src

VERSION=v.1.0.0
LDFLAGS=' -w -s -X main._VERSION_="$(VERSION)"'


$(PROGS):
	$(GO) install -ldflags $(LDFLAGS) gs
	$(GO) install -ldflags $(LDFLAGS) gate


all: $(PROGS)

gs:
	 $(GO) install -ldflags $(LDFLAGS)  gs

gate:
	 $(GO) install -ldflags $(LDFLAGS)  gate

clean:
	rm -rf pkg 

debug:
	 $(GO) install -ldflags $(LDFLAGS) -gcflags "-N -l" $(PROGS)

pb:
	protoc -I $(SRCDIR)/protocol   --go_out=plugins=grpc:$(SRCDIR)/protocol/  const.proto gate.proto gs.proto

fmt:
	 $(GO) fmt $(SRCDIR)/...
