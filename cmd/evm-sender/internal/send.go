package internal

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"github.com/staketab/evm-sender/cmd/evm-sender/internal/var"
	"github.com/staketab/evm-sender/cmd/evm-sender/pkg/func"
	"golang.org/x/crypto/sha3"
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

func CheckSendTx(logger *logrus.Logger) {
	config, err := ff.ReadConfigs()
	if err != nil {
		vars.ErrorLog.Fatal(err)
	}
	if config.Erc20.TokenContract != "" {
		SendErc20Tx(logger)
	} else {
		SendRangeTx(logger)
	}
}

func SendErc20Tx(logger *logrus.Logger) {
	config, err := ff.ReadConfigs()
	if err != nil {
		vars.ErrorLog.Fatal(err)
	}
	gasLimit := config.Default.GasLimit // in units
	toAddress := common.HexToAddress(config.Default.Recipient)
	tokenAddress := common.HexToAddress(config.Erc20.TokenContract) // assume we've added this to your config

	min := config.Default.Min
	max := config.Default.Max

	client, privateKey, nonce, gasPrice, chainID, err := ff.InitializeEthereumClientDefault(config)
	if err != nil {
		log.Fatal(err)
	}
	inTimeSeconds, err := strconv.Atoi(config.Default.InTime)
	if err != nil {
		log.Fatal(err)
	}
	symbol, _ := ff.GetTokenSymbol(config.Default.Rpc, tokenAddress)
	for {
		start := time.Now()
		logger.WithFields(logrus.Fields{"module": "send", "token": symbol, "count": config.Default.TxCount, "timer": inTimeSeconds}).Info("Starting to send a batch of transactions ")
		for i := 0; i < config.Default.TxCount; i++ {
			var value *big.Int
			if config.Default.Value.Cmp(big.NewInt(0)) != 0 {
				value = config.Default.Value
			} else {
				value = Rand(min, max)
			}

			transferFnSignature := []byte("transfer(address,uint256)")
			hash := sha3.NewLegacyKeccak256()
			hash.Write(transferFnSignature)
			methodID := hash.Sum(nil)[:4]
			paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
			paddedAmount := common.LeftPadBytes(value.Bytes(), 32)

			var data []byte
			data = append(data, methodID...)
			data = append(data, paddedAddress...)
			data = append(data, paddedAmount...)

			tx := types.NewTransaction(nonce, tokenAddress, big.NewInt(0), gasLimit, gasPrice, data)
			signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
			if err != nil {
				log.Fatal(err)
			}

			err = client.SendTransaction(context.Background(), signedTx)
			if err != nil {
				log.Fatal(err)
			}

			logger.WithFields(logrus.Fields{"module": "send", "token": symbol, "value": config.Default.Value, "hash": signedTx.Hash().Hex()}).Info("Tx sent")
			nonce++
		}
		elapsed := time.Since(start)
		sleepDuration := time.Duration(inTimeSeconds)*time.Second - elapsed
		if sleepDuration > 0 {
			time.Sleep(sleepDuration)
		}
	}
}

func SendRangeTx(logger *logrus.Logger) {
	config, err := ff.ReadConfigs()
	if err != nil {
		vars.ErrorLog.Fatal(err)
	}
	gasLimit := config.Default.GasLimit // in units
	toAddress := common.HexToAddress(config.Default.Recipient)
	memo := config.Default.Memo
	data := hex.EncodeToString([]byte(memo))

	min := config.Default.Min
	max := config.Default.Max

	client, privateKey, nonce, gasPrice, chainID, err := ff.InitializeEthereumClientDefault(config)
	if err != nil {
		log.Fatal(err)
	}
	inTimeSeconds, err := strconv.Atoi(config.Default.InTime)
	if err != nil {
		log.Fatal(err)
	}
	for {
		start := time.Now()
		logger.WithFields(logrus.Fields{"module": "send", "count": config.Default.TxCount, "timer": inTimeSeconds}).Info("Starting to send a batch of transactions ")
		for i := 0; i < config.Default.TxCount; i++ {
			var value *big.Int
			if config.Default.Value.Cmp(big.NewInt(0)) != 0 {
				value = config.Default.Value
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
		elapsed := time.Since(start)
		sleepDuration := time.Duration(inTimeSeconds)*time.Second - elapsed
		if sleepDuration > 0 {
			time.Sleep(sleepDuration)
		}
	}
}

func SendBackTx(logger *logrus.Logger) {
	config, err := ff.ReadConfigs()
	if err != nil {
		vars.ErrorLog.Fatal(err)
	}
	if config.SendBack.Enable {
		value := config.SendBack.Value       // in wei
		gasLimit := config.SendBack.GasLimit // in units
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
			start := time.Now()
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
			elapsed := time.Since(start)
			sleepDuration := time.Duration(inTimeSeconds)*time.Second - elapsed
			if sleepDuration > 0 {
				time.Sleep(sleepDuration)
			}
		}
	}
}
