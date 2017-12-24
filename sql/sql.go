package sql

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var dbconf map[string]string = map[string]string{
	"user":     "root",
	"password": "",
	"host":     "127.0.0.1:3306",
	"name":     "iotame",
}

func dburl() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=5s",
		dbconf["user"], dbconf["password"], dbconf["host"], dbconf["name"])
}

func UpdateNode(lastSolid, appName, appVersion, id string) {

	db, err := sql.Open("mysql", dburl())
	if err != nil {
		return
	}
	defer db.Close()

	statement, _ := db.Prepare("update nodes set last_solid = ?, app_name = ?, app_version = ?, connected_at = now() where id = ?")
	statement.Exec(lastSolid, appName, appVersion, id)
	defer statement.Close()
}

func Nodes() []map[string]string {
	list := []map[string]string{}

	db, err := sql.Open("mysql", dburl())
	if err != nil {
		return list
	}
	defer db.Close()

	rows, _ := db.Query("select id,connection_type from nodes where connection_type='tcp' order by id;")
	var id string
	var connectionType string
	for rows.Next() {
		rows.Scan(&id, &connectionType)
		node := map[string]string{
			"id":              id,
			"connection_type": connectionType,
		}
		list = append(list, node)
	}
	defer rows.Close()
	return list
}
