package models

// Lottery представляет запись о лотерее в базе данных
type Lottery struct {
	ContractAddress      string `json:"contractAddress"`      // Адрес контракта
	Network              string `json:"network"`              // Сеть (eth, bsc)
	Environment          string `json:"environment"`          // testnet или mainnet
	TicketPrice          string `json:"ticketPrice"`          // Цена билета в wei
	MaxTickets           uint64 `json:"maxTickets"`           // Максимум билетов
	OwnerFeePercent      uint64 `json:"ownerFeePercent"`      // % владельцу
	WinnerPrizePercent   uint64 `json:"winnerPrizePercent"`   // % победителю
	ReturnedPrizePercent uint64 `json:"returnedPrizePercent"` // % возврат
	StartTime            int64  `json:"startTime"`            // Время начала (Unix timestamp)
	Duration             uint64 `json:"duration"`             // Продолжительность в секундах
	CreatedAt            int64  `json:"createdAt"`            // Время создания (Unix timestamp)
}

// CreateLotteryRequest представляет запрос для создания новой лотереи
type CreateLotteryRequest struct {
	Network              string `json:"network"`              // Сеть: eth, bsc
	Environment          string `json:"environment"`          // testnet или mainnet
	TicketPrice          string `json:"ticketPrice"`          // Цена билета в wei
	MaxTickets           uint64 `json:"maxTickets"`           // Максимум билетов
	OwnerFeePercent      uint64 `json:"ownerFeePercent"`      // % владельцу
	WinnerPrizePercent   uint64 `json:"winnerPrizePercent"`   // % победителю
	ReturnedPrizePercent uint64 `json:"returnedPrizePercent"` // % возврат
	StartTime            int64  `json:"startTime"`            // Время начала (Unix timestamp)
	Duration             uint64 `json:"duration"`             // Продолжительность в секундах
}

// CreateLotteryResponse представляет ответ после успешного создания лотереи
type CreateLotteryResponse struct {
	ContractAddress string `json:"contractAddress"` // Адрес развернутого контракта
	TransactionHash string `json:"transactionHash"` // Хэш транзакции деплоя
}

// ErrorResponse представляет ошибку API
type ErrorResponse struct {
	Error string `json:"error"` // Сообщение об ошибке
}
