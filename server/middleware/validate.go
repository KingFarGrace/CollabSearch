package middleware

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// ValidateRegisterJSON is a middleware to validate user register data.
// @See RegisterJSON for detail validation rules.
func ValidateRegisterJSON() gin.HandlerFunc {
	var jsonObj entity.RegisterJSON
	return validateJSON(jsonObj)
}

// ValidateLoginJSON is a middleware to validate user login data.
// This middleware won't be handled if the request pass the token verification middleware.
// @See LoginJSON for detail validation rules.
func ValidateLoginJSON() gin.HandlerFunc {
	var jsonObj entity.LoginJSON
	return validateJSON(jsonObj)
}

// validateJSON
// An encapsulation of different middleware func.
// Pass verified json obj or abort with return code http.StatusBadRequest.
// jsonObj <T> should only be entities in package entity.
func validateJSON[T entity.LoginJSON | entity.RegisterJSON](jsonObj T) gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := context.ShouldBindJSON(&jsonObj); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Invalid request json."})
			return
		}
		if msg := valid(jsonObj); msg != "" {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": msg})
			return
		}
		context.Set("jsonObj", jsonObj)
		context.Next()
	}
}

// valid is an encapsulation of validate process.
// Return errMsg if failed to validated json data, otherwise return ok (empty string).
func valid[T entity.LoginJSON | entity.RegisterJSON](jsonObj T) string {
	validate := validator.New()
	ok := ""
	err := validate.Struct(jsonObj)
	if err != nil {
		errMsg := "Validate failed at: "
		for _, err := range err.(validator.ValidationErrors) {
			errMsg = util.Concat(errMsg, "\n", err.StructNamespace(), " ", err.Value())
		}
		return errMsg
	}
	return ok
}
