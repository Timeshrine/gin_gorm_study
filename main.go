package main

import (
	"gin+gorm_study/conf"
	"gin+gorm_study/routes"
)

func main() {
	conf.Init()
	r := routes.NerRouter()
	_ = r.Run(conf.HttpPort)
}
