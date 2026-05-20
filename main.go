// AC搜索接口服务 (AC search interface service).
//
// 入口点：通过数据库初始化（init 导入）→ 路由加载 → 服务启动。
package main

import (
	_ "app/databases"
	"app/routers"
)

func main() {
	r := routers.Load()
	r.Run()
}
