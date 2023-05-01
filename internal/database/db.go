package database

import (
	"fmt"
	"github.com/Anarr/entain/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Postgresql struct {
	conf *config.AppConfig
	DB   *gorm.DB
}

func New(conf *config.AppConfig) *Postgresql {
	return &Postgresql{
		conf: conf,
	}
}

func (p *Postgresql) Connect() error {
	pgConfig := p.conf.Postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable",
		pgConfig.Host, pgConfig.User, pgConfig.Password, pgConfig.Name, pgConfig.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: p.conf.Postgres.Name + ".",
		},
	})
	if err != nil {
		return err
	}

	p.DB = db

	return nil
}
