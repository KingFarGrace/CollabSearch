package middleware

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// ValidateRegisterJSON is a middleware to validate user register data.
// @See entity.RegisterJSON for detail validation rules.
func ValidateRegisterJSON() gin.HandlerFunc {
	var jsonObj entity.RegisterJSON
	return validateJSON(jsonObj)
}

// ValidateLoginJSON is a middleware to validate user login data.
// @See entity.LoginJSON for detail validation rules.
func ValidateLoginJSON() gin.HandlerFunc {
	var jsonObj entity.LoginJSON
	return validateJSON(jsonObj)
}

// ValidateUpdateJSON is a middleware to validate user update data.
// @See entity.User for detail validation rules.
func ValidateUpdateJSON() gin.HandlerFunc {
	var jsonObj entity.User
	return validateJSON(jsonObj)
}

func ValidateUidJSON() gin.HandlerFunc {
	var jsonObj entity.UidJSON
	return validateJSON(jsonObj)
}

// ValidateWorkspaceJSON is a middleware to validate workspace data.
// @See entity.Workspace for detail validation rules.
func ValidateWorkspaceJSON() gin.HandlerFunc {
	var jsonObj entity.Workspace
	return validateJSON(jsonObj)
}

// ValidateUserWorkspaceJSON is a middleware to validate user_workspace data.
// @See entity.UserWorkspace for detail validation rules.
func ValidateUserWorkspaceJSON() gin.HandlerFunc {
	var jsonObj entity.UserWorkspace
	return validateJSON(jsonObj)
}

// ValidateResultJSON is a middleware to validate user_workspace data.
// @See entity.Result for detail validation rules.
func ValidateResultJSON() gin.HandlerFunc {
	var jsonObj entity.Result
	return validateJSON(jsonObj)
}

// ValidateSearchingJSON is a middleware to validate user_workspace data.
// @See entity.SearchingJSON for detail validation rules.
func ValidateSearchingJSON() gin.HandlerFunc {
	var jsonObj entity.SearchingJSON
	return validateJSON(jsonObj)
}

// validateJSON
// An encapsulation of different middleware func.
// Pass verified json obj or abort with return code http.StatusBadRequest.
// jsonObj <T> must be defined in interface entity.Serializable.
func validateJSON[T entity.Serializable](jsonObj T) gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := context.ShouldBindBodyWith(&jsonObj, binding.JSON); err != nil {
			util.ErrorLogger(err, "ShouldBindBodyWith()")
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Msg": "Invalid request json.", "jsonObj": jsonObj})
			return
		}
		if msg := valid(jsonObj); msg != "" {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Msg": msg})
			return
		}
		context.Set("jsonObj", jsonObj)
		context.Next()
	}
}

// valid is an encapsulation of validate process.
// Return errMsg if failed to validated json data, otherwise return ok (empty string).
func valid[T entity.Serializable](jsonObj T) string {
	validate := validator.New()
	ok := ""
	err := validate.Struct(jsonObj)
	if err != nil {
		errMsg := "Validate failed at:"
		for _, err := range err.(validator.ValidationErrors) {
			errMsg = util.Concat(errMsg, " ", err.StructNamespace())
		}
		return errMsg
	}
	return ok
}
