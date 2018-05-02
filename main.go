package main

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"net/http"
	"encoding/json"
)
type AddrInfo struct {
	Address string
	N_tx int
	Total_sent int
	Final_balance int
}
var (
	buf bytes.Buffer
	num int32
	addrInfo AddrInfo
)
func main() {
	for num=1;num<=100;num++{
		binary.Write(&buf,binary.BigEndian,num)
		_, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), buf.Bytes())
		serCompr := pubKey.SerializeUncompressed()
		addrPubKey, err := btcutil.NewAddressPubKey(serCompr, &chaincfg.MainNetParams)
		if err != nil {
			fmt.Println(err)
		}
		resp, err := http.Get("https://blockchain.info/rawaddr/"+addrPubKey.EncodeAddress())
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &addrInfo)
		if err != nil {
			fmt.Println("error:", err)
		}
		if addrInfo.N_tx>0{
			fmt.Printf("PrivKey: %d; Addr: %s; n_tx: %d; total_sent: %d; ballance: %d\n", num, addrInfo.Address,addrInfo.N_tx,addrInfo.Total_sent, addrInfo.Final_balance)

		}
		buf.Reset()
	}


}
