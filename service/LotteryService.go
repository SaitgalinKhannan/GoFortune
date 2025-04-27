package service

import (
	"GoFortune/config"
	"GoFortune/contracts"
	"GoFortune/models"
	"GoFortune/repository"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

type LotteryService struct {
	repo     *repository.LotteryRepository
	client   *ethclient.Client
	ownerKey *ecdsa.PrivateKey
}

func NewLotteryService(repo *repository.LotteryRepository, client *ethclient.Client, ownerKey *ecdsa.PrivateKey) *LotteryService {
	return &LotteryService{
		repo:     repo,
		client:   client,
		ownerKey: ownerKey,
	}
}

func (s *LotteryService) CreateLottery(req models.CreateLotteryRequest, cfg *config.Config) (*models.CreateLotteryResponse, error) {
	if err := validateRequest(req, *cfg.Chains); err != nil {
		return nil, err
	}

	ticketPrice, ok := new(big.Int).SetString(req.TicketPrice, 10)
	if !ok || ticketPrice.Cmp(big.NewInt(0)) <= 0 {
		return nil, fmt.Errorf("invalid ticket price")
	}

	contractAddress, tx, err := contracts.DeployContract(s.client, s.ownerKey)
	if err != nil {
		return nil, err
	}

	lottery := &models.Lottery{
		ContractAddress:      contractAddress.Hex(),
		Network:              req.Network,
		Environment:          req.Environment,
		TicketPrice:          req.TicketPrice,
		MaxTickets:           req.MaxTickets,
		OwnerFeePercent:      req.OwnerFeePercent,
		WinnerPrizePercent:   req.WinnerPrizePercent,
		ReturnedPrizePercent: req.ReturnedPrizePercent,
		StartTime:            req.StartTime,
		Duration:             req.Duration,
		CreatedAt:            time.Now().Unix(),
	}

	if err := s.repo.Save(lottery); err != nil {
		return nil, err
	}

	return &models.CreateLotteryResponse{
		ContractAddress: contractAddress.Hex(),
		TransactionHash: tx.Hash().Hex(),
	}, nil
}

func (s *LotteryService) GetLotteries() ([]models.Lottery, error) {
	return s.repo.GetAll()
}

// Валидация запроса
func validateRequest(req models.CreateLotteryRequest, chains config.Chains) error {
	// Проверка сети
	if !chains.ContainsNetwork(req.Network) {
		return fmt.Errorf("network '%s' is not supported", req.Network)
	}

	// Проверка окружения
	if req.Environment != "testnet" && req.Environment != "mainnet" {
		return fmt.Errorf("environment must be 'testnet' or 'mainnet'")
	}

	// Проверка maxTickets
	if req.MaxTickets == 0 {
		return fmt.Errorf("maxTickets must be greater than 0")
	}

	// Проверка суммы процентов
	if req.OwnerFeePercent+req.WinnerPrizePercent+req.ReturnedPrizePercent != 100 {
		return fmt.Errorf("sum of percentages must be 100")
	}

	// Проверка startTime
	if req.StartTime < time.Now().Unix() {
		return fmt.Errorf("startTime must be in the future")
	}

	// Проверка duration
	if req.Duration == 0 {
		return fmt.Errorf("duration must be greater than 0")
	}

	return nil
}
