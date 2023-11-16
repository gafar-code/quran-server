package api

import (
	"net/http"

	db "github.com/gafar-code/quran-server/db/sqlc"
	"github.com/gin-gonic/gin"
)

type getPageRequest struct {
	Page int64 `uri:"page" binding:"required,min=1,max=604"`
}

type createPageRequest struct {
	PageID    int64                `json:"page_id"`
	ListAyah  []db.CreatePageAyah  `json:"list_ayah"`
	ListSurah []db.CreatePageSurah `json:"list_surah"`
}

func (server *Server) PostQuran(ctx *gin.Context) {
	var req createPageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseErr{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	arg := db.CreatePageParams{
		PageID:    req.PageID,
		ListAyah:  req.ListAyah,
		ListSurah: req.ListSurah,
	}

	getPage, err := server.page.CreatePage(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseErr{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, getPage)
}

func (server *Server) GetQuranPage(ctx *gin.Context) {
	var req getPageRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseErr{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	page, err := server.page.GetPage(ctx, req.Page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseErr{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ResponseData{
		Code:    200,
		Message: "success",
		Data:    page,
	})
}
