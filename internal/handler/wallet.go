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

// CreateWallet 创建钱包
// @Summary      创建钱包
// @Description  创建一个新钱包并返回钱包ID
// @Tags         钱包管理
// @Accept       json
// @Success      200  {object}  int32  "钱包ID"
// @Router       /wallets [post]
func (h *WalletHandler) CreateWallet(c *gin.Context) {
	wallet, err := h.service.CreateWallet()
	if err != nil {
		ginx.Error(c, err, "创建钱包失败")
		return
	}

	ginx.Success(c, wallet.WalletId)
}

// GetWallet 查询钱包余额
// @Summary      查询钱包余额
// @Description  根据钱包ID查询余额
// @Tags         钱包管理
// @Accept       json
// @Param        id   path      int32  true  "钱包ID"
// @Success      200  {object}  int64  "余额(分)"
// @Router       /wallets/{id} [get]
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

type TransferRequest struct {
	FromWalletID int32 `json:"from_wallet_id" binding:"required"` // 转出钱包
	ToWalletID   int32 `json:"to_wallet_id" binding:"required"`   // 转入钱包
	Amount       int64 `json:"amount" binding:"required,gt=0"`    // 转出金额（分）
}

// Transfer 钱包转账
// @Summary      钱包转账
// @Description  从一个钱包向另一个钱包转账
// @Tags         转账操作
// @Accept       json
// @Param        request  body  TransferRequest  true  "转账请求"
// @Success      200      {object}  string  "OK"
// @Router       /wallets/transfer [post]
func (h *WalletHandler) Transfer(c *gin.Context) {
	var req TransferRequest
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

	ginx.Success(c, "OK")
}
