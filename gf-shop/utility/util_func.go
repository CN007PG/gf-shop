package utility

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 绑定动态服务函数
func BindStaticServer(name string) {
	s := g.Server(name)
	s.Start()
}

// 绑定动态服务函数
func BindDynamicServer(name string, handlerOrObject interface{}, handlers ...ghttp.HandlerFunc) {
	s := g.Server(name)
	s.Group("/", func(group *ghttp.RouterGroup) {
		if len(handlers) > 0 {
			//设置中间件
			group.Middleware(handlers...)
		}
		//绑定对象
		group.Bind(
			handlerOrObject,
		)
	})
	s.Start()
}

// 返回数据函数
func ReturnJsonData(ctx context.Context, data interface{}) {
	g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
		"code": 1,
		"user": data,
	})
	return
}

// 返回错误函数
func ReturnErrorData(ctx context.Context, errcode string, msg string) {
	g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
		"code": errcode,
		"msg":  msg,
	})
	return
}
