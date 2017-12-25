package client

import "fmt"
import "github.com/iotaledger/giota"
import "github.com/andrewarrow/iotame/sql"

func Connect(node map[string]string) {
	url := node["url"]
	fmt.Println(url)
	api := giota.NewAPI(url, nil)
	resp, err := api.GetNodeInfo()
	if err == nil {
		fmt.Println(resp.LatestSolidSubtangleMilestone)
		sql.UpdateNode(resp.LatestSolidSubtangleMilestoneIndex, string(resp.LatestSolidSubtangleMilestone), resp.AppName, resp.AppVersion, node["id"])
	} else {
		fmt.Println(err)
	}
}
