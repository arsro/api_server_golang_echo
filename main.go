package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	
	/**
	 * Middleware
	 */
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//CORSを許可する設定???
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	
	/**
	* Routing
	*/
	e.GET("/", func(c echo.Context) error {
		 return c.JSON(http.StatusOK, "Hello, World!")
	})
	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
