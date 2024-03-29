package mock_learn

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

type ICrawler interface {
	GetUserList() ([]*User, error)
}

type MyCrawler struct {
	url string
}

func (c *MyCrawler) GetUserList() ([]*User, error) {
	resp, err := http.Get(c.url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userList []*User
	err = json.Unmarshal(data, &userList)
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func GetAndPrintUser(crawler ICrawler) {
	users, err := crawler.GetUserList()
	if err != nil {
		return
	}

	for _, u := range users {
		fmt.Print(u)
	}
}
