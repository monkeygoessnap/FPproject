package server

import (
	"FPproject/Backend/log"
	"FPproject/Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InsertUser(c *gin.Context) {
	var body models.User
	err := c.BindJSON(&body)
	if err != nil {
		log.Info.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
		return
	}
	id, err := h.db.InsertUser(body)
	if err != nil {
		log.Warning.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"id":     id,
	})
}

func verifyID() bool {
	//TODO, verify JWT ID header against request ID
	return true
}

func (h *Handler) DelUser(c *gin.Context) {
	userid := c.Param("id")
	if !verifyID() {
		log.Info.Println("Unauthorized Get", userid)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}
	id, err := h.db.DelUser(userid)
	if err != nil {
		log.Warning.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"id":     id,
	})
}

func (h *Handler) GetUser(c *gin.Context) {
	userid := c.Param("id")
	if !verifyID() {
		log.Info.Println("Unauthorized Del", userid)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}
	var user models.User
	user, err := h.db.GetUser(userid)
	if err != nil {
		log.Warning.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	userid := c.Param("id")
	if !verifyID() {
		log.Info.Println("Unauthorized Del", userid)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}
	var body models.User
	err := c.BindJSON(&body)
	if err != nil {
		log.Info.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
		return
	}
	id, err := h.db.UpdateUser(body)
	if err != nil {
		log.Warning.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"id":     id,
	})
}
