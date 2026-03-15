package handler

import (
	"strconv"
	"wallet/internal/service"
	"wallet/pkg/ginx"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	service *service.WalletService
}

func NewWalletHandler(service *service.WalletService) *WalletHandler {
	return &WalletHandler{service: service}
}

// CreateWallet 创建钱包并返回钱包ID
func (h *WalletHandler) CreateWallet(c *gin.Context) {
	wallet, err := h.service.CreateWallet()
	if err != nil {
		ginx.Error(c, err, "创建钱包失败")
		return
	}

	ginx.Success(c, wallet.WalletId)
}

// GetWallet 获取钱包信息
func (h *WalletHandler) GetWallet(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ginx.Error(c, err, "参数错误")
		return
	}

	wallet, ok, err := h.service.GetWallet(int32(id))
	if err != nil {
		ginx.Error(c, err, "查询错误")
		return
	}
	if !ok {
		ginx.NotFound(c, "ID不存在")
		return
	}

	ginx.Success(c, wallet.Balance)
}

var req struct {
	FromWalletID int32 `json:"from_wallet_id" binding:"required"` // 转出钱包
	ToWalletID   int32 `json:"to_wallet_id" binding:"required"`   // 转入钱包
	Amount       int64 `json:"amount" binding:"required,gt=0"`    // 转出金额（分）
}

// Transfer 转账
func (h *WalletHandler) Transfer(c *gin.Context) {
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.Error(c, err, "ID参数错误")
		return
	}

	// 参数验证
	if req.FromWalletID <= 0 {
		ginx.NotFound(c, "转出钱包ID不能为空或小于等于0")
		return
	}
	if req.ToWalletID <= 0 {
		ginx.NotFound(c, "转入钱包ID不能为空或小于等于0")
		return
	}
	if req.FromWalletID == req.ToWalletID {
		ginx.NotFound(c, "转出钱包和转入钱包不能相同")
		return
	}
	if req.Amount <= 0 {
		ginx.NotFound(c, "转账金额必须大于0")
		return
	}

	if err := h.service.Transfer(req.FromWalletID, req.ToWalletID, req.Amount); err != nil {
		ginx.Error(c, err, "转账失败")
		return
	}

	ginx.Success(c, "ok")
}
