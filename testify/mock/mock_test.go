package main

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockCrawler struct {
	mock.Mock
}

func (m *MockCrawler) GetUserList() ([]*User, error) {
	args := m.Called()
	return args.Get(0).([]*User), args.Error(1)
}

var (
	MockUsers []*User
)

func init() {
	MockUsers = append(MockUsers, &User{"dj", 18})
	MockUsers = append(MockUsers, &User{"zhangsan", 20})
}

func TestGetUserList(t *testing.T) {
	crawler := new(MockCrawler)
	crawler.On("GetUserList").Return(MockUsers, nil)

	GetAndPrintUsers(crawler)

	crawler.AssertExpectations(t)
}
