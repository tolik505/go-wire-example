package main

import (
	"testing"
	"wire/internal"
	mock "wire/internal/mocks"

	"go.uber.org/mock/gomock"
)

func TestNewUserHandler(t *testing.T) {
	ctr := gomock.NewController(t)
	httpMock := mock.NewMockHTTP(ctr)
	h, err := initUserHandler(internal.ConnectionInfo("test"), httpMock)
	if err != nil {
		t.Error(err)
	}
	// run handler
	println(h)
}
