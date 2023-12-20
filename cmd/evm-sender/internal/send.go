package internal

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"github.com/staketab/evm-sender/cmd/evm-sender/internal/var"
	"github.com/staketab/evm-sender/cmd/evm-sender/pkg/func"
	"log"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

func Rand(min *big.Int, max *big.Int) *big.Int {
	rand.Seed(time.Now().UnixNano())
	diff := new(big.Int).Sub(max, min)
	randNum := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), diff)
	randNum.Add(randNum, min)
	return randNum
}

func SendRangeTx(logger *logrus.Logger) {
	config, err := ff.ReadConfigs()
	if err != nil {
		vars.ErrorLog.Fatal(err)
	}

	//value := big.NewInt(config.Default.Value) // in wei
	gasLimit := config.Default.GasLimit // in units
	toAddress := common.HexToAddress(config.Default.Recipient)
	memo := config.Default.Memo
	data := hex.EncodeToString([]byte(memo))

	min := big.NewInt(config.Default.Min)
	max := big.NewInt(config.Default.Max)

	client, privateKey, nonce, gasPrice, chainID, err := ff.InitializeEthereumClientDefault(config)
	if err != nil {
		log.Fatal(err)
	}
	inTimeSeconds, err := strconv.Atoi(config.Default.InTime)
	if err != nil {
		log.Fatal(err)
	}
	for {
		logger.WithFields(logrus.Fields{"module": "send", "count": config.Default.TxCount, "timer": inTimeSeconds}).Info("Starting to send a batch of transactions ")
		for i := 0; i < config.Default.TxCount; i++ {
			var value *big.Int
			if config.Default.Value != 0 {
				value = big.NewInt(config.Default.Value)
			} else {
				value = Rand(min, max)
			}
			tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, []byte(data))
			signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
			if err != nil {
				log.Fatal(err)
			}

			err = client.SendTransaction(context.Background(), signedTx)
			if err != nil {
				log.Fatal(err)
			}

			logger.WithFields(logrus.Fields{"module": "send", "value": config.Default.Value, "hash": signedTx.Hash().Hex()}).Info("Tx sent")
			nonce++
		}
		time.Sleep(time.Duration(inTimeSeconds) * time.Second)
	}
}

func SendBackTx(logger *logrus.Logger) {
	config, err := ff.ReadConfigs()
	if err != nil {
		vars.ErrorLog.Fatal(err)
	}
	if config.SendBack.Enable {
		value := big.NewInt(config.SendBack.Value) // in wei
		gasLimit := config.SendBack.GasLimit       // in units
		toAddress := common.HexToAddress(config.SendBack.Recipient)
		memo := config.SendBack.Memo
		data := hex.EncodeToString([]byte(memo))

		client, privateKey, nonce, gasPrice, chainID, err := ff.InitializeEthereumClientSendBack(config)
		if err != nil {
			log.Fatal(err)
		}
		inTimeSeconds, err := strconv.Atoi(config.SendBack.InTime)
		if err != nil {
			log.Fatal(err)
		}
		for {
			logger.WithFields(logrus.Fields{"module": "send-back", "count": config.SendBack.TxCount, "timer": inTimeSeconds}).Info("Starting to send the transaction back")
			for i := 0; i < config.SendBack.TxCount; i++ {
				tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, []byte(data))
				signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
				if err != nil {
					log.Fatal(err)
				}

				err = client.SendTransaction(context.Background(), signedTx)
				if err != nil {
					log.Fatal(err)
				}

				logger.WithFields(logrus.Fields{"module": "send-back", "value": config.SendBack.Value, "hash": signedTx.Hash().Hex()}).Info("Tx sent back")
				nonce++
			}
			time.Sleep(time.Duration(inTimeSeconds) * time.Second)
		}
	}
}
