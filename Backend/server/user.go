package server

import (
	"FPproject/Backend/auth"
	"FPproject/Backend/log"
	"FPproject/Backend/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
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
		if err.(*mysql.MySQLError).Number == 1062 {
			c.JSON(http.StatusConflict, gin.H{
				"status": "username already exists",
			})
			return
		}
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

func (h *Handler) GetMerchants(c *gin.Context) {
	info, err := h.db.GetMerchants()
	if err != nil {
		log.Warning.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "internal server error",
		})
	}
	c.JSON(http.StatusOK, info)
}

func (h *Handler) DelUser(c *gin.Context) {

	userid := c.Keys["ID"].(string)
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

	userid := c.Keys["ID"].(string)
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
	var body models.User
	err := c.BindJSON(&body)
	if err != nil {
		log.Info.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
		return
	}
	body.ID = c.Keys["ID"].(string)
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

func (h *Handler) Login(c *gin.Context) {
	userCred := map[string]string{
		"username": "",
		"password": "",
	}
	err := c.BindJSON(&userCred)
	if err != nil {
		log.Info.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
		return
	}
	user, err := h.db.Validate(userCred["username"], userCred["password"])
	if err == sql.ErrNoRows {
		log.Info.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	} else if err != nil {
		log.Warning.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "internal server error",
		})
		return
	}

	token, err := auth.GenerateJWT(user.Username, user.ID)
	if err != nil {
		log.Info.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, token)
}
