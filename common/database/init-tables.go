package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/api/common/config"
)

func Initialize(dbDriver *sqlx.DB,cfg *config.AppConfig) {
	// statement, driverError := dbDriver.Prepare(train)
	// if driverError != nil {
	// 	cfg.Logger.Error(driverError.Error())
		
	// }
	// // Create train table
	// _, statementError := statement.Exec()
	// if statementError != nil {
	// 	log.Println("Table already exists!")
	// }
	// statement, _ = dbDriver.Prepare(station)
	// statement.Exec()
	// statement, _ = dbDriver.Prepare(schedule)
	// statement.Exec()

}
