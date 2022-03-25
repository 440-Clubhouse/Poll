package handlers

import (
    "net/http"
    "strconv"
    "time"

    "github.com/440-clubhouse/poll/models"
    "github.com/gin-gonic/gin"
)

func GetQuestionnaire(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "code": http.StatusBadRequest,
            "msg":  err,
        })
        return
    }

    var questionnaire models.Questionnaire
    if err = questionnaire.Get(id); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "code": http.StatusInternalServerError,
            "msg":  err,
        })
        return
    }

    c.JSON(http.StatusOK, questionnaire)
}

func GetAllQuestionnaires(c *gin.Context) {
    var questionnaires models.Questionnaires
    if err := questionnaires.GetAll(); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "code": http.StatusInternalServerError,
            "msg":  err,
        })
        return
    }

    c.JSON(http.StatusOK, questionnaires)
}

func CreateQuestionnaire(c *gin.Context) {
    var questionnaire models.Questionnaire
    if err := c.ShouldBindJSON(&questionnaire); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "code": http.StatusBadRequest,
            "msg":  err,
        })
        return
    }

    questionnaire.Date = time.Now().Format("2006-01-02 15:04:05")
    if err := questionnaire.Create(); err != nil {
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

func DeleteQuestionnaire(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "code": http.StatusBadRequest,
            "msg":  err,
        })
        return
    }

    var questionnaire models.Questionnaire
    if err = questionnaire.Delete(id); err != nil {
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
