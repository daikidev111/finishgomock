package a

import (
	"testing"

	"github.com/golang/mock/gomock"
)

// Test case where Finish is called correctly
func TestFinishCalled(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish() // want "detected an unnecessary call to Finish on gomock.Controller"
}

// Test case where Finish is not called
func TestFinishNotCalled(t *testing.T) {
	gomock.NewController(t)
}

// Test case where Finish is called explicitly
func TestFinishCalledExplicitly(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCtrl.Finish() // want "detected an unnecessary call to Finish on gomock.Controller"
}

// Test case with multiple unnecessary Finish calls
func TestMultipleFinishCalls(t *testing.T) {
	mockCtrl1 := gomock.NewController(t)
	defer mockCtrl1.Finish() // want "detected an unnecessary call to Finish on gomock.Controller"

	mockCtrl2 := gomock.NewController(t)
	mockCtrl2.Finish() // want "detected an unnecessary call to Finish on gomock.Controller"

	mockCtrl3 := gomock.NewController(t)
	defer mockCtrl3.Finish() // want "detected an unnecessary call to Finish on gomock.Controller"
}
