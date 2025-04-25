package main

import "hakuji/util"

func main() {
	r := util.Service()
	r.Run(":8080") // 在8080端口启动服务
}
