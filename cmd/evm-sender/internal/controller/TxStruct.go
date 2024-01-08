package controller

import "math/big"

type Config struct {
	Default  DefaultConfig  `toml:"DEFAULT"`
	Erc20    Erc20Config    `toml:"ERC20"`
	SendBack SendBackConfig `toml:"SEND-BACK"`
}

type DefaultConfig struct {
	Rpc        string   `toml:"rpc"`
	PrivateKey string   `toml:"private_key"`
	Recipient  string   `toml:"recipient"`
	Value      *big.Int `toml:"fixedValue"`
	GasLimit   uint64   `toml:"gas_limit"`
	Memo       string   `toml:"memo"`
	TxCount    int      `toml:"txCount"`
	InTime     string   `toml:"inTime"`
	Min        *big.Int `toml:"min"`
	Max        *big.Int `toml:"max"`
}

type Erc20Config struct {
	TokenContract string `toml:"tokenContract"`
}

type SendBackConfig struct {
	Enable     bool     `toml:"enable"`
	PrivateKey string   `toml:"private_key"`
	Recipient  string   `toml:"recipient"`
	Value      *big.Int `toml:"fixedValue"`
	GasLimit   uint64   `toml:"gas_limit"`
	Memo       string   `toml:"memo"`
	TxCount    int      `toml:"txCount"`
	InTime     string   `toml:"inTime"`
}
