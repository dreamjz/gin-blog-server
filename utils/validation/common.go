package validation

import (
	"gin-blog-server/models"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterStructValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(ArticleStructLevelValidation, models.Article{})
	}
}
