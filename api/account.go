package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/DevDreamy/go-finance/db/sqlc"
	"github.com/DevDreamy/go-finance/util"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	UserID      int32  `json:"user_id" binding:"required"`
	CategoryID  int32  `json:"category_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
	Value       int32     `json:"value" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	errOnValidateToken := util.GetTokenInHeaderAndVerify(ctx)
	if errOnValidateToken != nil {
		return
	}
	var req createAccountRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var categoryId = req.CategoryID
	var accountType = req.Type

	category, err := server.store.GetCategory(ctx, categoryId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if category.Type != accountType {
		ctx.JSON(http.StatusBadRequest, "Account type is different from Category type")
		return
	}

	arg := db.CreateAccountParams{
		UserID: req.UserID,
		CategoryID: categoryId,
		Title: req.Title,
		Type: accountType,
		Description: req.Description,
		Value: req.Value,
		Date: req.Date,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	errOnValidateToken := util.GetTokenInHeaderAndVerify(ctx)
	if errOnValidateToken != nil {
		return
	}
	var req getAccountRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
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

type getAccountsGraphRequest struct {
	UserID int32 `uri:"user_id" binding:"required"`
	Type   string `uri:"type" binding:"required"`
}

func (server *Server) getAccountsGraph(ctx *gin.Context) {
	errOnValidateToken := util.GetTokenInHeaderAndVerify(ctx)
	if errOnValidateToken != nil {
		return
	}
	var req getAccountsGraphRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAccountsGraphParams{
		UserID: req.UserID,
		Type: req.Type,
	}

	countGraph, err := server.store.GetAccountsGraph(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, countGraph)
}

type getAccountsReportsRequest struct {
	UserID int32 `uri:"user_id" binding:"required"`
	Type   string `uri:"type" binding:"required"`
}

func (server *Server) getAccountsReports(ctx *gin.Context) {
	errOnValidateToken := util.GetTokenInHeaderAndVerify(ctx)
	if errOnValidateToken != nil {
		return
	}
	var req getAccountsReportsRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAccountsReportsParams{
		UserID: req.UserID,
		Type: req.Type,
	}

	sumReports, err := server.store.GetAccountsReports(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sumReports)
}

type deleteAccountRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteAccount(ctx *gin.Context) {
	errOnValidateToken := util.GetTokenInHeaderAndVerify(ctx)
	if errOnValidateToken != nil {
		return
	}
	var req deleteAccountRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.DeleteAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, true)
}

type updateAccountRequest struct {
	ID          int32  `json:"id" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value        int32 `json:"value"`
}

func (server *Server) updateAccount(ctx *gin.Context) {
	errOnValidateToken := util.GetTokenInHeaderAndVerify(ctx)
	if errOnValidateToken != nil {
		return
	}
	var req updateAccountRequest
	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateAccountParams{
		ID: req.ID,
		Title: req.Title,
		Description: req.Description,
		Value: req.Value,
	}

	account, err := server.store.UpdateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountsRequest struct {
	UserID      int32     `json:"user_id" binding:"required"`
	Type        string    `json:"type"`
	CategoryID  int32     `json:"category_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (server *Server) getAccounts(ctx *gin.Context) {
	errOnValidateToken := util.GetTokenInHeaderAndVerify(ctx)
	if errOnValidateToken != nil {
		return
	}
	var req getAccountsRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAccountsParams{
		UserID: req.UserID,
		Type: req.Type,
		CategoryID: sql.NullInt32{
			Int32: req.CategoryID,
			Valid: req.CategoryID > 0,
		},
		Title: req.Title,
		Description: req.Description,
		Date: sql.NullTime{
			Time: req.Date,
			Valid: !req.Date.IsZero(),
		},
	}

	accounts, err := server.store.GetAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}