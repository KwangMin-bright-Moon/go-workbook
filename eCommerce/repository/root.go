package repository

import (
	"eCommerce/config"
	"eCommerce/repository/mongo"
)

type Repository struct {
	config *config.Config

	Mongo *mongo.Mongo
}

func NewRepository(config *config.Config) (*Repository, error) {
	r := &Repository{
		config: config,
	}

	var err error
	
	if r.Mongo, err = mongo.NewMongo(config); err != nil {
		panic(err)
	} 

	return r, nil
}