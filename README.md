# demo_gin-csrf 

[![Build Status](https://travis-ci.org/mikeee/demo_gin-csrf.svg?branch=master)](https://travis-ci.org/mikeee/demo_gin-csrf)
[![Maintainability](https://api.codeclimate.com/v1/badges/595a3c21d8e6d64a78b2/maintainability)](https://codeclimate.com/github/mikeee/demo_gin-csrf/maintainability)


Quick test of an example for a gin middleware that helps prevent csrf attacks

## tl;dr

This creates a simple gin-powered server (running on port 8000) that implements utrack/gin-csrf with gin-contrib/sessions and has a complete test for it.

## Usage

Feel free to run the main package and play around with it.

```bash
go run main.go
```

Testing:

```bash
go test -v
```

## Links

- [github.com/utrack/gin-csrf](https://github.com/utrack/gin-csrf)