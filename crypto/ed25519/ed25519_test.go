package ed25519_test

//import "github.com/agl/ed25519"
import (
	"crypto/rand"
	"testing"

	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/sha3"
)

func TestEd25519(t *testing.T) {
	// generate a priv-pub key pairs randomly
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	// check for error if any
	if nil != err {
		t.Fatal("ed25519.GenerateKey: ", err)
	}

	rawMsg := []byte("Hello World")
	// get the hash of the raw message, a.k.a its digest
	digest := sha3.Sum256(rawMsg)

	// note: the length of keys and signature in ed25519 is fixed as
	// + ed25519.PublicKeySize(32 bytes) for public keys
	// + ed25519.PrivateKeySize(64 bytes) for private keys
	// + ed25519.SignatureSize(64 bytes) for signatures
	if len(privKey) != ed25519.PrivateKeySize {
		t.Fatal("invalid length of private key")
	}
	sig := ed25519.Sign(privKey, digest[:])

	if !ed25519.Verify(pubKey, digest[:], sig) {
		t.Fatal("ed25519.Verify: failed")
	}
}
