package middlewares

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)


// CORSInit initializes CORS middleware of Gin.
func CORSInit(r *gin.Engine) {
    r.Use(cors.Default())
}