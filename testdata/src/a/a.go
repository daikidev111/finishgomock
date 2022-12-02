package a

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func okFirst(t *testing.T) {
	mock := gomock.NewController(t)
	mock.Finish() // want "identifier is GoMock Finish"
}

func okSecond(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish() // want "identifier is GoMock Finish"
}

func failFirst(t *testing.T) {
	gomock.NewController(t)
}
