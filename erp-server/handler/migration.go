package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"erp-server/model"

	"github.com/go-gormigrate/gormigrate/v2"
)

type Migration struct {
	db *gorm.DB
}

type IMigration interface {
	Migrate(c *gin.Context)
}

func NewMigration(db *gorm.DB) *Migration {
	return &Migration{db: db}
}

func (h *Migration) BaseMigrate(ctx *gin.Context) error {
	err := h.db.Exec(`
			CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		`).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *Migration) Migrate(ctx *gin.Context) {
	// put your migrations at the begin of the list
	migrate := gormigrate.New(h.db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			// delete table money, product, order
			// update column price to float in table money, product, order
			// add table order_item
			ID: "20230815182141",
			Migrate: func(tx *gorm.DB) error {
				if err := h.db.Exec(`
			drop table money, product, order if exists;
		`).Error; err != nil {
					fmt.Errorf(err.Error())
				}
				if err := h.db.AutoMigrate(
					&model.Money{}, &model.Order{}, &model.Product{}, &model.OrderItem{}); err != nil {
					return err
				}
				return nil
			},
		},
		{
			// add uuid extension
			// add table user, business, product, order, money
			ID: "20230814200514",
			Migrate: func(tx *gorm.DB) error {
				h.BaseMigrate(ctx)
				if err := h.db.AutoMigrate(
					&model.User{},
					&model.Business{},
					&model.Product{},
					&model.Order{},
					&model.Money{}); err != nil {
					return err
				}
				return nil
			},
		},
	})
	err := migrate.Migrate()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
}