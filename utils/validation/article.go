package validation

import (
	"gin-blog-server/models"

	"github.com/go-playground/validator/v10"
)

func ArticleStructLevelValidation(sl validator.StructLevel) {
	article := sl.Current().Interface().(models.Article)
	if *article.Status < 0 || *article.Status > 1 {
		sl.ReportError(article.Status, "Status", "Status", "status", "")
	}
	if article.Importance < 1 || article.Importance > 7 {
		sl.ReportError(article.Importance, "Importance", "Importance", "importance", "")
	}
}
