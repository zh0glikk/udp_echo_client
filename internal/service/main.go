package service

import (
	"bufio"
	"context"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"net"
	"os"

	"udp_echo_client/internal/config"
)

type service struct {
	logger           *logan.Entry
	cfg 			 config.Config
}

func NewService(cfg config.Config) *service {
	return &service{
		logger:    cfg.Log(),
		cfg:	   cfg,
	}
}

func (s *service) Run(ctx context.Context) error{
	for {
		conn, _ := net.Dial("udp", fmt.Sprintf("%v:%v", s.cfg.Server().Ip, s.cfg.Server().Port))

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text + "\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)
	}

	return nil
}

