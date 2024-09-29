package service

import "eCommerce/config"

type Service struct {
	config *config.Config
}

func NewService(config *config.Config) (*Service, error) {
	r := &Service{
		config: config,
	}

	return r, nil
}