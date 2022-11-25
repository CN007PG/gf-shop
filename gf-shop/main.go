package main

import (
	_ "gf-shop/internal/logic"
	_ "gf-shop/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-shop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
