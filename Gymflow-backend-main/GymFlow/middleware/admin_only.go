package middleware//d все


import (
"net/http"
"github.com/gin-gonic/gin"
)


func AdminOnlyMiddleware(c *gin.Context) {
if c.GetString("user_role") != "admin" {
c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admin access required"})
return
}
c.Next()
}