package bitcoin

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/bitcoinsv/bsvd/bsvec"
	"github.com/bitcoinsv/bsvd/chaincfg"
	"github.com/bitcoinsv/bsvutil"
	"github.com/ejfhp/trail"
	"github.com/ejfhp/trail/trace"
	"golang.org/x/crypto/sha3"
)

func WIF() (string, error) {
	source := rand.NewSource(int64(time.Now().Minute()) * int64(time.Now().Second()) * int64(time.Now().Nanosecond()))
	r := rand.New(source)
	var seed []byte
	b := make([]byte, 8)
	for i := 0; i < 10; i++ {
		binary.LittleEndian.PutUint64(b, math.Float64bits(r.Float64()))
		seed = append(seed, b...)
		seed = hasher(15, seed)
	}
	priv, _ := bsvec.PrivKeyFromBytes(bsvec.S256(), seed)
	wif, err := bsvutil.NewWIF(priv, &chaincfg.MainNetParams, true)
	if err != nil {
		return "", fmt.Errorf("cannot generate WIF: %v", err)
	}
	return wif.String(), nil
}

func AddressOf(wifkey string) (string, error) {
	tr := trace.New().Source("keys.go", "", "AddressOf")
	w, err := bsvutil.DecodeWIF(wifkey)
	if err != nil {
		trail.Println(trace.Alert("cannot decode WIF").UTC().Add("wif", wifkey).Error(err).Append(tr))
		return "", fmt.Errorf("cannot decode WIF: %w", err)
	}
	_, err = bsvec.ParsePubKey(w.SerializePubKey(), bsvec.S256())
	if err != nil {
		trail.Println(trace.Alert("cannot parse").UTC().Add("wif", wifkey).Error(err).Append(tr))
		return "", err
	}
	// fmt.Printf("pubk: %s\n", string(pubk.SerializeCompressed()))
	// fmt.Printf("pubk ser: %s\n", string(w.SerializePubKey()))
	add, err := bsvutil.NewAddressPubKey(w.SerializePubKey(), &chaincfg.MainNetParams)
	if err != nil {
		trail.Println(trace.Alert("cannot generate address from WIF").UTC().Add("wif", wifkey).Error(err).Append(tr))
		return "", fmt.Errorf("cannot generate address from WIF: %w", err)
	}
	return add.EncodeAddress(), nil

}

func DecodeWIF(wifkey string) (*bsvec.PrivateKey, error) {
	t := trace.New().Source("keys.go", "", "DecodeWIF")
	wif, err := bsvutil.DecodeWIF(wifkey)
	if err != nil {
		trail.Println(trace.Alert("cannot decode WIF").UTC().Error(err).Append(t))
		return nil, fmt.Errorf("cannot decode WIF: %w", err)
	}
	priv := wif.PrivKey
	return priv, nil
}

func hasher(repeat int, hash []byte) []byte {
	var out [32]byte
	in := hash
	for i := 0; i < repeat; i++ {
		out = sha3.Sum256(in)
		copy(in, out[:])
	}
	return out[:]
}
