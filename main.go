package main

import (
	"GoFortune/utils"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

//var gasLimit uint64 = 300000

func main() {
	utils.LoadEnv()
	privateKey := utils.MustGetEnv("PRIVATE_KEY")
	fmt.Printf("private key: %s\n", privateKey)
	rpcUrl := utils.MustGetEnv("RPC_URL")
	fmt.Printf("RPC URL: %s\n", rpcUrl)
	contractAddr := utils.MustGetEnv("CONTRACT_ADDR")
	fmt.Printf("Contract address: %s\n", contractAddr)

	// Подключение к BNB testnet
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalf("Не удалось подключиться к BNB testnet: %v", err)
	}
	fmt.Println("Успешно подключились к BNB testnet")

	// Убираем "0x" из privateKey, если он есть
	if len(privateKey) > 2 && privateKey[:2] == "0x" {
		privateKey = privateKey[2:]
	}

	// Создаем аккаунты для тестирования с обработкой ошибок
	ownerKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("Ошибка при создании ownerKey: %v", err)
	}
	user1Key, err := crypto.HexToECDSA("ec7477e2254a54e422227cc72c6232a7da34f5b93c79b9235e920f64d4acda2d")
	if err != nil {
		log.Fatalf("Ошибка при создании user1Key: %v", err)
	}
	user2Key, err := crypto.HexToECDSA("dda4c2cad31eb06f7929ab7a83da619a9887d330a76fd7f62e48e0e797955ed1")
	if err != nil {
		log.Fatalf("Ошибка при создании user2Key: %v", err)
	}

	// Получаем адреса из приватных ключей
	ownerAddress := crypto.PubkeyToAddress(ownerKey.PublicKey)
	user1Address := crypto.PubkeyToAddress(user1Key.PublicKey)
	user2Address := crypto.PubkeyToAddress(user2Key.PublicKey)

	fmt.Printf("Адрес владельца: %s\n", ownerAddress.Hex())
	fmt.Printf("Адрес пользователя 1: %s\n", user1Address.Hex())
	fmt.Printf("Адрес пользователя 2: %s\n", user2Address.Hex())

	// Создаем инстанс контракта
	contractAddress := common.HexToAddress(contractAddr)
	instance, err := NewSingleLottery(contractAddress, client)
	if err != nil {
		log.Fatalf("Не удалось создать экземпляр контракта: %v", err)
	}

	// Вызываем различные сценарии тестирования
	testLotteryInfo(client, instance, contractAddress)
	testBuyTickets(client, instance, user1Key, user2Key)
	testDrawLottery(client, instance, ownerKey)
	testClaimRewards(client, instance, ownerKey, user1Key, user2Key)
}

// Получение информации о лотерее
func testLotteryInfo(client *ethclient.Client, instance *SingleLottery, contractAddress common.Address) {
	ctx := context.Background()

	ticketPrice, err := instance.TICKETPRICE(nil)
	if err != nil {
		log.Printf("Ошибка при получении цены билета: %v", err)
		return
	}

	maxTickets, err := instance.MAXTICKETS(nil)
	if err != nil {
		log.Printf("Ошибка при получении максимального количества билетов: %v", err)
		return
	}

	ticketsSold, err := instance.TicketsSold(nil)
	if err != nil {
		log.Printf("Ошибка при получении количества проданных билетов: %v", err)
		return
	}

	owner, err := instance.Owner(nil)
	if err != nil {
		log.Printf("Ошибка при получении владельца: %v", err)
		return
	}

	balance, err := client.BalanceAt(ctx, contractAddress, nil)
	if err != nil {
		log.Printf("Ошибка при получении баланса контракта: %v", err)
		return
	}

	fmt.Println("\n=== Информация о лотерее ===")
	fmt.Printf("Адрес контракта: %s\n", contractAddress.Hex())
	fmt.Printf("Владелец: %s\n", owner.Hex())
	fmt.Printf("Цена билета: %s BNB\n", weiToBNB(ticketPrice))
	fmt.Printf("Максимум билетов: %s\n", maxTickets.String())
	fmt.Printf("Продано билетов: %s\n", ticketsSold.String())
	fmt.Printf("Баланс контракта: %s BNB\n", weiToBNB(balance))

	// Получение оставшихся билетов
	remainingTickets, err := instance.GetRemainingTickets(nil)
	if err != nil {
		log.Printf("Ошибка при получении оставшихся билетов: %v", err)
		return
	}
	fmt.Printf("Осталось билетов: %s\n", remainingTickets.String())
}

