package main

import (
	"net/http"
	"theGame/database"
	"theGame/users/admin"
	"theGame/users/player"
	"theGame/users/support"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitPlayer()
	router :=gin.Default()
	router.GET("/admin/:id", getAdminByID)
	router.GET("/admin", getAdmin)
	router.GET("/player", getPlayer)
	router.GET("/player/:id", getPlayerByID)
	router.GET("/support", getSupport)
	router.GET("/support/:id", getSupportByID)
	router.Run(":8083")
		
}

func getAdmin(c *gin.Context){
	c.IndentedJSON(http.StatusOK, admin.ADMIN_MAP)
}

func getPlayer(c *gin.Context){
	c.IndentedJSON(http.StatusOK, player.PLAYER_MAP)
}

func getSupport(c *gin.Context){
	c.IndentedJSON(http.StatusOK, support.SUPPORT_MAP)
}

func getAdminByID(c *gin.Context){
	id :=c.GetInt("id")

	for _, a :=range admin.ADMIN_MAP{
		if  a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}

func getPlayerByID(c *gin.Context){
	id :=c.GetInt("id")

	for _, a :=range player.PLAYER_MAP{
		if  a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}

func getSupportByID(c *gin.Context){
	id :=c.GetInt("id")

	for _, a :=range support.SUPPORT_MAP{
		if  a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}