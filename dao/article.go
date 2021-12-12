package dao

import (
	"gin-blog-server/models"
	"gin-blog-server/models/response"
)

func CreateArticle(article *models.Article) error {
	return db.Create(article).Error
}

func FindArticleList(offset, limit int) ([]response.ArticleListResult, error) {
	var articleList []response.ArticleListResult
	err := db.Model(&models.Article{}).Offset(offset).Limit(limit).Find(&articleList).Error
	return articleList, err
}

func FindArticleContentByID(id uint) (*response.ArticleContent, error) {
	var content response.ArticleContent
	err := db.Model(&models.Article{}).Where("id = ?", id).Take(&content).Error
	return &content, err
}

func CountArticle() (int64, error) {
	var count int64
	err := db.Model(&models.Article{}).Count(&count).Error
	return count, err
}

func UpdateArticleByID(article *models.Article) error {
	return db.Save(article).Error
}

func DeleteArticleByID(id uint) error {
	return db.Where("id = ?", id).Delete(&models.Article{}).Error
}
