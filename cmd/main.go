package main

import (
	"altair-test/internal/service"
	"altair-test/pkg/utils"
	"log"
)

func main() {
	tools := utils.NewTools()
	s := service.NewService(tools)

	if err := s.Exec(); err != nil {
		log.Fatal(err)
	}
}
