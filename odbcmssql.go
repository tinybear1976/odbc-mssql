package odbcmssql

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

var (
	serverTags map[string]*sqlx.DB = make(map[string]*sqlx.DB)
)

func New(serverTag, server, userid, password, database string, port int) error {
	connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d;encrypt=disable", server, database, userid, password, port)
	db, err := sqlx.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
		return err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("PING:", err)
		return err
	}
	serverTags[serverTag] = db
	return nil
}

func Connect(serverTag string) (*sqlx.DB, error) {
	sqldb, ok := serverTags[serverTag]
	if !ok {
		return nil, errors.New(fmt.Sprintf("odbc mssqldb[%s] not existing", serverTag))
	}
	return sqldb, nil
}

func Destroy() {
	for k := range serverTags {
		delete(serverTags, k)
	}
}
