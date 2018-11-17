package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

type Signature struct {
	r    big.Int
	s    big.Int
	hash []byte
}

func CreateKeypair() ecdsa.PrivateKey {
	result, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	return *result
}

func SignInput(input *txinput, key ecdsa.PrivateKey) {
	//hash := sha256.Sum256([]byte(strconv.Itoa(input.from.value)))
	hash := input.from.ComputeHashByte()
	r, s, err := ecdsa.Sign(rand.Reader, &key, hash[:])
	if err != nil {
		panic(err)
	}
	input.sig = Signature{r: *r, s: *s, hash:hash[:]}
}

func CheckInput(input txinput) bool {
	valid := ecdsa.Verify(&input.from.pubkey, input.from.ComputeHashByte(), &input.sig.r, &input.sig.s)
	return valid
}
