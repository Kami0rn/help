package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Kami0rn/SoyJuu/controller"
	"github.com/Kami0rn/SoyJuu/entities"
)

func main() {
	entities.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	//User Routes
	r.GET("/customer", controller.ListCustomer)
	r.GET("/customer/:id", controller.GetCustomer)
	r.POST("/customer", controller.CreateCustomer)	
	r.PATCH("/customer/:id", controller.UpdateCustomer)
	r.DELETE("/customer/:id", controller.DeleteCustomer)

	//Run the server

	r.Run()

}


func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
	
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		
		
		if c.Request.Method == "OPTIONS" {
		
			c.AbortWithStatus(204)
			
			return
		
		}
		
		
		c.Next()
	
	}
	
}
