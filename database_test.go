package godatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "udin:udin@tcp(host:3306)/test_go")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
