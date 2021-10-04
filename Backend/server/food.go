package server

import (
	"FPproject/Backend/log"
	"FPproject/Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InsertFood(c *gin.Context) {
	var body models.Food
	err := c.BindJSON(&body)
	if err != nil {
		log.Info.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
		return
	}
	id, err := h.db.InsertFood(body)
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

func (h *Handler) DelFood(c *gin.Context) {
	foodid := c.Param("id")
	id, err := h.db.DelFood(foodid)
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

func (h *Handler) GetFood(c *gin.Context) {
	foodid := c.Param("id")

	var body models.Food
	body, err := h.db.GetFood(foodid)
	if err != nil {
		log.Warning.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, body)
}

func (h *Handler) GetFoodByMerchant(c *gin.Context) {
	merchantid := c.Param("id")
	var body []models.Food
	body, err := h.db.GetFoodByMerchant(merchantid)
	if err != nil {
		log.Warning.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, body)
}

func (h *Handler) UpdateFood(c *gin.Context) {

	var body models.Food
	err := c.BindJSON(&body)
	if err != nil {
		log.Info.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
		return
	}
	id, err := h.db.UpdateFood(body)
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
