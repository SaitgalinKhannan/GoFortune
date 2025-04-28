package repository

import (
	"GoFortune/models"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type LotteryRepository struct {
	db *sql.DB
}

func NewLotteryRepository(db *sql.DB) *LotteryRepository {
	return &LotteryRepository{db: db}
}

func (r *LotteryRepository) GetAll() ([]models.Lottery, error) {
	rows, err := r.db.Query(`
        SELECT contract_address, network, environment, ticket_price, max_tickets,
               owner_fee_percent, winner_prize_percent, returned_prize_percent,
               start_time, duration, created_at
        FROM lotteries
        ORDER BY created_at DESC`)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(rows)

	lotteries := make([]models.Lottery, 0)

	for rows.Next() {
		var l models.Lottery
		if err := rows.Scan(
			&l.ContractAddress,
			&l.Network,
			&l.Environment,
			&l.TicketPrice,
			&l.MaxTickets,
			&l.OwnerFeePercent,
			&l.WinnerPrizePercent,
			&l.ReturnedPrizePercent,
			&l.StartTime,
			&l.Duration,
			&l.CreatedAt,
		); err != nil {
			return nil, err
		}
		lotteries = append(lotteries, l)
	}

	return lotteries, nil
}

func (r *LotteryRepository) Save(lottery *models.Lottery) error {
	query := `
    INSERT INTO lotteries (
        contract_address, network, environment, ticket_price, max_tickets,
        owner_fee_percent, winner_prize_percent, returned_prize_percent,
        start_time, duration, created_at
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := r.db.Exec(query,
		lottery.ContractAddress,
		lottery.Network,
		lottery.Environment,
		lottery.TicketPrice,
		lottery.MaxTickets,
		lottery.OwnerFeePercent,
		lottery.WinnerPrizePercent,
		lottery.ReturnedPrizePercent,
		lottery.StartTime,
		lottery.Duration,
		time.Now().Unix(),
	)

	return err
}

// DeleteByContractAddress удаляет лотерею по адресу контракта
// Возвращает:
// - количество удаленных записей (0 или 1)
// - ошибку операции
func (r *LotteryRepository) DeleteByContractAddress(contractAddress string) (bool, error) {
	query := `DELETE FROM lotteries WHERE contract_address = $1`
	result, err := r.db.Exec(query, contractAddress)
	if err != nil {
		return false, fmt.Errorf("delete failed: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("get rows affected failed: %w", err)
	}

	return rowsAffected > 0, nil
}

// Update обновляет данные лотереи по адресу контракта
// Возвращает:
// - количество обновленных записей (0 или 1)
// - ошибку операции
func (r *LotteryRepository) Update(lottery *models.Lottery) (bool, error) {
	query := `
        UPDATE lotteries
        SET 
            start_time = $1,
            duration = $2
        WHERE contract_address = $3 
    `
	result, err := r.db.Exec(query,
		lottery.StartTime,
		lottery.Duration,
		lottery.ContractAddress,
	)

	if err != nil {
		return false, fmt.Errorf("update failed: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("get rows affected failed: %w", err)
	}

	return rowsAffected > 0, nil
}
