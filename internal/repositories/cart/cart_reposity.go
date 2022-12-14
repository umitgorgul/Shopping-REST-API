package cart_repository

import (
	"github/umitgorgul/Shopping-Cart-REST-API/internal/models/cart"

	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{
		db: db,
	}
}

func (r *cartRepository) Migration() error {
	err := r.db.Migrator().DropTable(&cart.Cart{})
	err = r.db.Migrator().DropTable(&cart.CartProduct{})
	if err != nil {
		return err
	}
	err = r.db.AutoMigrate(&cart.Cart{})
	err = r.db.AutoMigrate(&cart.CartProduct{})
	if err != nil {
		return err
	}
	return nil
}

func (r *cartRepository) Create(cart *cart.Cart) error {
	return r.db.Create(cart).Error
}

func (r *cartRepository) FindByID(id int) (*cart.Cart, error) {
	c := new(cart.Cart)
	err := r.db.Preload("CartProducts").First(c, id).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *cartRepository) Update(cart *cart.Cart) error {
	return r.db.Save(cart).Error
}

func (r *cartRepository) Delete(id int) error {
	return r.db.Where("id = ?", id).Delete(&cart.Cart{}).Error
}
