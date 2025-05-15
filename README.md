<div align="center">


# Learn Cryptography in Go

![GitHub top language](https://img.shields.io/github/languages/top/alerone/passly?color=%2377CDFF)
![GitHub last commit](https://img.shields.io/github/last-commit/alerone/passly?color=%23bc0bbf)
![GitHub Created At](https://img.shields.io/github/created-at/alerone/passly?color=%230dba69)
![GitHub repo size](https://img.shields.io/github/repo-size/alerone/passly?color=%23390385)

<br>

<img src="https://github.com/user-attachments/assets/463387cc-2c8e-4264-8860-451d1e790648" alt="golang y http" width="250" height="250"/>

</div>


This repository contains code examples and implementations from the [Learn Cryptography in Go](https://www.boot.dev/courses/learn-cryptography-golang) course on Boot.dev. Throughout this course, I've explored fundamental and advanced cryptographic techniques implemented in Go.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Running Tests](#running-tests)

## Features

- **Alphabetic Ciphers**: Implementations of Caesar ciphers.
- **Block Ciphers**: Use of libraries for DES and AES ciphers. A custom implementation of DES.
- **Stream Ciphers**: Examples of OTP and XOR-based stream ciphers.
- **Hash Functions**: Toy hash and cryptographic hash functions.
- **Message Authentication**: Checksums, MAC, and HMAC implementations.
- **Asymmetric Cryptography**: RSA and ECDSA for key generation, signing, and verification. Custom implementation of RSA.
- **Key Derivation**: Password-based KDF using bcrypt.
- **Initialization Vectors & Nonces**: Secure IV/nonce generation for ciphers.
- **JWT**: Synchronous and asynchronous JWT creation and validation.

## Prerequisites

- Go 1.20+ installed
- `go.mod` and `go.sum` are included for dependency management

## Installation

```sh
git clone https://github.com/yourusername/cryptography-go-course
cd cryptography-go-course
go mod download
```

## Usage

Command-line tools are available under the `cmd` directory:

- **hash**: Generate toy hashes for input data.
- **streamCipher**: Encrypt/decrypt using stream ciphers.

```sh
# Example: Run the stream cipher tool
go run ./cmd/streamCipher/
```

## Project Structure

```
.
├── alphabet.go
├── bruteforce.go
├── cipher.go
├── cmd
│   ├── blockCipher
│   ├── hash/hash.go
│   └── streamCipher/main.go
├── decoding.go
├── encoding.go
├── go.mod
├── go.sum
├── internal
│   ├── aes
│   ├── blockCipher
│   ├── caesarCipher
│   ├── checksum
│   ├── des
│   ├── ecc
│   ├── ecdsa
│   ├── feistel
│   ├── hash
│   ├── iv
│   ├── kdf
│   ├── otp
│   ├── roundKey
│   ├── rsa
│   ├── sBox
│   └── streamCipher
├── main.go
├── main_test.go
├── randomKey.go
├── xor.go
└── xor_test.go
```

## Running Tests

Run all unit tests:

```sh
go test ./...
```
