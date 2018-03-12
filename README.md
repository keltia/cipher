# cipher — Old paper & pencil ciphers in Go.

[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/keltia/cipher) [![license](https://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/keltia/cipher/master/LICENSE) [![build](https://img.shields.io/travis/keltia/cipher.svg?style=flat)](https://travis-ci.org/keltia/cipher) [![Go Report Card](https://goreportcard.com/badge/github.com/keltia/cipher)](https://goreportcard.com/report/github.com/keltia/cipher)

`cipher` is a [Go](https://golang.org/) port of my [old-crypto](https://github.com/keltia/old-crypto) Ruby code. 

It features a simple CLI-based tool called `old-crypto` which serve both as a collection of use-cases for the library and an easy way to use it.

**Work in progress, still incomplete**

## Table of content

- [Features](#features)
- [Installation](#installation)
- [TODO](#todo)
- [Contributing](#contributing)

## Features

It currently implement a few of the Ruby code, namely:

- null
- Caesar (you can choose the shift number)
- Playfair
- Chaocipher
- Simple transposition (can be used with other ciphers as super-encipherement)
- Polybius square bigrammatic cipher (for ADFGVX = polybius + transposition)

It does not try to reinvent the wheel and implements the `cipher.Block` interface defined in the Go standard library (see `src/crypto/cipher/cipher.go`).

That means that all ciphers have `BlockSize(), Encrypt() & Decrypt()`.  You can create one with `NewCipher()` then use `Encrypt()`/`Decrypt`.  `BlockSize()` is of course implemented as well otherwise the interface would not be matched. 

## Installation

Like many Go-based tools, installation is very easy

    go get github.com/keltia/cipher/cmd/...

or

    git clone https://github.com/keltia/cipher
    make install

The library is fetched, compiled and installed in whichever directory is specified by `$GOPATH`.  The `old-crypto` binary will also be installed (on windows, this will be called `old-crypto.exe`).

To run the tests, you will need:

- `github.com/stretchr/assert`

NOTE: please use and test the Windows version (use `make windows`to generate it).  It should work but I lack resources to play much with it.

## Benchmarks

I tried to provide benchmarks for all ciphers (including key scheduling/expansion) and in some cases several implementations (and associated benchamarks).

## TODO

- more ciphers
- more tests (and better ones!)
- better display of results
- refactoring to reduce code duplication: always in progress
- even more tests

## Contributing

Please see CONTRIBUTING.md for some simple rules.