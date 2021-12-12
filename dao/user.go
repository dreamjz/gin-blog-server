package dao

import "gin-blog-server/models"

func FindUsername(keywords string) ([]string, error) {
	var names []string
	err := db.Model(&models.User{}).Select("username").
		Where("username REGEXP ?", keywords).Find(&names).Error
	return names, err
}
