package repository

import (
	"GoFortune/models"
	"database/sql"
	"github.com/sirupsen/logrus"
	"time"
)

type LotteryRepository struct {
	db *sql.DB
}

func NewLotteryRepository(db *sql.DB) *LotteryRepository {
	return &LotteryRepository{db: db}
}

func (r *LotteryRepository) Save(lottery *models.Lottery) error {
	/*query := `
	  INSERT INTO lotteries (
	      contract_address, network, environment, ticket_price, max_tickets,
	      owner_fee_percent, winner_prize_percent, returned_prize_percent,
	      start_time, duration, created_at
	  ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`*/

	query := `
    INSERT INTO lotteries (
        contract_address, network, environment, ticket_price, max_tickets,
        owner_fee_percent, winner_prize_percent, returned_prize_percent,
        start_time, duration, created_at
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	logrus.Printf("Executing query: %s\nValues: %+v", query, lottery)

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

	var lotteries []models.Lottery

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
