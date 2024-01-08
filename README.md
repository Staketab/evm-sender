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
fixedValue = "0"
gas_limit = 22000
memo = "From Staketab with LOVE!"
txcount = 3
inTime = "60"
min = "1000000000000000000"
max = "2000000000000000000"

[SEND-BACK]
enable = false
private_key = ""
recipient = ""
fixedValue = "1000000000000000000"
gas_limit = 22000
memo = "From Staketab with LOVE!"
txCount = 1
inTime = "60"
```
Specify in the DEFAULT config:
- ETH endpoint RPC `rpc = "https://rpc:443"`
- Sender's private key `private_key`
- Recipient's address `recipient`
- Transaction count sent within a specific time(`inTime env`) `txcount = 3`
- Min Max send range or use `fixedValue != 0` to send fixed value (if `fixedValue = 0` the `min` `max` range will be used)

Specify in the SEND-BACK config:
- Enable it with `enable = true` if needed
- Sender's private key `private_key`
- Recipient's address `recipient`
- Transaction count sent within a specific time(`inTime env`) `txcount = 1`
- Fixed value to send `fixedValue`

## 2. Start:
```
evm-sender start
```