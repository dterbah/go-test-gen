## Go Test Generator (AKA go-test-gen)

This CLI enables to create automatically tests for your Golang projects.

![CI](https://github.com/dterbah/go-test-gen/actions/workflows/go-test.yml/badge.svg)
[![codecov](https://codecov.io/gh/dterbah/go-test-gen/branch/main/graph/badge.svg)](https://codecov.io/gh/dterbah/go-test-gen)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=dterbah_go-test-gen&metric=ncloc)](https://sonarcloud.io/summary/new_code?id=dterbah_go-test-gen)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=dterbah_go-test-gen&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=dterbah_go-test-gen)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=dterbah_go-test-gen&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=dterbah_go-test-gen)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=dterbah_go-test-gen&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=dterbah_go-test-gen)

<img src="assets/logo.png" width="100" />

## Installation

To install this CLI in your environment, you can use the following command :

```bash
    go get github.com/dterbah/go-test-gen
```

## Usages

The command to generate test of your project is the following one :

```bash
    go-test-gen generate --project="/path/to/golang/project"
    # or
    go-test-gen generate -p="/path/to/golang/project"
```
