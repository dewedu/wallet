-- 创建数据库
CREATE DATABASE IF NOT EXISTS walletdb CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE walletdb;

-- 钱包表
CREATE TABLE wallets (
                         wallet_id INT PRIMARY KEY AUTO_INCREMENT COMMENT '钱包ID',
                         user_id BIGINT NOT NULL COMMENT '用户ID',
                         balance BIGINT NOT NULL DEFAULT 0 COMMENT '余额(分)',
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间'
) COMMENT '钱包表';

-- 转账记录表
CREATE TABLE transfers (
                           transfer_id INT PRIMARY KEY AUTO_INCREMENT COMMENT '转账记录ID',
                           from_wallet_id INT NOT NULL COMMENT '转出钱包ID',
                           to_wallet_id INT NOT NULL COMMENT '转入钱包ID',
                           amount BIGINT NOT NULL DEFAULT 0 COMMENT '转账金额(分)',
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '转账时间'
) COMMENT '转账记录表';

-- 创建索引
CREATE INDEX idx_wallets_user ON wallets(user_id) COMMENT '用户钱包索引';