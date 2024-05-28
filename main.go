package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "myshop/internal/logic"
	_ "myshop/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"myshop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
