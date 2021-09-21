package server

import (
	"FPproject/Backend/database"
	"FPproject/Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//must be same as database user functions
type repository interface {
	InsertUser(user models.User) (string, error)
	DelUser(id string) (string, error)
	UpdateUser(user models.User) (string, error)
	GetUser(id string) (models.User, error)

	InsertAdd(id string, add models.Address) (string, error)
	DelAdd(id string) (string, error)
	UpdateAdd(add models.Address) (string, error)
	GetAdd(id string) (models.Address, error)

	InsertUH(id string, h models.UserHealth) (string, error)
	DelUH(id string) (string, error)
	UpdateUH(h models.UserHealth) (string, error)
	GetUH(id string) (models.UserHealth, error)

	InsertFood(f models.Food) (string, error)
	DelFood(id string) (string, error)
	UpdateFood(f models.Food) (string, error)
	GetFood(id string) (models.Food, error)

	InsertCI(ci models.CartItem) (string, error)
	DelCI(id string) (string, error)
	UpdateCI(f models.CartItem) (string, error)
	GetCI(id string) (models.CartItem, error)
}

type Handler struct {
	db repository
}

//passes dependency
func handler(db repository) *Handler {
	return &Handler{
		db: db,
	}
}

func Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func InitServer() {
	router := gin.Default()
	db := connectDB()
	defer db.Close()

	r := database.New(db)
	h := handler(r)

	router.GET("/healthcheck", Healthcheck)

	router.POST("/user", h.InsertUser)
	router.DELETE("/user/:id", h.DelUser)
	router.GET("/user/:id", h.GetUser)
	router.PUT("/user/:id", h.UpdateUser)

	router.POST("/add/:id", h.InsertAdd)
	router.DELETE("/add/:id", h.DelAdd)
	router.GET("/add/:id", h.GetAdd)
	router.PUT("/add/:id", h.UpdateAdd)

	router.POST("/uh/:id", h.InsertUH)
	router.DELETE("/uh/:id", h.DelUH)
	router.GET("/uh/:id", h.GetUH)
	router.PUT("/uh/:id", h.UpdateUH)

	router.POST("/food", h.InsertFood)
	router.DELETE("/food/:id", h.DelFood)
	router.GET("/food/:id", h.GetFood)
	router.PUT("/food/:id", h.UpdateFood)
	//custom getall TODO

	router.POST("/ci", h.InsertCI)
	router.DELETE("/ci/:id", h.DelCI)
	router.GET("/ci/:id", h.GetCI)
	router.PUT("/ci/:id", h.UpdateCI)
	//custom getall TODO

	router.Run()
}
