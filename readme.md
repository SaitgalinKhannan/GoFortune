# GoFortune - Decentralized Lottery Platform

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Solidity](https://img.shields.io/badge/Solidity-%23363636.svg?style=for-the-badge&logo=solidity&logoColor=white)
![Ethereum](https://img.shields.io/badge/Ethereum-3C3C3D?style=for-the-badge&logo=Ethereum&logoColor=white)

GoFortune is a decentralized lottery platform built on Ethereum Virtual Machine (EVM) compatible blockchains. The project consists of a Solidity smart contract for lottery mechanics and a Go backend service for contract management and deployment.

## 🎯 Features

### Smart Contract Features
- **Fair Lottery System**: Transparent and verifiable lottery draws using blockchain randomness
- **Flexible Prize Distribution**: Configurable percentages for owner fees, winner prizes, and participant returns
- **Multiple Ticket Purchase**: Users can buy multiple tickets in a single transaction
- **Secure Reward Claims**: Protected reward claiming mechanism with reentrancy guards
- **Emergency Refunds**: Owner can initiate emergency refunds if needed
- **Ownership Transfer**: Contract ownership can be transferred to new addresses

### Backend Features
- **Multi-Chain Support**: Works with BNB Smart Chain, Ethereum, and other EVM chains
- **RESTful API**: Complete API for lottery management
- **Database Integration**: PostgreSQL database for lottery tracking
- **Contract Deployment**: Automated smart contract deployment
- **Configuration Management**: Environment-based configuration system

## 🏗️ Architecture

```
GoFortune/
├── contracts/           # Solidity smart contracts
├── config/             # Configuration files
├── handlers/           # HTTP request handlers
├── models/             # Data models
├── repository/         # Database layer
├── service/            # Business logic
└── utils/              # Utility functions
```

## 🚀 Quick Start

### Prerequisites
- Go 1.24+
- Node.js and npm/yarn
- PostgreSQL database
- Solidity compiler (solc)
- abigen tool

### Installation

1. **Clone the repository**
```bash
git clone https://github.com/yourusername/GoFortune.git
cd GoFortune
```

2. **Install Go dependencies**
```bash
go mod download
```

3. **Install Solidity compiler**
```bash
npm install -g solc
```

4. **Set up environment variables**
   Create `.env` file in the `config/` directory:
```env
MNEMONIC=your_mnemonic_phrase_here
PRIVATE_KEY=your_private_key_here
```

5. **Configure database**
   Update database credentials in `repository/init_db.go` or use environment variables.

6. **Set up database**
```sql
CREATE DATABASE lottery_db;
-- Run the SQL script from repository/db.sql
```

### Compile Smart Contract

Generate ABI:
```bash
solc --abi contracts/SingleLottery.sol -o contracts/abi
```

Generate bytecode:
```bash
solc --bin contracts/SingleLottery.sol -o contracts/bin
```

Generate Go bindings:
```bash
abigen --abi contracts/abi/SingleLottery.abi --pkg main --type SingleLottery --out SingleLottery.go --bin contracts/bin/SingleLottery.bin
```

### Run the Application

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## 📡 API Endpoints

### Create Lottery
```http
POST /lottery/create
Content-Type: application/json

{
  "network": "BNB Smart Chain Testnet",
  "environment": "testnet",
  "ticketPrice": "10000000000000000",
  "maxTickets": 100,
  "ownerFeePercent": 10,
  "winnerPrizePercent": 50,
  "returnedPrizePercent": 40,
  "startTime": 1640995200,
  "duration": 86400
}
```

### Get All Lotteries
```http
GET /lottery/all
```

### Delete Lottery
```http
GET /lottery/delete/{contractAddress}
```

## 🎲 Smart Contract Usage

### Constructor Parameters
- `_ticketPrice`: Price per ticket in wei
- `_maxTickets`: Maximum number of tickets
- `_ownerFeePercent`: Percentage fee for contract owner
- `_winnerPrizePercent`: Percentage prize for winner
- `_returnedPrizePercent`: Percentage returned to participants

**Note**: The three percentages must sum to 100%.

### Key Functions

#### For Users
- `buyTickets(uint256 _amount)`: Purchase lottery tickets
- `claimReward()`: Claim pending rewards
- `getMyInfo()`: Get personal ticket count and pending rewards

#### For Owner
- `drawLottery(uint256 _randomSeed)`: Conduct the lottery draw
- `emergencyRefund()`: Initiate emergency refund to all participants
- `transferOwnership(address newOwner)`: Transfer contract ownership

#### View Functions
- `getRemainingTickets()`: Get number of tickets still available
- `getContractBalance()`: Get current contract balance

## 🔧 Configuration

### Supported Networks
The project supports multiple EVM-compatible networks configured in `config/chains.json`:

- **BNB Smart Chain Mainnet** (Chain ID: 56)
- **BNB Smart Chain Testnet** (Chain ID: 97)
- **Ethereum Mainnet** (Chain ID: 1)

### Environment Settings
- `testnet`: For testing purposes
- `mainnet`: For production deployment

## 🛡️ Security Features

- **Reentrancy Protection**: Prevents reentrancy attacks in reward claiming
- **Input Validation**: Comprehensive validation of all user inputs
- **Access Control**: Owner-only functions for critical operations
- **Safe Math**: Protection against integer overflow/underflow
- **Transparent Randomness**: Uses blockchain-based randomness sources

## 🧪 Testing

Run the test suite:
```bash
go run test.go
```

The test file includes comprehensive testing scenarios:
- Contract deployment
- Ticket purchasing
- Lottery drawing
- Reward claiming
- Balance verification

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Useful Links

- [Solidity Documentation](https://docs.soliditylang.org/)
- [Go Ethereum Documentation](https://geth.ethereum.org/docs/)
- [BNB Smart Chain Documentation](https://docs.bnbchain.org/)

**Built with ❤️ for the decentralized future**