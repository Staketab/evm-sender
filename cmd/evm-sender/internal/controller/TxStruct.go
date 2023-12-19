package controller

type Config struct {
	Default DefaultConfig `toml:"DEFAULT"`
}

type DefaultConfig struct {
	Rpc        string `toml:"rpc"`
	PrivateKey string `toml:"private_key"`
	Recipient  string `toml:"recipient"`
	Value      int64  `toml:"value"`
	GasLimit   uint64 `toml:"gas_limit"`
	Memo       string `toml:"memo"`
	TxCount    int    `toml:"txcount"`
	Min        int64  `toml:"min"`
	Max        int64  `toml:"max"`
}
