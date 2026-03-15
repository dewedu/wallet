package service

import (
	"errors"
	"math/rand"
	"time"
	"wallet/data/model"

	"gorm.io/gorm"
)

type WalletService struct {
	db *gorm.DB
}

func NewWalletService(db *gorm.DB) *WalletService {
	return &WalletService{db: db}
}

// CreateWallet 创建钱包
func (s *WalletService) CreateWallet() (*model.Wallets, error) {
	wallet := &model.Wallets{
		WalletId: 0,
		UserId:   int64(rand.Int()),
	}

	if err := s.db.Create(wallet).Error; err != nil {
		return nil, err
	}
	return wallet, nil
}

// GetWallet 获取钱包
func (s *WalletService) GetWallet(id int32) (*model.Wallets, bool, error) {
	var wallet model.Wallets
	if err := s.db.First(&wallet, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &wallet, true, nil
}

// Transfer 转账
func (s *WalletService) Transfer(fromID, toID int32, amount int64) error {
	// 开启事务
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 1. 检查源钱包是否存在
		var fromWallet model.Wallets
		if err := tx.First(&fromWallet, fromID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("转出钱包不存在")
			}
			return errors.New("查询转出钱包失败")
		}

		// 2. 检查目标钱包是否存在
		var toWallet model.Wallets
		if err := tx.First(&toWallet, toID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("转入钱包不存在")
			}
			return errors.New("查询转入钱包失败")
		}

		// 3. 检查源钱包余额是否充足
		if fromWallet.Balance < amount {
			return errors.New("转出钱包余额不足")
		}

		// 4. 执行转账 - 更新转出钱包余额
		if err := tx.Model(&fromWallet).
			Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return errors.New("更新转出钱包余额失败")
		}

		// 5. 执行转账 - 更新转入钱包余额
		if err := tx.Model(&toWallet).
			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return errors.New("更新转入钱包余额失败")
		}

		// 6. 记录转账记录
		transfer := &model.Transfers{
			FromWalletId: fromID,
			ToWalletId:   toID,
			Amount:       amount,
			CreatedAt:    time.Now(),
		}

		if err := tx.Create(transfer).Error; err != nil {
			return errors.New("记录转账信息失败")
		}

		return nil
	})
}
