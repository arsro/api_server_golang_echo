package api

import (
	"github.com/labstack/echo"
	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
	
	"net/http"
	"strconv"
	
	"api/model"
	"api/config"
)

func HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World!")
}

func GetUser() echo.HandlerFunc {
	return func (c echo.Context) (err error) {
		user_id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			logrus.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, "user_id is null")
		}
		tx := c.Get(config.TX_KEY).(*dbr.Tx)
		
		user := new(model.User)
		if err := user.Select(tx, user_id); err != nil {
			logrus.Error(err)
			return echo.NewHTTPError(http.StatusNotFound, "user not exists.")
		}
		return c.JSON(http.StatusOK, user)
	}
}


func PostUser() echo.HandlerFunc {
	return func (c echo.Context) (err error) {
		u := new(model.User)
		
		if err := c.Bind(&u); err != nil{
			logrus.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		
		tx := c.Get(config.TX_KEY).(*dbr.Tx)
		
		user := model.NewUser(u.Name)
		if err := user.Insert(tx); err != nil{
			logrus.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		
		return c.JSON(http.StatusCreated, user)
	}
}


func PutUser() echo.HandlerFunc {
	return func (c echo.Context) (err error) {
		u := new(model.User)
		if err := c.Bind(u); err != nil {
			logrus.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		
		
		tx := c.Get(config.TX_KEY).(*dbr.Tx)
		if err := u.Update(tx); err != nil{
			logrus.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		
		return c.JSON(http.StatusCreated, u)
	}
	
}

func DeleteUser() echo.HandlerFunc  {
	return func (c echo.Context) (err error) {
		user_id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			logrus.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, "user_id is null")
		}
		tx := c.Get(config.TX_KEY).(*dbr.Tx)
		
		u := new(model.User)
		if err := u.Delete(tx, user_id); err != nil{
			logrus.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		
		return c.JSON(http.StatusCreated, user_id)
	}
	
}
