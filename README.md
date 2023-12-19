# EVM sender
EVM sender

## Clone the repository:
```
git clone https://github.com/staketab/evm-sender.git
```

## Build the binary:
```
cd evm-sender
make install
```

## USAGE:
```
ETH Sender

Usage:
  evm-sender [flags]
  evm-sender [command]

Available Commands:
  help        Help about any command
  init        Initialize config
  start       Start sender
  version     Print the CLI version

Flags:
  -h, --help   help for evm-sender
```

## 1. Init config:
This command will create a config along the path `/home/USER/.evm-sender/config.toml`
```
evm-sender init
```
Example:
```
[DEFAULT]
rpc = "https://rpc:443"
private_key = ""
recipient = ""
value = 1000000
gas_limit = 22000
memo = "From Staketab with LOVE!"
txcount = 3
min = 1000000000000000000
max = 2000000000000000000
```
Specify the in the config:
- ETH endpoint RPC `rpc = "https://rpc:443"`
- Sender `private_key`
- Recepient address `recipient`
- Tx count sent in 1 minute `txcount = 3`
- Min Max send range
- 
## 2. Start:
```
evm-sender start
```