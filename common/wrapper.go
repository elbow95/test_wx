package common

import (
	"fmt"
	"log"
	"wxcloudrun-golang/util"

	"github.com/gin-gonic/gin"
)

func HandlerWrapper[Req any, Resp any](handler func(c *gin.Context, req *Req) (Resp, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		req := new(Req)
		if err := c.ShouldBind(req); err != nil {
			log.Println(err)
			Failed(c, fmt.Sprintf("%s,请稍后刷新重试", ParamInvalid))
			return
		}
		log.Printf("http req: %+v", util.MarshalJsonIgnoreError(req))
		data, err := handler(c, req)
		if err != nil {
			Failed(c, err.Error())
			return
		}
		log.Printf("http data: %+v", util.MarshalJsonIgnoreError(data))
		Success(c, data)
	}
}
