package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	
	"net/http"
	
	"api/api"
	"api/database"
	CustomMW "api/middleware"
)

func Init() *echo.Echo {
	
	e := echo.New()

	// Debug mode
	e.Debug = true
	
	SetMiddleware(e)
	SetRoutting(e)
	return e
}



func SetMiddleware(e *echo.Echo) {
	
	// Set Bundle MiddleWare
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())
	// Setting CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	session, err := database.Init()
	if err != nil {
		echo.NewHTTPError(http.StatusNotFound, err)
	}
	// set cutom middleware
	e.Use(CustomMW.TransactionHandler(session))
}

func SetRoutting(e *echo.Echo) {
	//Routing
	e.GET("/", api.HelloWorld)
	
	v1 := e.Group("/api")
	{
		v1.GET( 	"/users/:id",	api.GetUser()	)
		v1.POST(	"/users",		api.PostUser()	)
		v1.PUT( 	"/users",		api.PutUser()	)
		v1.DELETE(	"/users/:id", 	api.DeleteUser())
	}
}
