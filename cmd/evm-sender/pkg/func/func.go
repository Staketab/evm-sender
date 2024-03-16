package ff

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pelletier/go-toml"
	"github.com/staketab/evm-sender/cmd/evm-sender/internal/controller"
	"github.com/staketab/evm-sender/cmd/evm-sender/internal/var"
	"math/big"
	"os"
	"os/user"
	"path/filepath"
)

func CreateDirectory(path string) error {
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to controller current user: %s", err)
	}

	fullPath := filepath.Join(usr.HomeDir, path)

	err = os.MkdirAll(fullPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to cr directory: %s", err)
	}

	vars.InfoLog.Printf("Directory created: %s\n", fullPath)
	return nil
}

func CreateConfigFile(path string) error {
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to controller current user: %s", err)
	}
	filePath := filepath.Join(usr.HomeDir, path)
	content := []byte(`[DEFAULT]
rpc = ""
private_key = ""
recipient = ""
fixedValue = "0"
gas_limit = 22000
memo = "From Staketab with LOVE!"
txCount = 3
inTime = "60"
min = "1000000000000000000"
max = "9000000000000000000"

[ERC20]
tokenContract = ""

[SEND-BACK]
enable = false
private_key = ""
recipient = ""
fixedValue = "1000000000000000000"
gas_limit = 22000
memo = "From Staketab with LOVE!"
txCount = 1
inTime = "60"
`)

	err2 := os.WriteFile(filePath, content, 0644)
	if err2 != nil {
		return fmt.Errorf("failed to cr config file: %s", err)
	}

	vars.InfoLog.Printf("Config file created: %s\n", filePath)
	return nil
}

func ReadConfigs() (controller.Config, error) {
	usr, err := user.Current()
	if err != nil {
		return controller.Config{}, fmt.Errorf("Failed to controller current user: %s", err)
	}
	filePath := filepath.Join(usr.HomeDir, vars.ConfigFilePath)
	config := controller.Config{}
	tomlFile, err := toml.LoadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("failed to load config file: %s", err)
	}

	err = tomlFile.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal config: %s", err)
	}

	return config, nil
}

func InitializeEthereumClientDefault(config controller.Config) (*ethclient.Client, *ecdsa.PrivateKey, uint64, *big.Int, *big.Int, error) {
	client, err := ethclient.Dial(config.Default.Rpc)
	if err != nil {
		return nil, nil, 0, nil, nil, err
	}

	privateKey, err := crypto.HexToECDSA(config.Default.PrivateKey)
	if err != nil {
		return nil, nil, 0, nil, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, 0, nil, nil, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, nil, 0, nil, nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, nil, 0, nil, nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, nil, 0, nil, nil, err
	}

	return client, privateKey, nonce, gasPrice, chainID, nil
}

func InitializeEthereumClientSendBack(config controller.Config) (*ethclient.Client, *ecdsa.PrivateKey, uint64, *big.Int, *big.Int, error) {
	client, err := ethclient.Dial(config.Default.Rpc)
	if err != nil {
		return nil, nil, 0, nil, nil, err
	}

	privateKey, err := crypto.HexToECDSA(config.SendBack.PrivateKey)
	if err != nil {
		return nil, nil, 0, nil, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, 0, nil, nil, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, nil, 0, nil, nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, nil, 0, nil, nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, nil, 0, nil, nil, err
	}

	return client, privateKey, nonce, gasPrice, chainID, nil
}

func UnpackString(data []byte) (string, error) {
	if len(data) < 96 {
		return "", fmt.Errorf("data too short to contain a string")
	}
	offset := new(big.Int).SetBytes(data[0:32]).Uint64()
	if offset+32 > uint64(len(data)) {
		return "", fmt.Errorf("incorrect offset or string length")
	}
	strLen := new(big.Int).SetBytes(data[offset : offset+32]).Uint64()
	if offset+32+strLen > uint64(len(data)) {
		return "", fmt.Errorf("incorrect string length")
	}
	strData := data[offset+32 : offset+32+strLen]

	return string(strData), nil
}

func GetTokenSymbol(RPCURL string, tokenAddress common.Address) (string, error) {
	client, err := ethclient.Dial(RPCURL)
	if err != nil {
		return "", fmt.Errorf("failed to connect to the Ethereum node: %v", err)
	}
	data := common.Hex2Bytes("06fdde03")
	msg := ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}
	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return "", fmt.Errorf("failed to call the contract: %v", err)
	}
	symbol, err := UnpackString(result)
	if err != nil {
		return "", fmt.Errorf("failed to decode the token symbol: %v", err)
	}
	return symbol, nil
}
