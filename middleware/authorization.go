package middleware

import (
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/yodfhafx/go-crud/models"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, ok := ctx.Get("sub")
		if !ok {
			// AbortWithStatusJSON = next handler won't working
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		enforcer, err := casbin.NewEnforcer("config/acl_model.conf", "config/policy.csv")
		if err != nil {
			log.Fatal("error: enforcer:", err)
		}

		ok, _ = enforcer.Enforce(user.(*models.User), ctx.Request.URL.Path, ctx.Request.Method)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "you are not allowed to access this resource",
			})
			return
		}
		// call next handler
		ctx.Next()
	}
}
