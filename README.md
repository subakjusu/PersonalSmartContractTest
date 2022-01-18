# PersonalSmartContractTest
Example smart contract deployed in go for posterity. Referenced: https://towardsdev.com/creating-a-simple-ethereum-smart-contract-in-golang-138b9439f64e

# Commands to set before executing script
go env -w GO111MODULE=on
go mod init
go mod tidy

# Outputs
## Smart contract deployment:
ubuntu@ip-172-31-8-253:~/go_projects/src/github.com/testSmartContract$ go run deploy.go
0x549129CECd93C366a4Fb327970CaA8604A54dbb2

## Ganache output:
Available Accounts
==================
(0) 0x45E22A133f580F0c801AdebB5B7ef0A163d7f734 (1000 ETH)
(1) 0xe81d2757e1f16A7Ed87B2b64d45d4f1Ee1Eb1459 (1000 ETH)
(2) 0x7C8898980DDcE2794B30EC1b0F329F2938aA1d8C (1000 ETH)
(3) 0x963029Ac43286Bc8B42A11Eb737e6855533a251e (1000 ETH)
(4) 0x07Db759a9F5A3776FD18262f1e019FA375eEA3FA (1000 ETH)
(5) 0x52913a33176DF62E60454503Ffb2a781358770DE (1000 ETH)
(6) 0x5f11D3deF2099941193f41fbeDc9dB7d2E12a437 (1000 ETH)
(7) 0xD80C617a2e8f68051583d5418AC30bFb563DDC26 (1000 ETH)
(8) 0x9AA103907B4511C5B01aD57ac803e12193a1145D (1000 ETH)
(9) 0xB05cd9735313D5F4EA82586A2e7F513bD04Ae8E1 (1000 ETH)

HD Wallet
==================
Mnemonic:      thrive fork bless involve theory edge always fine please exile dish knock
Base HD Path:  m/44'/60'/0'/0/{account_index}

Default Gas Price
==================
2000000000

BlockGas Limit
==================
30000000

Call Gas Limit
==================
50000000

Chain Id
==================
1337

RPC Listening on 127.0.0.1:8545
eth_getTransactionCount
eth_chainId
eth_sendRawTransaction
eth_getTransactionCount
eth_chainId
eth_sendRawTransaction

  Transaction: 0x56bba91ac808cc75e9757761c5911bfa315c3b3126dac94bef26328db05e24a0
  Contract created: 0x549129cecd93c366a4fb327970caa8604a54dbb2
  Gas usage: 154969
  Block number: 2
  Block time: Tue Jan 18 2022 18:44:21 GMT+0000 (Coordinated Universal Time)

eth_call
eth_call
eth_call
eth_call

## Echo output:
ubuntu@ip-172-31-8-253:~/go_projects/src/github.com/testSmartContract$ go run main.go

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.2.1
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323
{"time":"2022-01-18T19:05:02.583785271Z","id":"","remote_ip":"127.0.0.1","host":"localhost:1323","method":"GET","uri":"/hello","user_agent":"curl/7.68.0","status":200,"error":"","latency":14233258,"latency_human":"14.233258ms","bytes_in":0,"bytes_out":14}
{"time":"2022-01-18T19:05:15.31049746Z","id":"","remote_ip":"127.0.0.1","host":"localhost:1323","method":"GET","uri":"/greet/owenyuwono","user_agent":"curl/7.68.0","status":200,"error":"","latency":10238370,"latency_human":"10.23837ms","bytes_in":0,"bytes_out":13}
{"time":"2022-01-18T19:05:20.723742891Z","id":"","remote_ip":"127.0.0.1","host":"localhost:1323","method":"GET","uri":"/greet/greetings","user_agent":"curl/7.68.0","status":200,"error":"","latency":9323748,"latency_human":"9.323748ms","bytes_in":0,"bytes_out":12}
{"time":"2022-01-18T19:05:24.787828271Z","id":"","remote_ip":"127.0.0.1","host":"localhost:1323","method":"GET","uri":"/greet/randomtextforthequeryparameter","user_agent":"curl/7.68.0","status":200,"error":"","latency":9107618,"latency_human":"9.107618ms","bytes_in":0,"bytes_out":33}
