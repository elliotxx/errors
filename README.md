<div align = "center">
<p>
    <img width="160" src="https://github.com/elliotxx/errors/blob/master/golang-logo.png?sanitize=true">
</p>
<h2>An errors package for web development</h2>
<a title="Go Reference" target="_blank" href="https://pkg.go.dev/github.com/elliotxx/errors"><img src="https://pkg.go.dev/badge/github.com/elliotxx/errors.svg"></a>
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/elliotxx/errors"><img src="https://goreportcard.com/badge/github.com/elliotxx/errors?style=flat-square"></a>
<a title="Coverage Status" target="_blank" href="https://coveralls.io/github/elliotxx/errors?branch=master"><img src="https://img.shields.io/coveralls/github/elliotxx/errors/master"></a>
<a title="Code Size" target="_blank" href="https://github.com/elliotxx/errors"><img src="https://img.shields.io/github/languages/code-size/elliotxx/errors.svg?style=flat-square"></a>
<br>
<a title="GitHub release" target="_blank" href="https://github.com/elliotxx/errors/releases"><img src="https://img.shields.io/github/release/elliotxx/errors.svg"></a>
<a title="License" target="_blank" href="https://github.com/elliotxx/errors/blob/master/LICENSE"><img src="https://img.shields.io/github/license/elliotxx/errors"></a>
<a title="GitHub Commits" target="_blank" href="https://github.com/elliotxx/errors/commits/master"><img src="https://img.shields.io/github/commit-activity/m/elliotxx/errors.svg?style=flat-square"></a>
<a title="Last Commit" target="_blank" href="https://github.com/elliotxx/errors/commits/master"><img src="https://img.shields.io/github/last-commit/elliotxx/errors.svg?style=flat-square&color=FF9900"></a>
</p>
</div>

`elliotxx/errors` is a golang package for web development, inspired by [pkg/errors](https://github.com/pkg/errors)!

## 📜 Language

[English](https://github.com/elliotxx/errors/blob/master/README.md) | [简体中文](https://github.com/elliotxx/errors/blob/master/README-zh.md)

## ⚡ Usage

```
go get -u github.com/elliotxx/errors
```

## ✨ Features
* More detailed error information, e.g. error code, message, cause, stack trace
* More powerful API
* Lightweight

## 📚 Examples
```go
package errcodes

import "github.com/elliotxx/errors"

var (
	Success                    = errors.NewErrorCode("00000", "success")
	NotFound                   = errors.NewErrorCode("A0100", "not found")
	AccessPermissionError      = errors.NewErrorCode("A0200", "abnormal access permission")
	AbnormalUserOperation      = errors.NewErrorCode("A0300", "abnormal user operation")
	InvalidParams              = errors.NewErrorCode("A0400", "invalid params")
	BlankRequiredParams        = errors.NewErrorCode("A0401", "required parameter is blank")
	MalformedParams            = errors.NewErrorCode("A0403", "parameter format mismatch")
	ServerError                = errors.NewErrorCode("A0500", "server error")
	TooManyRequests            = errors.NewErrorCode("A0501", "too many requests")
	ConcurrentExceedLimit      = errors.NewErrorCode("A0502", "the request parallel number exceeds the limit")
	WaitUserOperation          = errors.NewErrorCode("A0503", "please wait for user operation")
	RepeatedRequest            = errors.NewErrorCode("A0504", "repeated request")
)
```