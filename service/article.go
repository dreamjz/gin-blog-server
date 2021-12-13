package service

import (
	"errors"
	"gin-blog-server/dao"
	"gin-blog-server/models"
	"gin-blog-server/models/request"
	"gin-blog-server/models/response"
	"log"

	"gorm.io/gorm"
)

var (
	ErrCreateArticle   = errors.New("create article error")
	ErrQueryArticle    = errors.New("query article error")
	ErrArticleNotFound = errors.New("article not found")
	ErrUpdateArticle   = errors.New("update article error")
	ErrDeleteArticle   = errors.New("delete article error")
)

func CreateArticle(article *models.Article) error {

	err := dao.CreateArticle(article)
	if err != nil {
		log.Print("Create article error: ", err.Error())
		return ErrCreateArticle
	}
	return nil
}

func QueryArticleList(pagination request.Pagination) (response.PageResult, error) {
	var result response.PageResult
	limit := pagination.PageSize
	offset := (pagination.Page - 1) * limit
	total, err := dao.CountArticle()
	if err != nil {
		log.Print("Count article error: ", err.Error())
		return result, ErrQueryArticle
	}
	if total < 1 {
		log.Print("No article found")
		return result, nil
	}
	articleList, err := dao.FindArticleList(offset, limit)
	if err != nil {
		log.Println("query article list error: ", err.Error())
		return result, ErrQueryArticle
	}
	result = response.PageResult{
		List:     articleList,
		Total:    total,
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
	}
	return result, nil
}

func QueryArticleByID(id uint) (*response.ArticleDetail, error) {
	content, err := dao.FindArticleByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Print("Article not found")
			return nil, ErrArticleNotFound
		}
		return nil, ErrQueryArticle
	}
	return content, nil
}

func UpdateArticleByID(article *models.Article) error {
	if err := dao.UpdateArticleByID(article); err != nil {
		log.Print("Update article error: ", err.Error())
		return ErrUpdateArticle
	}
	return nil
}

func DeleteArticleByID(id uint) error {
	if err := dao.DeleteArticleByID(id); err != nil {
		log.Print("Delete article error: ", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrArticleNotFound
		}
		return ErrDeleteArticle
	}
	return nil
}
