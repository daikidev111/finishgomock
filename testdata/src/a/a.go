package a

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func okFirst(t *testing.T) {
	mock := gomock.NewController(t)
	mock.Finish() // want "detected an unnecessary call to Finish on gomock.Controllers"
}

func okSecond(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish() // want "detected an unnecessary call to Finish on gomock.Controllers"
}

func failFirst(t *testing.T) {
	gomock.NewController(t)
}
