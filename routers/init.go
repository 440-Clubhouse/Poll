package routers

import (
    "github.com/440-clubhouse/poll/handlers"
    "github.com/gin-gonic/gin"
)

// Init initializes routers of Gin.
func Init(r *gin.Engine) {
    questions := r.Group("/api/questions")
    {
        questions.GET("/", handlers.GetAllQuestions)
        questions.GET("/:id", handlers.GetQuestion)
        questions.POST("/", handlers.CreateQuestion)
        questions.DELETE("/:id", handlers.DeleteQuestion)
    }
}
