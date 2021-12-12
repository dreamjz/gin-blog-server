package service

import (
	"errors"
	"gin-blog-server/dao"
	"log"

	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrQueryUser    = errors.New("query user error")
)

func SearchUsername(keywords string) ([]string, error) {
	names, err := dao.FindUsername(keywords)
	if err != nil {
		log.Print("Search username error: ", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, ErrQueryUser
	}
	return names, nil
}
