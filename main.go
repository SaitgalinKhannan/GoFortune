package main

import (
	"GoFortune/config"
	"GoFortune/handlers"
	"GoFortune/repository"
	"GoFortune/service"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func main() {
	// Укажите пути к файлам
	envPath := "./config/.env"
	chainsPath := "./config/chains.json"

	// Загрузка конфигурации
	cfg := config.LoadConfig(envPath, chainsPath)
	ownerKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatalf("Ошибка при создании ownerKey: %v", err)
	}

	db, err := repository.InitializeDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	// Подключение к сети
	client, err := ethclient.Dial(cfg.Chains.Chains[1].RPCUrl)
	if err != nil {
		log.Fatalf("Не удалось подключиться к BNB testnet: %v", err)
	}
	fmt.Println("Успешно подключились к BNB testnet")

	// Создаем зависимости
	lotteryRepo := repository.NewLotteryRepository(db)
	lotteryService := service.NewLotteryService(lotteryRepo, client, ownerKey)
	lotteryHandler := handlers.NewLotteryHandler(lotteryService, cfg)

	// Регистрируем роуты
	r := mux.NewRouter()
	/*r.HandleFunc("/create-lottery", lotteryHandler.CreateLottery(cfg)).Methods("POST")
	r.HandleFunc("/lotteries", lotteryHandler.GetLotteries).Methods("GET")*/
	handlers.SetupRoutes(r, lotteryHandler)

	// Запускаем сервер
	logrus.Println("Server started on :8080")
	logrus.Fatal(http.ListenAndServe(":8080", r))
}
