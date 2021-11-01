/*
Package server provides the router functions
*/
package server

import (
	"FPproject/Backend/auth"
	"FPproject/Backend/database"
	"FPproject/Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//must be same as database user functions
type repository interface {
	Validate(um, pw string) (models.User, error)
	InsertUser(user models.User) (string, error)
	DelUser(id string) (string, error)
	UpdateUser(user models.User) (string, error)
	GetUser(id string) (models.User, error)
	GetMerchants() ([]models.User, error)

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
	GetFoodByMerchant(id string) ([]models.Food, error)

	InsertCI(ci models.CartItem) (string, error)
	DelCI(id, userid string) (string, error)
	UpdateCI(f models.CartItem) (string, error)
	GetCI(id string) (models.CartItem, error)
	GetCIByUser(id string) ([]models.CartItem, error)
	DelAllCI(id string) (string, error)
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

	apiVersion := "/api/v1"

	router := gin.Default()
	db := connectDB()
	defer db.Close()

	r := database.New(db)
	h := handler(r)

	//unauthenticated routes
	public := router.Group(apiVersion)

	public.GET("/healthcheck", Healthcheck)
	public.POST("/login", h.Login)
	public.POST("/register", h.InsertUser)

	//authenticated routes
	private := router.Group(apiVersion)
	private.Use(auth.AuthJWT())

	private.DELETE("/user", h.DelUser)
	private.GET("/user", h.GetUser)
	private.PUT("/user", h.UpdateUser)
	private.GET("/merc", h.GetMerchants)

	private.POST("/add", h.InsertAdd)
	private.DELETE("/add", h.DelAdd)
	private.GET("/add", h.GetAdd)
	private.PUT("/add", h.UpdateAdd)
	private.GET("/mercadd/:id", h.GetMercAdd)

	private.POST("/uh", h.InsertUH)
	private.DELETE("/uh", h.DelUH)
	private.GET("/uh", h.GetUH)
	private.PUT("/uh", h.UpdateUH)

	private.POST("/food", h.InsertFood)
	private.DELETE("/food/:id", h.DelFood)
	private.GET("/food/:id", h.GetFood)
	private.PUT("/food/:id", h.UpdateFood)
	private.GET("/allfood/:id", h.GetFoodByMerchant)

	private.POST("/ci", h.InsertCI)
	private.DELETE("/ci/:id", h.DelCI)
	private.GET("/ci/:id", h.GetCI)
	private.PUT("/ci", h.UpdateCI)
	private.GET("/allci", h.GetCIByUser)
	private.DELETE("/ci", h.DelAllCI)

	//log.Error.Fatal(http.ListenAndServeTLS(":8080", "certs/cert.pem", "certs/key.pem", router))
	router.RunTLS(":8080", "certs/cert.pem", "certs/key.pem")
	//router.Run()
}
