package handlers

import (
    "net/http"
    "strconv"
    "time"

    "github.com/440-clubhouse/poll/models"
    "github.com/gin-gonic/gin"
)

func GetQuestion(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "code": http.StatusBadRequest,
            "msg":  err,
        })
        return
    }

    var question models.Question
    if err = question.Get(id); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "code": http.StatusInternalServerError,
            "msg":  err,
        })
        return
    }

    c.JSON(http.StatusOK, question)
}

func GetAllQuestions(c *gin.Context) {
    var questions models.Questions
    if err := questions.GetAll; err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "code": http.StatusInternalServerError,
            "msg":  err,
        })
        return
    }

    c.JSON(http.StatusOK, questions)
}

func CreateQuestion(c *gin.Context) {
    var question models.Question
    if err := c.ShouldBindJSON(&question); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "code": http.StatusBadRequest,
            "msg":  err,
        })
        return
    }

    question.Date = time.Now().Format("2006-01-02 15:04:05")
    if err := question.Create(); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "code": http.StatusInternalServerError,
            "msg":  err,
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "code": http.StatusOK,
        "msg":  "success",
    })
}

func DeleteQuestion(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "code": http.StatusBadRequest,
            "msg":  err,
        })
        return
    }

    var question models.Question
    if err = question.Delete(id); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "code": http.StatusInternalServerError,
            "msg":  err,
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "code": http.StatusOK,
        "msg":  "success",
    })
}
