package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	"log"
)

var (
	buf bytes.Buffer
	num int32
)

func main() {
	connCfg := &rpcclient.ConnConfig{
		Host:         "127.0.0.1:8337",
		User:         "taketa",
		Pass:         "weuwdfyu7",
		HTTPPostMode: true,
		DisableTLS:   true,
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	for num = 1; num <= 3; num++ {
		binary.Write(&buf, binary.BigEndian, num)
		_, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), buf.Bytes())
		serCompr := pubKey.SerializeUncompressed()
		addrPubKey, err := btcutil.NewAddressPubKey(serCompr, &chaincfg.MainNetParams)

		if err != nil {
			fmt.Println(err)
		}
		acc, err := client.GetAccount(addrPubKey)
		if err != nil {
			fmt.Println("Acc error: ", err, acc)
		}
		bal, err := client.GetBalance(acc)
		if err != nil {
			fmt.Println("Bal error: ", err)
		}
		reseived, err := client.GetReceivedByAddress(addrPubKey)
		if err != nil {
			fmt.Println("Received error: ", err)
		}
		//if reseived>0{
		fmt.Printf("Address: %s; reseived: %d; balance: %d\n", addrPubKey.EncodeAddress(), reseived, bal)
		//}

		buf.Reset()
	}
}
