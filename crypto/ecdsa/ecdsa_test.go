package ecdsa_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/asn1"
	"math/big"
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
		t.Fatal("the signature should be valid")
	}
}

type ecdsaSig struct {
	R, S *big.Int
}

func TestECDSAWithASN1(t *testing.T) {
	// configure the parameters for curves
	curve := elliptic.P256()
	// generate the priv-pub key pairs
	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)

	msg := []byte("Hello World")
	// get the message digest, a.k.a. its digest
	digest := sha3.Sum256(msg)

	// sign the message with the key and
	// get the signature encoded according to ASN1
	asn1Sig, err := privKey.Sign(rand.Reader, digest[:], nil)
	if nil != err {
		t.Fatal(err)
	}

	// unmarshal the signature for verification
	var decodedSig ecdsaSig
	if _, err := asn1.Unmarshal(asn1Sig, &decodedSig); nil != err {
		t.Fatal(err)
	}

	// verify the signature (r,s) on the digest
	// by the corresponding public key
	if !ecdsa.Verify(&privKey.PublicKey, digest[:], decodedSig.R, decodedSig.S) {
		t.Fatal("the signature should be valid")
	}
}
