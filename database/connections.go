package database

import (
	"database/sql"
	"masterplan-backend/models"

	//mysql driver imported as _ because it's not used explicitly
	_ "github.com/go-sql-driver/mysql"
)

//DBconn used globally to execute queries
var DBconn *sql.DB

// ConnectToDatabases to establish connection with the DB which will be used
func ConnectToDatabases() {
	dbDriver := "mysql"
	var err error
	DBconn, err = sql.Open(dbDriver, models.Config.Database.Username+":"+models.Config.Database.Password+"@tcp("+models.Config.Database.Host+")/"+models.Config.Database.DatabaseName+models.Config.Database.Flags)
	if err != nil {
		panic(err.Error())
	}
	err = DBconn.Ping() //Ping the DB to confirm the connection as Go lazy loads connection in this case.
	if err != nil {
		panic(err.Error())
	}
}
