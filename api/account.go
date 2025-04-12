package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/juanparraiv/simple-bank/db/sqlc"
)

// createAccountRequest defines the request body for the create account request
type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
	// Balance  int64  `json:"balance" binding:"required,min=0"`
}
type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type listAccountsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
type listAccountsResponse struct {
	Accounts []db.Account `json:"accounts"`
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} createAccount
// @Router /accounts [post]
// @Param account body createAccountRequest true "Create account request"
// @Success 200 {object} db.Account
// @Description Create a new account
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		// Balance:  req.Balance,
	}
	// Create the account in the database
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// @BasePath /api/v1

// PingExample godoc

// @Schemes
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} getAccount
// @Router /accounts/:id [get]
// @Param id path int true "Account ID"
// @Success 200 {object} db.Account
// @Description Get account by ID
func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} listAccount
// @Router /accounts [get]
// @Param page_id query int true "Page ID"
// @Param page_size query int true "Page Size"
func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	fmt.Println(arg)
	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if len(accounts) == 0 {
		ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf("no accounts found")))
		return
	}
	// if err == sql.ErrNoRows {
	// 	ctx.JSON(http.StatusNotFound, errorResponse(err))
	// 	return
	// }
	fmt.Println(accounts)
	ctx.JSON(http.StatusOK, listAccountsResponse{
		Accounts: accounts,
	})
}
