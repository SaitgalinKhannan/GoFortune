package config

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
)

// Chain представляет информацию о блокчейне.
type Chain struct {
	Name     string `json:"name"`     // Название блокчейна
	ChainID  int    `json:"chainId"`  // Идентификатор цепи (chain ID)
	Currency string `json:"currency"` // Валюта блокчейна
	RPCUrl   string `json:"rpcUrl"`   // URL RPC-узла
}

// Chains представляет массив блокчейнов.
type Chains struct {
	Chains []Chain `json:"chains"`
}

// ContainsNetwork проверяет, содержится ли указанная сеть в списке блокчейнов.
func (c *Chains) ContainsNetwork(network string) bool {
	for _, chain := range c.Chains {
		if chain.Name == network {
			return true
		}
	}
	return false
}

// Config объединяет все настройки приложения.
type Config struct {
	PrivateKey string  // Приватный ключ из .env
	Chains     *Chains // Информация о блокчейнах из JSON
}

// LoadConfig загружает конфигурацию из .env и JSON.
func LoadConfig(envPath, chainsPath string) *Config {
	// Загрузка .env файла
	viper.SetConfigFile(envPath)
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	privateKey := viper.GetString("PRIVATE_KEY")
	if privateKey == "" {
		logrus.Fatal("PRIVATE_KEY не найден в .env файле")
	}

	// Загрузка JSON файла с блокчейнами
	chains, err := loadChainsFromJSON(chainsPath)
	if err != nil {
		logrus.Fatalf("Ошибка загрузки JSON файла: %v", err)
	}

	return &Config{
		PrivateKey: privateKey,
		Chains:     chains,
	}
}

// loadChainsFromJSON загружает данные о блокчейнах из JSON файла.
func loadChainsFromJSON(filePath string) (*Chains, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var chains Chains
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл: %w", err)
	}

	if err := json.Unmarshal(bytes, &chains); err != nil {
		return nil, fmt.Errorf("не удалось распарсить JSON: %w", err)
	}

	return &chains, nil
}