// Тест покупки билетов
func testBuyTickets(client *ethclient.Client, instance *SingleLottery, user1Key, user2Key *ecdsa.PrivateKey) {
	ctx := context.Background()

	// Получаем цену билета
	ticketPrice, err := instance.TICKETPRICE(nil)
	if err != nil {
		log.Printf("Ошибка при получении цены билета: %v", err)
		return
	}

	fmt.Println("\n=== Тестирование покупки билетов ===")

	// Покупка билетов от пользователя 1
	user1Addr := crypto.PubkeyToAddress(user1Key.PublicKey)
	user1Nonce, err := client.PendingNonceAt(ctx, user1Addr)
	if err != nil {
		log.Printf("Ошибка при получении nonce для user1: %v", err)
		return
	}

	// Вычисляем стоимость 3-х билетов
	ticketsCost := new(big.Int).Mul(ticketPrice, big.NewInt(3))

	// Создаем транзакцию для покупки билетов
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Printf("Ошибка при получении chain ID: %v", err)
		return
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("Ошибка при получении gasPrice: %v", err)
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(user1Key, chainID)
	if err != nil {
		log.Printf("Ошибка при создании транзактора: %v", err)
		return
	}
	auth.Nonce = big.NewInt(int64(user1Nonce))
	auth.Value = ticketsCost // Отправляем стоимость билетов
	//auth.GasLimit = gasLimit // Устанавливаем лимит газа
	auth.GasPrice = gasPrice // Используем рекомендуемую цену газа

	// Покупаем 3 билета
	fmt.Printf("Пользователь 1 покупает 3 билета за %s BNB...\n", weiToBNB(ticketsCost))
	tx, err := instance.BuyTickets(auth, big.NewInt(3))
	if err != nil {
		log.Printf("Ошибка при покупке билетов: %v", err)
		return
	}
	fmt.Printf("Транзакция отправлена: %s\n", tx.Hash().Hex())

	// Ждем подтверждения транзакции
	receipt, err := waitForTxReceipt(client, tx.Hash())
	if err != nil {
		log.Printf("Ошибка при ожидании подтверждения транзакции: %v", err)
		return
	}
	fmt.Printf("Транзакция подтверждена в блоке %d\n", receipt.BlockNumber)

	// Проверяем количество купленных билетов
	ticketsSold, _ := instance.TicketsSold(nil)
	fmt.Printf("Всего продано билетов: %s\n", ticketsSold.String())

	// Покупка билетов от пользователя 2
	user2Addr := crypto.PubkeyToAddress(user2Key.PublicKey)
	user2Nonce, err := client.PendingNonceAt(ctx, user2Addr)
	if err != nil {
		log.Printf("Ошибка при получении nonce для user2: %v", err)
		return
	}

	// Вычисляем стоимость 2-х билетов
	ticketsCost = new(big.Int).Mul(ticketPrice, big.NewInt(2))

	auth, err = bind.NewKeyedTransactorWithChainID(user2Key, chainID)
	if err != nil {
		log.Printf("Ошибка при создании транзактора: %v", err)
		return
	}
	auth.Nonce = big.NewInt(int64(user2Nonce))
	auth.Value = ticketsCost
	//auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	// Покупаем 2 билета
	fmt.Printf("Пользователь 2 покупает 2 билета за %s BNB...\n", weiToBNB(ticketsCost))
	tx, err = instance.BuyTickets(auth, big.NewInt(2))
	if err != nil {
		log.Printf("Ошибка при покупке билетов: %v", err)
		return
	}
	fmt.Printf("Транзакция отправлена: %s\n", tx.Hash().Hex())

	// Ждем подтверждения транзакции
	receipt, err = waitForTxReceipt(client, tx.Hash())
	if err != nil {
		log.Printf("Ошибка при ожидании подтверждения транзакции: %v", err)
		return
	}
	fmt.Printf("Транзакция подтверждена в блоке %d\n", receipt.BlockNumber)

	// Проверяем количество купленных билетов
	ticketsSold, _ = instance.TicketsSold(nil)
	fmt.Printf("Всего продано билетов: %s\n", ticketsSold.String())

	// Получаем информацию о пользователях
	user1Info, err := instance.GetMyInfo(&bind.CallOpts{From: user1Addr})
	if err != nil {
		log.Printf("Ошибка при получении информации о пользователе 1: %v", err)
	} else {
		fmt.Printf("Пользователь 1: билетов = %s, ожидающих наград = %s BNB\n", user1Info.Tickets.String(), weiToBNB(user1Info.Rewards))
	}

	user2Info, err := instance.GetMyInfo(&bind.CallOpts{From: user2Addr})
	if err != nil {
		log.Printf("Ошибка при получении информации о пользователе 1: %v", err)
	} else {
		fmt.Printf("Пользователь 2: билетов = %s, ожидающих наград = %s BNB\n", user2Info.Tickets.String(), weiToBNB(user2Info.Rewards))
	}
}

