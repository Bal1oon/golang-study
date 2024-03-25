package cmd

import (
	"fmt"
	"golang-study/config"
	"golang-study/network"
	"golang-study/repository"
	"golang-study/service"
)

type Cmd struct {
	config     *config.Config
	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filePath string) *Cmd {
	// c := new(Cmd)
	// c.config = config.NewConfig(filePath)
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	c.network.ServerStart(c.config.Server.Port)
	// network가 실행되어 서버 연결이 진행되면 이 아래는
	// 서버 종료 시까지 실행되지 않음
	fmt.Println(c.config.Server.Port)
	return c
}
