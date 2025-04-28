package handlers

import (
	"GoFortune/config"
	"GoFortune/models"
	"GoFortune/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type LotteryHandler struct {
	service *service.LotteryService
	cfg     *config.Config
}

func NewLotteryHandler(service *service.LotteryService, cfg *config.Config) *LotteryHandler {
	return &LotteryHandler{service: service, cfg: cfg}
}

func (h *LotteryHandler) GetLotteries(w http.ResponseWriter, _ *http.Request) {
	lotteries, err := h.service.GetLotteries()
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Database error")
		return
	}

	respondJSON(w, http.StatusOK, lotteries)
}

func (h *LotteryHandler) CreateLottery(w http.ResponseWriter, r *http.Request) {
	var req models.CreateLotteryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	resp, err := h.service.CreateLottery(req, h.cfg)
	if err != nil {
		logrus.Error(err)
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, resp)
}

func (h *LotteryHandler) DeleteLotteries(w http.ResponseWriter, r *http.Request) {
	contractAddress := mux.Vars(r)["contractAddress"]

	result, err := h.service.DeleteLottery(contractAddress)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Database error")
		return
	}

	if !result {
		sendError(w, http.StatusNotFound, "Can`t delete lottery: "+contractAddress)
		return
	}

	respondJSON(w, http.StatusOK, result)
}

func sendError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, models.ErrorResponse{Error: message})
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
