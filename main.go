package main

//import "fmt"

//import "github.com/iotaledger/giota"
//import "database/sql"
//import "time"
//import _ "github.com/go-sql-driver/mysql"
//import "strings"
import "github.com/andrewarrow/iotame/client"
import "github.com/andrewarrow/iotame/sql"

func main() {
	for _, node := range sql.Nodes() {
		client.Connect(node)
	}
}

/*
var seen []string = []string{"999999999999999999999999999999999999999999999999999999999999999999999999999999999"}
var unseen []string = []string{}

func InsertTx(id string, ts time.Time, value int64, address string) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/iotame?timeout=5s")
	if err != nil {
		return
	}
	defer db.Close()

	statement, _ := db.Prepare("INSERT INTO transactions (id, ts, value, address) VALUES (?, ?, ?, ?)")
	statement.Exec(id, ts, value, address)
	defer statement.Close()
}

func InsertNode(id string, connectionType string) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/iotame?timeout=5s")
	if err != nil {
		return
	}
	defer db.Close()

	statement, _ := db.Prepare("insert into nodes (id, connection_type) values (?,?);")
	statement.Exec(id, connectionType)
	defer statement.Close()
}

func UpdateNode(appName, appVersion, id string) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/iotame?timeout=5s")
	if err != nil {
		return
	}
	defer db.Close()

	statement, _ := db.Prepare("update nodes set app_name = ?, app_version = ?, connected_at = now() where id = ?")
	statement.Exec(appName, appVersion, id)
	defer statement.Close()
}

func ListNodes() []string {

	var list []string = []string{}
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/iotame?timeout=5s")
	if err != nil {
		return list
	}
	defer db.Close()

	rows, _ := db.Query("select id,connection_type from nodes where connection_type='tcp';")
	var id string
	var connectionType string
	for rows.Next() {
		rows.Scan(&id, &connectionType)
		if connectionType == "tcp" {
			if strings.HasSuffix(id, ":443") {
				list = append(list, fmt.Sprintf("https://%s", id))
			} else {
				list = append(list, fmt.Sprintf("http://%s", id))
			}
		} else {
			list = append(list, fmt.Sprintf("udp://%s", id))
		}
	}
	defer rows.Close()
	return list
}

func follow(node string, api *giota.API) {
	t, _ := giota.ToTrytes(node)
	resp, err := api.GetTrytes([]giota.Trytes{t})
	if err == nil {
		for _, tx := range resp.Trytes {
			InsertTx(node, tx.Timestamp, tx.Value, string(tx.Address))
			trunk := string(tx.TrunkTransaction)
			branch := string(tx.BranchTransaction)
			fmt.Println(tx.Timestamp)
			follow(trunk, api)
			follow(branch, api)
		}
	}
}

func main() {
	for _, n := range ListNodes() {
		fmt.Println(n)
		api := giota.NewAPI(n, nil)
		resp, err := api.GetNodeInfo()
		fmt.Println(err)
		if err == nil {
			//fmt.Println(resp.Neighbors)
			//fmt.Println(resp.AppVersion)
			//fmt.Println(resp.AppName)
			tokens := strings.Split(n, "/")
			UpdateNode(resp.AppName, resp.AppVersion, tokens[2])
		}
	}
}
func main_n() {
	for _, n := range ListNodes() {
		fmt.Println(n)
		api := giota.NewAPI(n, nil)
		resp, err := api.GetNeighbors()
		fmt.Println("done")
		fmt.Println(err)
		if err == nil {
			for _, nn := range resp.Neighbors {
				fmt.Println(nn.ConnectionType, nn.Address)
				InsertNode(string(nn.Address), nn.ConnectionType)
			}
		}
	}
}

func main22() {
	server := "http://iota.bitfinex.com:80"
	//server = "http://176.9.3.149:14265"
	server = "https://iotanode.us:443"
	api := giota.NewAPI(server, nil)
	resp, err := api.GetNeighbors()
	fmt.Println(err)
	if err == nil {
		for _, n := range resp.Neighbors {
			fmt.Println(n.ConnectionType, n.Address)
			InsertNode(string(n.Address), n.ConnectionType)
		}
	}
}

func main2() {
	server := "http://iota.bitfinex.com:80"
	server = "http://176.9.3.149:14265"
	api := giota.NewAPI(server, nil)
	resp, err := api.GetNodeInfo()
	if err == nil {
		unseen = append(unseen, string(resp.LatestSolidSubtangleMilestone))
	}

	x := 0
	for {
		if len(unseen) == 0 {
			break
		}
		//tofetch = get_and_remove_first(unseen)
		tofetch := unseen[0]
		fmt.Println(x, tofetch)
		x = x + 1
		unseen = unseen[1:len(unseen)]

		t, _ := giota.ToTrytes(tofetch)
		resp, err := api.GetTrytes([]giota.Trytes{t})
		if err == nil {
			seen = append(seen, tofetch)
			for _, tx := range resp.Trytes {
				//fmt.Println(i)
				fmt.Println(tx.AttachmentTimestamp, tx.Value, tx.Timestamp)
				//fmt.Println(tx.Timestamp)
				//fmt.Println(tx.AttachmentTimestamp)

				trunk := string(tx.TrunkTransaction)
				branch := string(tx.BranchTransaction)
				address := string(tx.Address)

				InsertTx(tofetch, tx.Timestamp, tx.Value, address)

				contains_seen := false
				for _, s := range seen {
					if s == trunk {
						contains_seen = true
						break
					}
				}
				contains_unseen := false
				for _, s := range unseen {
					if s == trunk {
						contains_unseen = true
						break
					}
				}
				if !contains_seen && !contains_unseen {
					unseen = append(unseen, trunk)
				}
				contains_seen = false
				for _, s := range seen {
					if s == branch {
						contains_seen = true
						break
					}
				}
				contains_unseen = false
				for _, s := range unseen {
					if s == branch {
						contains_unseen = true
						break
					}
				}
				if !contains_seen && !contains_unseen {
					unseen = append(unseen, branch)
				}
				//fmt.Println(len(seen), len(unseen))

			}
		}

	}

}*/
