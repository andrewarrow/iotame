package main

import "fmt"
import "github.com/iotaledger/giota"

var seen []string = []string{"999999999999999999999999999999999999999999999999999999999999999999999999999999999"}
var unseen []string = []string{}

func main() {
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
			for i, tx := range resp.Trytes {
				fmt.Println(i)
				//fmt.Println(tx.TrunkTransaction)
				//fmt.Println(tx.BranchTransaction)

				trunk := string(tx.TrunkTransaction)
				branch := string(tx.BranchTransaction)
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
				//fmt.Println(seen, unseen)

			}
		}

	}

}

/*
func findTx(bundle giota.Trytes, trunk giota.Trytes, branch giota.Trytes, address giota.Trytes, api *giota.API) {
	fmt.Println("bar", bundle)

	aa, _ := address.ToAddress()
	tt, _ := trunk.ToAddress()
	bb, _ := branch.ToAddress()

	ftr := &giota.FindTransactionsRequest{Approvees: []giota.Trytes{address},
		Addresses: []giota.Address{address, tt, bb},
		Bundles:   []giota.Trytes{bundle}}

	resp, err := api.FindTransactions(ftr)
	if err == nil {
		for i, h := range resp.Hashes {
			//foo(i, h, api)
			fmt.Println(i, h)
		}
	}
}*/

func loadMilestone(i int, t giota.Trytes, api *giota.API) {
	fmt.Println("loadMilestone", i, t)
	resp, err := api.GetTrytes([]giota.Trytes{t})
	if err == nil {
		for i, tx := range resp.Trytes {
			fmt.Println(i)
			fmt.Println(tx.Bundle)
			fmt.Println(tx.TrunkTransaction)
			fmt.Println(tx.BranchTransaction)
			fmt.Println(tx.Address)
			//findTx(tx.Bundle, tx.TrunkTransaction, tx.BranchTransaction, tx.Address, api)
		}
		//tx := resp.Trytes[0]
		//fmt.Println("Address")
		//fmt.Println(tx.Address)
		//fmt.Println("Value")
		//fmt.Println(tx.Value)
		//fmt.Println("Timestamp")
		//fmt.Println(tx.Timestamp)
		//fmt.Println("TrunkTransaction")
		//fmt.Println(tx.TrunkTransaction)
		//fmt.Println("AttachmentTimestamp")
		//fmt.Println(tx.AttachmentTimestamp)
		//fmt.Println("Bundle")
		//fmt.Println(tx.Bundle)
		//fmt.Println("Tag")
		//fmt.Println(tx.Tag)
		//fmt.Println("CurrentIndex")
		//fmt.Println(tx.CurrentIndex)
		//fmt.Println("LastIndex")
		//fmt.Println(tx.LastIndex)
		//foo(0, tx.TrunkTransaction, api)
		//foo(0, tx.BranchTransaction, api)
		//bar(tx.Bundle, api)
	}
}

func main3() {
	tritsFrom := []int8{-1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, -1, 1}
	trits, _ := giota.ToTrits(tritsFrom)

	trytes := trits.Trytes()
	//fmt.Println(trytes)
	//trytesFrom := "ABCDEAAC9ACB9PO..."
	//trytes2, _ := giota.ToTrytes(trytesFrom)

	//hash := trytes.Hash()

	//api := giota.NewAPI("http://localhost:14265", nil)
	//resp, err := api.FindTransactions([]giota.Trytes{"DEXRPL...SJRU"})

	index := 0
	security := 2
	adr, _ := giota.NewAddress(trytes, index, security) //without checksum.
	//fmt.Println(err)
	fmt.Println(adr)
	//adrWithChecksum := adr.WithChecksum() //adrWithChecksum is trytes type.
	//fmt.Println(adrWithChecksum)

	/*
		tx, err := giota.NewTransaction(trytes)
		mwm := 15
		if tx.HasValidNonce(mwm) {
		}
		trytes2 := tx.trytes()

		key := giota.NewKey(seed, index, security)
		norm := bundleHash.Normalize()
		sign := giota.Sign(norm[:27], key[:6561/3])

		if giota.ValidateSig(adr, []giota.Trytes{sign}, bundleHash) {
		}

		trs := []giota.Transfer{
			giota.Transfer{
				Address: "KTXF...QTIWOWTY",
				Value:   20,
				Tag:     "MOUDAMEPO",
			},
		}
		_, pow := giota.GetBestPoW()
		bdl, err = giota.Send(api, seed, security, trs, mwm, pow)
	*/
}
