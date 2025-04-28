package handlers

import (
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, h *LotteryHandler) {
	lotterySubrouter := r.PathPrefix("/lottery").Subrouter()

	lotterySubrouter.HandleFunc("/create", h.CreateLottery).Methods("POST")
	lotterySubrouter.HandleFunc("/all", h.GetLotteries).Methods("GET")
	lotterySubrouter.HandleFunc("/delete/{contractAddress}", h.DeleteLotteries).Methods("GET")
}
