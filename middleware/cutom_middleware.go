package middleware

import (
	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	
	"api/config"
)


/**
 * dbのトランジション管理
 */
func TransactionHandler(db *dbr.Session) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) (err error) {
			logrus.Debug("start ransaction")

			tx, err := db.Begin()
			if err != nil {
				logrus.Error("transction error: ", err)
				return err
			}
			defer tx.RollbackUnlessCommitted()// この function の終了時に実行される
			
			// Set saves data in the context.
			c.Set(config.TX_KEY, tx)
			if err := next(c); err != nil {
				tx.Rollback()
				logrus.Error("transction rollback: ", err)
				return err
			}
			
			logrus.Debug("commit ransaction")
			tx.Commit()
	
			return nil
		})
	}
}
