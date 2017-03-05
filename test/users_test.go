package api

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"testing"
	"net/http"
	"net/http/httptest"
	
	"api/api"
	"api/model"
)

var (
	Users []model.User
)


func TestHelloWorld(t *testing.T)  {
	// Setup
	e := echo.New()
	req := new(http.Request)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	
	// Assertions
	if assert.NoError(t, api.HelloWorld(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// assert.Equal(t, "hogehoge", res.Body.String())
	}
	t.Log("hogehoge", res)
}

// func TestGetUser(t *testing.T)  {
// 	// Setup
// 	e := echo.New()
// 	req := new(http.Request)
// 	res := httptest.NewRecorder()
//   	c := e.NewContext(req, rec)
//
// 	c.SetPath("api/users/:id")
// 	c.SetParamNames("id")
// 	c.SetParamValues("3")
//
// 	// Assertions
// 	if assert.NoError(t, api.GetUser(c)) {
// 		assert.Equal(t, http.StatusOK, res.Code)
// 		assert.Equal(t, "hogehoge", res.Body.String())
// 	}
// }
