create database lottery_db;

CREATE TABLE lotteries
(
    id                     BIGINT      NOT NULL GENERATED ALWAYS AS IDENTITY,
    contract_address       VARCHAR(42) NOT NULL,
    network                VARCHAR(50) NOT NULL,
    environment            VARCHAR(20) NOT NULL,
    ticket_price           VARCHAR(78) NOT NULL,
    max_tickets            BIGINT      NOT NULL,
    owner_fee_percent      INT         NOT NULL CHECK (owner_fee_percent <= 100),
    winner_prize_percent   INT         NOT NULL CHECK (winner_prize_percent <= 100),
    returned_prize_percent INT         NOT NULL CHECK (returned_prize_percent <= 100),
    start_time             BIGINT      NOT NULL,
    duration               BIGINT      NOT NULL,
    created_at             BIGINT      NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (contract_address)
);

CREATE INDEX idx_network ON lotteries(network);
CREATE INDEX idx_environment ON lotteries(environment);