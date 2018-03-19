# Main Makefile for erc-eas
#
# Copyright 2015-2017 © by Ollivier Robert for the EEC
#

GOBIN=   ${GOPATH}/bin

BIN=	old-crypto
EXE=	${BIN}.exe

SRCS= cmd/old-crypto/main.go \
	  caesar/cipher.go crypto.go wheatstone/cipher.go \
      crypto_test.go playfair/cipher.go vic/cipher.go \
      null/cipher.go chaocipher/cipher.go \
      adfgvx/cipher.go nihilist/cipher.go straddling/cipher.go

SRCST= caesar/cipher_test.go chaocipher/cipher_test.go null/cipher_test.go \
	   playfair/cipher_test.go adfgvx/cipher_test.go straddling/cipher_test.go \
	   nihilist/cipher_test.go wheatstone/cipher_test.go \
	   vic/cipher_test.go

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
	/bin/rm -f ${BIN} ${EXE}

push:
	git push --all
	git push --tags
