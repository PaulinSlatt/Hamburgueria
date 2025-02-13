package routes

import (
	"GoEcho/controllers"

	"github.com/labstack/echo/v4"
)

func HandleRequest(e *echo.Echo) {

	e.POST("/burgers", controllers.CreateBurger)
	e.GET("/burgers", controllers.GetBurgers)
	e.PATCH("/burgers/:id", controllers.UpdateBurgerStatus)
	e.DELETE("/burgers/:id", controllers.DeleteBurger)

	e.GET("/ingredientes", controllers.GetIngredientes)

	//e.GET("/paes", controllers.GetPaes)

	//e.GET("/carnes", controllers.GetCarnes)

	e.GET("/status", controllers.GetStatus)
}
