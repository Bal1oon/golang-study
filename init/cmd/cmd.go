package cmd

import (
	"fmt"
	"golang-study/config"
	"golang-study/network"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config:  config.NewConfig(filePath),
		network: network.NewNetwork(),
	}

	c.network.ServerStart(c.config.Server.Port)
	// network가 실행되어 서버 연결이 진행되면 이 아래는
	// 서버 종료 시까지 실행되지 않음
	fmt.Println(c.config.Server.Port)
	return c
}
