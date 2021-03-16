package routes

import (
	"database/sql"
	"snippetBox-microservice/basket/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) *gin.Engine {

	r := gin.Default()

	r.GET("/", controllers.Home)
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.GET("/index", controllers.Productlist)
	r.POST("/addwishlists", controllers.Addwishlist)
	r.POST("/addtoorder", controllers.Addtoorder)
	r.GET("/wishlist", controllers.Wishlist)
	r.GET("/deletewishs", controllers.Deletewishlist)
	r.GET("/removeall", controllers.Deletewishlistall)
	r.GET("/signout", controllers.Logout)
	r.GET("/myorders", controllers.Viewmyorders)

	return r
}