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
	userid := c.Keys["ID"].(string)
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
	userid := c.Keys["ID"].(string)
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
	userid := c.Keys["ID"].(string)
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

func (h *Handler) GetMercAdd(c *gin.Context) {
	id := c.Param("id")
	add, err := h.db.GetAdd(id)
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
	var body models.Address
	err := c.BindJSON(&body)
	if err != nil {
		log.Info.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
		return
	}
	body.ID = c.Keys["ID"].(string)
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
