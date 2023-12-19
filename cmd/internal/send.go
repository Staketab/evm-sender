package internal

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	vars "github.com/staketab/evm-sender/cmd/internal/var"
	ff "github.com/staketab/evm-sender/cmd/pkg/func"
	"log"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Rand(min *big.Int, max *big.Int) *big.Int {
	rand.Seed(time.Now().UnixNano())
	diff := new(big.Int).Sub(max, min)
	randNum := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), diff)
	randNum.Add(randNum, min)
	return randNum
}

func SendTx() {
	config, err := ff.ReadConfigs()
	if err != nil {
		vars.ErrorLog.Fatal(err)
	}

	client, err := ethclient.Dial(config.Default.Rpc)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(config.Default.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//value := big.NewInt(config.Default.Value) // in wei
	gasLimit := config.Default.GasLimit // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(config.Default.Recipient)

	memo := config.Default.Memo
	data := hex.EncodeToString([]byte(memo))

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	min := big.NewInt(config.Default.Min)
	max := big.NewInt(config.Default.Max)

	for {
		for i := 0; i < config.Default.TxCount; i++ {
			value := Rand(min, max)
			tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, []byte(data))
			signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
			if err != nil {
				log.Fatal(err)
			}

			err = client.SendTransaction(context.Background(), signedTx)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Tx sent: %s\n", signedTx.Hash().Hex())
			nonce++
		}

		time.Sleep(60 * time.Second)
	}
}
