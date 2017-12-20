package main

import "fmt"
import "github.com/iotaledger/giota"

func foo(t giota.Trytes, api *giota.API) {
	fmt.Println(t)
	resp, err := api.GetTrytes([]giota.Trytes{t})
	if err == nil {
		tx := resp.Trytes[0]
		fmt.Println("Address")
		fmt.Println(tx.Address)
		fmt.Println("Value")
		fmt.Println(tx.Value)
		fmt.Println("Timestamp")
		fmt.Println(tx.Timestamp)
		fmt.Println("TrunkTransaction")
		fmt.Println(tx.TrunkTransaction)
		fmt.Println("AttachmentTimestamp")
		fmt.Println(tx.AttachmentTimestamp)
		fmt.Println("Bundle")
		fmt.Println(tx.Bundle)
		foo(tx.TrunkTransaction, api)
	}
}

func main() {
	server := "http://iota.bitfinex.com:80"
	server = "http://176.9.3.149:14265"
	api := giota.NewAPI(server, nil)
	resp, err := api.GetNodeInfo()
	if err == nil {
		foo(resp.LatestMilestone, api)
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
