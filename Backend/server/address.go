package server

import (
	"FPproject/Backend/log"
	"FPproject/Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InsertAdd(c *gin.Context) {
	var body models.Address
	err := c.BindJSON(&body)
	if err != nil {
		log.Info.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
		return
	}
	//use params first then next commit maybe change to JWT ID header
	userid := c.Param("id")
	id, err := h.db.InsertAdd(userid, body)
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

func (h *Handler) DelAdd(c *gin.Context) {
	userid := c.Param("id")
	if !verifyID() {
		log.Info.Println("Unauthorized Del", userid)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}
	id, err := h.db.DelAdd(userid)
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

func (h *Handler) GetAdd(c *gin.Context) {
	userid := c.Param("id")
	if !verifyID() {
		log.Info.Println("Unauthorized Get", userid)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}
	var add models.Address
	add, err := h.db.GetAdd(userid)
	if err != nil {
		log.Warning.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, add)
}

func (h *Handler) UpdateAdd(c *gin.Context) {
	userid := c.Param("id")
	if !verifyID() {
		log.Info.Println("Unauthorized Del", userid)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}
	var body models.Address
	err := c.BindJSON(&body)
	if err != nil {
		log.Info.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
		return
	}
	id, err := h.db.UpdateAdd(body)
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
