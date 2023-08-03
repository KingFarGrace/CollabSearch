package middleware

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	router "github.com/KingFarGrace/CollabSearch/server/router/response"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate *validator.Validate

// ValidateRegisterJSON is a middleware to validate user register data.
// @See RegisterJSON for detail rules.
// Pass verified RegisterJSON or Abort with return code http.StatusBadRequest.
func ValidateRegisterJSON() gin.HandlerFunc {
	return func(context *gin.Context) {
		var jsonObj entity.RegisterJSON
		resp := router.AccountResponse{}
		if err := context.ShouldBindJSON(&jsonObj); err != nil {
			resp.New(router.AccountGroupCode, 1, "Invalid request json.")
			context.JSON(http.StatusBadRequest, resp)
			context.Abort()
			return
		}
		validate = validator.New()
		err := validate.Struct(&jsonObj)
		if err != nil {
			errMsg := "Validate failed at: "
			for _, err := range err.(validator.ValidationErrors) {
				errMsg = util.Concat(errMsg, "\n", err.StructNamespace(), " ", err.Value())
			}
			resp.New(router.AccountGroupCode, 1, errMsg)
			context.JSON(http.StatusBadRequest, resp)
			context.Abort()
			return
		}
		context.Set("jsonObj", jsonObj)
		context.Next()
	}
}
