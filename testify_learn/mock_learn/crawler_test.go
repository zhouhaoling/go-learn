package mock_learn

import (
	"testing"

	"github.com/stretchr/testify/mock"
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
	MockUsers = append(MockUsers, &User{
		Name: "zhangsan",
		Age:  20,
	})
}

func TestGetUserList(t *testing.T) {
	crawler := new(MockCrawler)
	crawler.On("GetUserList").Return(MockUsers, nil)

	GetAndPrintUser(crawler)
	crawler.AssertExpectations(t)
}
