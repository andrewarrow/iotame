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
