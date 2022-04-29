package key

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ramin0x53/eth_bruteforce/config"
)

func ShaConvert(txt string) []byte {
	sha256 := sha256.Sum256([]byte(txt))
	return sha256[:]
}

func AddrGenerator(key []byte) common.Address {
	privateKey, err := crypto.ToECDSA(key)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return fromAddress
}

func GetBalance(addr common.Address) float64 {
	client, err := ethclient.Dial(config.Apikey)
	if err != nil {
		log.Fatal(err)
	}

	balance, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		log.Fatal(err)
	}

	balancefloat := new(big.Float).SetInt(balance)
	i := new(big.Float)
	i.SetString("1000000000000000000")
	f := new(big.Float).Quo(balancefloat, i)
	b, _ := f.Float64()

	return b
}
