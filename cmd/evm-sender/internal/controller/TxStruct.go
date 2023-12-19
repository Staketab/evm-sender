package controller

type Config struct {
	Default  DefaultConfig  `toml:"DEFAULT"`
	SendBack SendBackConfig `toml:"SEND-BACK"`
}

type DefaultConfig struct {
	Rpc        string `toml:"rpc"`
	PrivateKey string `toml:"private_key"`
	Recipient  string `toml:"recipient"`
	Value      int64  `toml:"fixedValue"`
	GasLimit   uint64 `toml:"gas_limit"`
	Memo       string `toml:"memo"`
	TxCount    int    `toml:"txCount"`
	InTime     string `toml:"inTime"`
	Min        int64  `toml:"min"`
	Max        int64  `toml:"max"`
}

type SendBackConfig struct {
	Enable     bool   `toml:"enable"`
	PrivateKey string `toml:"private_key"`
	Recipient  string `toml:"recipient"`
	Value      int64  `toml:"fixedValue"`
	GasLimit   uint64 `toml:"gas_limit"`
	Memo       string `toml:"memo"`
	TxCount    int    `toml:"txCount"`
	InTime     string `toml:"inTime"`
}
