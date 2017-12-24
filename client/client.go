package client

import "fmt"
import "github.com/iotaledger/giota"
import "github.com/andrewarrow/iotame/sql"
import "strings"

func Connect(node map[string]string) {
	s := ""
	if strings.HasSuffix(node["id"], ":443") {
		s = "s"
	}
	url := fmt.Sprintf("http%s://%s", s, node["id"])
	fmt.Println(url)
	api := giota.NewAPI(url, nil)
	resp, err := api.GetNodeInfo()
	if err == nil {
		fmt.Println(resp.LatestSolidSubtangleMilestone)
		sql.UpdateNode(string(resp.LatestSolidSubtangleMilestone), resp.AppName, resp.AppVersion, node["id"])
	} else {
		fmt.Println(err)
	}
}
