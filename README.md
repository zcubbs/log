# Log

A versatile Go logging wrapper that allows seamless integration and toggling between various logging libraries such as Logrus, Zap, Charmlog and the standard log package. It also supports different logging formats like JSON and text, making it a flexible choice for various application needs.

[![tag](https://img.shields.io/github/tag/zcubbs/logwrapper)](https://github.com/zcubbs/logwrapper/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-%23007d9c)
[![GoDoc](https://godoc.org/github.com/zcubbs/logwrapper?status.svg)](https://pkg.go.dev/github.com/zcubbs/logwrapper)
[![Lint](https://github.com/zcubbs/logwrapper/actions/workflows/lint.yaml/badge.svg)](https://github.com/zcubbs/logwrapper/actions/workflows/lint.yaml)
[![Scan](https://github.com/zcubbs/logwrapper/actions/workflows/scan.yaml/badge.svg?branch=main)](https://github.com/zcubbs/logwrapper/actions/workflows/scan.yaml)
![Build Status](https://github.com/zcubbs/logwrapper/actions/workflows/test.yaml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/zcubbs/logwrapper)](https://goreportcard.com/report/github.com/zcubbs/logwrapper)
[![Contributors](https://img.shields.io/github/contributors/zcubbs/logwrapper)](https://github.com/zcubbs/logwrapper/graphs/contributors)
[![License](https://img.shields.io/github/license/zcubbs/logwrapper.svg)](./LICENSE)

## Installation

To install the LogWrapper package, run the following command:

```bash
go get github.com/zcubbs/log
```

## Usage

Here is a simple example demonstrating the usage of LogWrapper with different logging types and formats:

```go
package main

import (
    "github.com/zcubbs/log"
    "github.com/zcubbs/log/structuredlogger"
)

func main() {
    log.Info("This is an info message in JSON format")

    log.SetFormat(structuredlogger.TextFormat)
    log.Info("This is an info message in text format")
}
```

For a more comprehensive example, please refer to [`example.go`](./examples/example.go) in this repository.

## Supported Loggers

- [Logrus](https://github.com/sirupsen/logrus)
- [Zap](https://github.com/uber-go/zap)
- [Charmlog](https://github.com/charmbracelet/log)
- Standard log package

## Supported Formats

- JSON
- Text (Human-Readable)

## Contributing

We welcome contributions from the community! Feel free to open issues for bug reports or feature requests, and submit pull requests for improvements to the codebase.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
