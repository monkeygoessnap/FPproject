package server

import (
	"FPproject/Backend/log"
	"FPproject/Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InsertUH(c *gin.Context) {
	var body models.UserHealth
	err := c.BindJSON(&body)
	if err != nil {
		log.Info.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
		return
	}
	userid := c.Keys["ID"].(string)
	id, err := h.db.InsertUH(userid, body)
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

func (h *Handler) DelUH(c *gin.Context) {
	userid := c.Keys["ID"].(string)
	id, err := h.db.DelUH(userid)
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

func (h *Handler) GetUH(c *gin.Context) {
	userid := c.Keys["ID"].(string)
	var body models.UserHealth
	body, err := h.db.GetUH(userid)
	if err != nil {
		log.Warning.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, body)
}

func (h *Handler) UpdateUH(c *gin.Context) {

	var body models.UserHealth
	err := c.BindJSON(&body)

	if err != nil {
		log.Info.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
		return
	}
	body.ID = c.Keys["ID"].(string)
	id, err := h.db.UpdateUH(body)
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
