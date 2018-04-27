package main

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func main() {
	var num int32 = 1
	he := fmt.Sprintf("%x", num)
	fmt.Println("int to hex: ", he)
	heBt := []byte(he)
	fmt.Println("hex to bytes: ", heBt)
	//elliptic curve Secp256k1 y2 = x3 + 7
	// S256 returns a Curve which implements secp256k1.
	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), heBt)
	fmt.Println("privKey: ", privKey.D)
	fmt.Println("pubKey", pubKey)

	//SerializeCompressed serializes a public key in a 33-byte compressed format.
	serCompr := pubKey.SerializeCompressed()

	addrPubKey, err := btcutil.NewAddressPubKey(serCompr, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("addrPubKey: ", addrPubKey)
	// Hash160 calculates the hash ripemd160(sha256(b)).
	//	btcutil.Hash160()
	//fmt.Println("serCompr: ",serCompr)

}
