package main

import (
	"crypto/cipher"
	"encoding/hex"
	"fmt"

	"github.com/vapor/crypto/ed25519"
	"github.com/vapor/crypto/ed25519/chainkd"
	kbed "go.dedis.ch/kyber/sign/eddsa"
)

func main() {
	// xprvStr := "f2767279cd01ed8793808e0542a18958e1a2f3a6b6fe5328ec79596a022bc6f085951a98a631917563f86bb91db9159dd2969ff9d690fc12b250baff2b6f6a1d"

	// xprv := &chainkd.XPrv{}
	// if err := xprv.UnmarshalText([]byte(xprvStr)); err != nil {
	//  fmt.Println("err1")
	// }

	// fmt.Println(xprv)
	// fmt.Println(xprv.XPub())
	// toSign := []byte{1}
	// signed := xprv.Sign(toSign)
	// fmt.Println(hex.EncodeToString(signed))

	seed := []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
	}

	edPrv := ed25519.NewKeyFromSeed(seed)
	fmt.Println("edPrv", edPrv)
	fmt.Println("edPub", edPrv.Public())
	kdPrv := chainkd.RootXPrv(seed)
	fmt.Println("kdPrv", kdPrv)
	fmt.Println("kdPrv.XPub", kdPrv.XPub())
	fmt.Println("kdPrv.XPub.PublicKey", kdPrv.XPub().PublicKey())
	fmt.Println("kdPrv.ExpandedPrivateKey.Public", kdPrv.ExpandedPrivateKey().Public())

	s := "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f03a107bff3ce10be1d70dd18e74bc09967e4d6309ba50d5f1ddc8664125531b8"
	ss, err := hex.DecodeString(s)
	if err != nil {
		fmt.Println("err", err)
	}

	stream := ConstantStream(ss)
	edDSA := kbed.NewEdDSA(stream)
	// marshalled, err := edDSA.MarshalBinary()
	fmt.Println(edDSA.Secret)
	fmt.Println(edDSA.Public)
	// fmt.Println(edDSA)
	// fmt.Println(edDSA)

	// unmarshalled := &EdDSA{}
	// err = unmarshalled.UnmarshalBinary(marshalled)
	// assert.Nil(t, err)
	// assert.Equal(t, edDSA, unmarshalled)

}

// ConstantStream is a cipher.Stream which always returns
// the same value.
func ConstantStream(buff []byte) cipher.Stream {
	return &constantStream{buff}
}

type constantStream struct {
	seed []byte
}

// XORKexStream implements the cipher.Stream interface
func (cs *constantStream) XORKeyStream(dst, src []byte) {
	copy(dst, cs.seed)
}
