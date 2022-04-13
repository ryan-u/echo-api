package main
import (
    "test/controller"
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
 
    e := echo.New()
    
    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
 
    //CORS
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
    }))
    
    // Root route => handler  
    e.GET("/", func(c echo.Context) error {       
        return c.String(http.StatusOK, "Hello, World!\n")  
    })
    
    e.GET("/dogs/:data", controller.GetDogs)
    e.POST("/dogs", controller.AddDog)
    
    // Run Server
    e.Logger.Fatal(e.Start(":8000"))
}