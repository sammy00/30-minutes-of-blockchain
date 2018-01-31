package ecdsa_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"golang.org/x/crypto/sha3"
)

func TestECDSA(t *testing.T) {
	// configure the parameters for curves
	curve := elliptic.P256()
	// generate the priv-pub key pairs
	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)

	msg := []byte("Hello World")
	// get the message digest, a.k.a. its digest
	digest := sha3.Sum256(msg)

	// sign the message to get the signature as (r,s)
	r, s, err := ecdsa.Sign(rand.Reader, privKey, digest[:])
	if nil != err {
		t.Fatal(err)
	}

	// verify the signature (r,s) on the digest
	// by the corresponding public key
	if !ecdsa.Verify(&privKey.PublicKey, digest[:], r, s) {
		t.Fatal("the signature should be verified")
	}
}
