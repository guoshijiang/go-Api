//==================================================================
//创建时间：2017-4-23 首次创建
//功能描述：router
//创建人：郭世江
//修改记录：若要修改请记录
//==================================================================

package handler


import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func CreateHttpRouter() http.Handler {
	r := httprouter.New()
	r.GET("/api/heart/login", httprouter.Handle(GetAccountInfo))
	r.GET("/api/heart/data", httprouter.Handle(GetHeartData))
	r.GET("/api/heart/update", httprouter.Handle(UpdateResult))
	return r
}


