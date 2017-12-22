package client

import "fmt"
import "net"
import "time"

//import "github.com/andrewarrow/simplecoin.life/sql"

func Connect(node map[string]string) {
	if node["connection_type"] == "udp" {
		id := node["id"]
		fmt.Println(id)

		conn, err := net.DialTimeout("udp", id, time.Second*3)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		fmt.Println("writing")
		conn.Write([]byte("{\"command\": \"getNodeInfo\"}"))

		buffer := make([]byte, 1)
		fmt.Println("reading")
		//conn.Read(buffer)
		fmt.Println(buffer)

	}
}