// Тест проведения розыгрыша
func testDrawLottery(client *ethclient.Client, instance *SingleLottery, ownerKey *ecdsa.PrivateKey) {
	ctx := context.Background()

	fmt.Println("\n=== Тестирование проведения розыгрыша ===")

	// Подготавливаем транзакцию
	ownerAddr := crypto.PubkeyToAddress(ownerKey.PublicKey)
	nonce, err := client.PendingNonceAt(ctx, ownerAddr)
	if err != nil {
		log.Printf("Ошибка при получении nonce: %v", err)
		return
	}

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Printf("Ошибка при получении chain ID: %v", err)
		return
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("Ошибка при получении gasPrice: %v", err)
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(ownerKey, chainID)
	if err != nil {
		log.Printf("Ошибка при создании транзактора: %v", err)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // Не отправляем эфир
	//auth.GasLimit = gasLimit   // Устанавливаем лимит газа
	auth.GasPrice = gasPrice // Используем рекомендуемую цену газа

	// Случайное число для розыгрыша
	randomSeed := big.NewInt(time.Now().UnixNano())
	fmt.Printf("Проведение розыгрыша с seed = %s...\n", randomSeed.String())

	// Вызываем функцию drawLottery
	tx, err := instance.DrawLottery(auth, randomSeed)
	if err != nil {
		log.Printf("Ошибка при проведении розыгрыша: %v", err)
		return
	}
	fmt.Printf("Транзакция отправлена: %s\n", tx.Hash().Hex())

	// Ждем подтверждения транзакции
	receipt, err := waitForTxReceipt(client, tx.Hash())
	if err != nil {
		log.Printf("Ошибка при ожидании подтверждения транзакции: %v", err)
		return
	}
	fmt.Printf("Транзакция подтверждена в блоке %d\n", receipt.BlockNumber)

	// Проверяем, что лотерея завершена
	isFinished, err := instance.LotteryFinished(nil)
	if err != nil {
		log.Printf("Ошибка при проверке статуса лотереи: %v", err)
		return
	}
	fmt.Printf("Лотерея завершена: %t\n", isFinished)

	// Получаем победителя и призовой фонд
	winner, err := instance.Winner(nil)
	if err != nil {
		log.Printf("Ошибка при получении победителя: %v", err)
		return
	}

	winnerPrize, err := instance.WinnerPrize(nil)
	if err != nil {
		log.Printf("Ошибка при получении приза победителя: %v", err)
		return
	}

	fmt.Printf("Победитель: %s\n", winner.Hex())
	fmt.Printf("Приз победителя: %s BNB\n", weiToBNB(winnerPrize))
}

// Тест клейма наград
func testClaimRewards(client *ethclient.Client, instance *SingleLottery, ownerKey, user1Key, user2Key *ecdsa.PrivateKey) {
	ctx := context.Background()

	fmt.Println("\n=== Тестирование получения наград ===")

	// Получаем адреса
	ownerAddr := crypto.PubkeyToAddress(ownerKey.PublicKey)
	user1Addr := crypto.PubkeyToAddress(user1Key.PublicKey)
	user2Addr := crypto.PubkeyToAddress(user2Key.PublicKey)

	// Получаем ожидающие награды для каждого пользователя
	ownerReward, err := instance.PendingRewards(nil, ownerAddr)
	if err != nil {
		log.Printf("Ошибка при получении наград владельца: %v", err)
		return
	}

	user1Reward, err := instance.PendingRewards(nil, user1Addr)
	if err != nil {
		log.Printf("Ошибка при получении наград пользователя 1: %v", err)
		return
	}

	user2Reward, err := instance.PendingRewards(nil, user2Addr)
	if err != nil {
		log.Printf("Ошибка при получении наград пользователя 2: %v", err)
		return
	}

	fmt.Printf("Ожидающие награды владельца: %s BNB\n", weiToBNB(ownerReward))
	fmt.Printf("Ожидающие награды пользователя 1: %s BNB\n", weiToBNB(user1Reward))
	fmt.Printf("Ожидающие награды пользователя 2: %s BNB\n", weiToBNB(user2Reward))

	// Получаем начальные балансы
	ownerBalance, err := client.BalanceAt(ctx, ownerAddr, nil)
	if err != nil {
		log.Printf("Ошибка при получении баланса владельца: %v", err)
		return
	}

	user1Balance, err := client.BalanceAt(ctx, user1Addr, nil)
	if err != nil {
		log.Printf("Ошибка при получении баланса пользователя 1: %v", err)
		return
	}

	user2Balance, err := client.BalanceAt(ctx, user2Addr, nil)
	if err != nil {
		log.Printf("Ошибка при получении баланса пользователя 2: %v", err)
		return
	}

	// Клейм наград для владельца
	if ownerReward.Cmp(big.NewInt(0)) > 0 {
		chainID, _ := client.NetworkID(ctx)
		gasPrice, _ := client.SuggestGasPrice(ctx)
		nonce, _ := client.PendingNonceAt(ctx, ownerAddr)

		auth, err := bind.NewKeyedTransactorWithChainID(ownerKey, chainID)
		if err != nil {
			log.Printf("Ошибка при создании транзактора: %v", err)
			return
		}
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		//auth.GasLimit = gasLimit
		auth.GasPrice = gasPrice

		fmt.Println("Владелец забирает награду...")
		tx, err := instance.ClaimReward(auth)
		if err != nil {
			log.Printf("Ошибка при клейме наград владельца: %v", err)
			return
		}
		fmt.Printf("Транзакция отправлена: %s\n", tx.Hash().Hex())

		// Ждем подтверждения
		receipt, err := waitForTxReceipt(client, tx.Hash())
		if err != nil {
			log.Printf("Ошибка при ожидании подтверждения транзакции: %v", err)
			return
		}
		fmt.Printf("Транзакция подтверждена в блоке %d\n", receipt.BlockNumber)

		// Проверяем новый баланс
		newOwnerBalance, _ := client.BalanceAt(ctx, ownerAddr, nil)
		gasCost := new(big.Int).Mul(gasPrice, big.NewInt(int64(receipt.GasUsed)))
		expectedBalance := new(big.Int).Add(ownerBalance, ownerReward)
		expectedBalance = new(big.Int).Sub(expectedBalance, gasCost)

		fmt.Printf("Баланс владельца до: %s BNB\n", weiToBNB(ownerBalance))
		fmt.Printf("Баланс владельца после: %s BNB\n", weiToBNB(newOwnerBalance))
		fmt.Printf("Затраты на газ: %s BNB\n", weiToBNB(gasCost))

		// Проверяем, что баланс изменился правильно (с учетом затрат на газ)
		diff := new(big.Int).Sub(newOwnerBalance, expectedBalance)
		if diff.Abs(diff).Cmp(big.NewInt(1e14)) < 0 { // Допускаем небольшую погрешность
			fmt.Println("Владелец успешно получил награду!")
		} else {
			fmt.Printf("Ошибка: Баланс изменился неправильно. Разница: %s BNB\n", weiToBNB(diff))
		}
	}

	// Аналогично для пользователя 1
	if user1Reward.Cmp(big.NewInt(0)) > 0 {
		claimReward(client, instance, user1Key, user1Addr, user1Balance, user1Reward, "Пользователь 1")
	}

	// Аналогично для пользователя 2
	if user2Reward.Cmp(big.NewInt(0)) > 0 {
		claimReward(client, instance, user2Key, user2Addr, user2Balance, user2Reward, "Пользователь 2")
	}
}

// Вспомогательная функция для клейма наград
func claimReward(client *ethclient.Client, instance *SingleLottery, privateKey *ecdsa.PrivateKey,
	address common.Address, initialBalance, reward *big.Int, userLabel string) {

	ctx := context.Background()
	chainID, _ := client.NetworkID(ctx)
	gasPrice, _ := client.SuggestGasPrice(ctx)
	nonce, _ := client.PendingNonceAt(ctx, address)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Printf("Ошибка при создании транзактора для %s: %v", userLabel, err)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	//auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	fmt.Printf("%s забирает награду...\n", userLabel)
	tx, err := instance.ClaimReward(auth)
	if err != nil {
		log.Printf("Ошибка при клейме наград для %s: %v", userLabel, err)
		return
	}
	fmt.Printf("Транзакция отправлена: %s\n", tx.Hash().Hex())

	// Ждем подтверждения
	receipt, err := waitForTxReceipt(client, tx.Hash())
	if err != nil {
		log.Printf("Ошибка при ожидании подтверждения транзакции: %v", err)
		return
	}
	fmt.Printf("Транзакция подтверждена в блоке %d\n", receipt.BlockNumber)

	// Проверяем новый баланс
	newBalance, _ := client.BalanceAt(ctx, address, nil)
	gasCost := new(big.Int).Mul(gasPrice, big.NewInt(int64(receipt.GasUsed)))
	expectedBalance := new(big.Int).Add(initialBalance, reward)
	expectedBalance = new(big.Int).Sub(expectedBalance, gasCost)

	fmt.Printf("Баланс %s до: %s BNB\n", userLabel, weiToBNB(initialBalance))
	fmt.Printf("Баланс %s после: %s BNB\n", userLabel, weiToBNB(newBalance))
	fmt.Printf("Затраты на газ: %s BNB\n", weiToBNB(gasCost))

	// Проверяем, что баланс изменился правильно (с учетом затрат на газ)
	diff := new(big.Int).Sub(newBalance, expectedBalance)
	if diff.Abs(diff).Cmp(big.NewInt(1e14)) < 0 { // Допускаем небольшую погрешность
		fmt.Printf("%s успешно получил награду!\n", userLabel)
	} else {
		fmt.Printf("Ошибка: Баланс %s изменился неправильно. Разница: %s BNB\n", userLabel, weiToBNB(diff))
	}
}

// Вспомогательная функция для ожидания подтверждения транзакции
func waitForTxReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	ctx := context.Background()
	for i := 0; i < 30; i++ { // Пробуем 30 раз с интервалом в 2 секунды
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}
		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("транзакция не подтверждена в течение отведенного времени")
}

// Вспомогательная функция для конвертации wei в BNB
func weiToBNB(wei *big.Int) string {
	if wei == nil {
		return "0"
	}
	// 1 BNB = 10^18 wei
	f := new(big.Float).SetInt(wei)
	d := new(big.Float).SetInt(big.NewInt(1e18))
	result := new(big.Float).Quo(f, d)

	// Форматируем до 8 знаков после запятой
	return fmt.Sprintf("%.8f", result)
}
