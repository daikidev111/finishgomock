<div align="center">

# finishgomock

</div>

`finishgomock` is a golang linter that detects an unnecessary goMock Finish on gomock.Controller. 
As per the official document for gomock, it is not required to call ctrl.Finish() in a test method that uses goMock controller object anymore (only applied from go 1.14+). This linter is designed in a way that it flags and outputs an error when ctrl.Finish() is used in a test method, which achieves a reduction of unnecessary code.
For more details, please visit the official document: https://pkg.go.dev/github.com/golang/mock/gomock#NewController

### Installation

```shell
go get -u github.com/daikidev111/finishgomock/cmd/finishgomock
```

### Usage

```go
package main

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func fail(t *testing.T) {
	mock := gomock.NewController(t)
	mock.Finish()
}

func failSecond(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()
}

func pass(t *testing.T) {
	gomock.NewController(t)
}
```

```console
go vet -vettool=(which finishgomock) ./...

# output

```

## CI

### GitHub Actions

```yaml
- name: install finishgomock
  run: go install github.com/daikidev111/finishgomock/cmd/finishgomock
- name: exec finishgomock
  run: go vet -vettool=`which finishgomock` ./...
```

### CircleCI

```yaml
- run:
    name: install finishgomock
    command: go install github.com/daikidev111/finishgomock/cmd/finishgomock
- run:
    name: exec finishgomock
    command: go vet -vettool=`which finishgomock` ./...
```


