package a

import (
	"testing"

	"github.com/golang/mock/gomock"
)

// func f(t *testing.T) {
// 	mockF := gomock.NewController(t)
// 	a := "hi"
// 	print(a)
// 	mockF.Finish()
// }

func ok(t *testing.T) {
	mockOk := gomock.NewController(t)
	defer mockOk.Finish() // want "identifier is GoMock Finish"
}