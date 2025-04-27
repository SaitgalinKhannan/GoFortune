# Генерирует ABI

```shell
solc --abi contracts/SingleLottery.sol -o contracts/abi
```

# Генерирует биндинги для Go

```shell
$ abigen --abi contracts/abi/SingleLottery.abi --pkg main --type SingleLottery --out SingleLottery.go
```

# Compiled bytecode

```shell
$ solc --bin contracts/SingleLottery.sol -o contracts/bin
```
abigen --abi contracts/abi/SingleLottery.abi --pkg main --type SingleLottery --out SingleLottery.go --bin contracts/bin/SingleLottery.bin

# Генерирует биндинги с деплоем для Go

```shell
$ abigen --abi contracts/abi/SingleLottery.abi --pkg main --type SingleLottery --out SingleLottery.go --bin contracts/bin/SingleLottery.bin
```