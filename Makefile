# Main Makefile for erc-eas
#
# Copyright 2015-2017 © by Ollivier Robert for the EEC
#

GOBIN=   ${GOPATH}/bin

BIN=	old-crypto
EXE=	${BIN}.exe

SRCS= caesar/cipher.go		crypto.go			null/cipher_test.go \
      caesar/cipher_test.go		crypto_test.go			playfair/cipher.go \
      cmd/old-crypto/main.go		null/cipher.go			playfair/cipher_test.go

OPTS=	-ldflags="-s -w" -v

all: ${BIN} ${EXE}

${BIN}: ${SRCS}
	go build -o ${BIN} ${OPTS} cmd/${BIN}/main.go

${EXE}: ${SRCS}
	GOOS=windows go build -o ${EXE} ${OPTS}  cmd/${BIN}/main.go

pkg: ${BIN} ${EXE}
	-/bin/mkdir pkg
	tar cvf - ${BIN} ${XTRA} | xz -c >pkg/${BIN}.tar.xz
	zip pkg/${BIN}.zip ${XTRA} ${EXE}

test:
	go test ./...

bench:
	go test -bench=. -benchmem ./...

lint:
	gometalinter ./...

install: ${BIN}
	go install ${OPTS} .

clean:
	go clean -v
	/bin/rm -f pkg/${BIN}.tar.xz pkg/${BIN}.zip

push:
	git push --all
	git push --tags
