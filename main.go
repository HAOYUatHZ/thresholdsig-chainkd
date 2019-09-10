package main

import (
	"encoding/hex"
	"fmt"

	"github.com/vapor/crypto/ed25519/chainkd"
)

func main() {
	xprvStr := "f2767279cd01ed8793808e0542a18958e1a2f3a6b6fe5328ec79596a022bc6f085951a98a631917563f86bb91db9159dd2969ff9d690fc12b250baff2b6f6a1d"

	xprv := &chainkd.XPrv{}
	if err := xprv.UnmarshalText([]byte(xprvStr)); err != nil {
		fmt.Println("err1")
	}

	fmt.Println(xprv)
	toSign := []byte{1}
	signed := xprv.Sign(toSign)
	fmt.Println(hex.EncodeToString(signed))
}
