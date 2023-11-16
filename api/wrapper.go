package api

import "github.com/gin-gonic/gin"

type WrapperServer struct {
	server *Server
}

// (POST /auth/change_password)
func (wrapper *WrapperServer) PostQuran(c *gin.Context) {
	// wrapper.server.PostQuran(c)
}

// (DELETE /auth/delete_account)
func (wrapper *WrapperServer) GetQuranPage(c *gin.Context, page int32) {
	wrapper.server.GetQuranPage(c)
}
