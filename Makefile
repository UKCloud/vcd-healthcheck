NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m
DEPS = $(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
UNAME := $(shell uname -s)
ifeq ($(UNAME),Darwin)
ECHO=echo
else
ECHO=/bin/echo -e
endif

OS = linux windows
ARCH = amd64

PKG_windows = ZIP
PKG_linux = TAR
ZIP_CMD = zip
ZIP_FILE_SUFFIX = zip
TAR_CMD = tar zcvf
TAR_FILE_SUFFIX = tar.gz

GIT_TAG := $(shell git describe --tags)

all: test bin dist

get-deps:
	@$(ECHO) "$(OK_COLOR)==> Fetching dependencies$(NO_COLOR)"
	go get -v ./...
	go get -u github.com/golang/lint/golint
	go get -u gopkg.in/check.v1
	#go get -u golang.org/x/tools/cmd/vet

test: get-deps
	@$(ECHO) "$(OK_COLOR)==> Testing$(NO_COLOR)"
	#go vet ./...
	$(GOPATH)/bin/golint ./...
	go test -v -check.v ./...

bin: get-deps
	@$(ECHO) "$(OK_COLOR)==> Building$(NO_COLOR)"
	go build

clean:
	@rm -rf dist/ vcd-healthcheck vcd-healthcheck.exe

format:
	go fmt ./...

dist:	$(OS)

$(OS): distdir	
	@$(ECHO) "$(OK_COLOR)==> Building $@ Packages for $(GIT_TAG) ...$(NO_COLOR)"
	$(eval BASEDIR = "$(PWD)")

	$(foreach GOARCH,$(ARCH), \
	    $(eval DISTARCH = "$@-$(GOARCH)") \
            $(eval BINARY = "vcd-healthcheck$(if $(findstring windows,$@),".exe","")") \
	    $(eval PKG = $(PKG_$@)) \
	    $(eval PKGCMD = $($(PKG)_CMD)) \
	    $(eval PKGSUFFIX = $($(PKG)_FILE_SUFFIX)) \
	    \
            @$(ECHO) "$(OK_COLOR)Building $(BINARY) for $(DISTARCH)$(NO_COLOR)"; \
            GOOS=$@ GOARCH=$(GOARCH) go build -o $(BINARY) -ldflags "-X main.VERSION=$(GIT_TAG)"; \
	    $(PKGCMD) dist/vcd-healthcheck.$(DISTARCH).$(PKGSUFFIX) $(BINARY) \
	)

distdir:
	@mkdir -p dist


.PHONY: all clean deps format test updatedeps
