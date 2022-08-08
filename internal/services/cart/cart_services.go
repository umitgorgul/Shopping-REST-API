package cart_service

import (
	"github/umitgorgul/Shopping-Cart-REST-API/internal/models/cart"
	"github/umitgorgul/Shopping-Cart-REST-API/internal/models/product"
)

type CartService struct {
	repository cart.Repository
	proRepo    product.Repository
}

func NewCartService(repository cart.Repository, proRepo product.Repository) *CartService {
	return &CartService{repository: repository, proRepo: proRepo}
}

func (s *CartService) Create(req *CreateCartRequest) error {
	if err := CreateCartValidate(req); err != nil {
		return err
	}
	//map products and sum their prices and equalize them to cart total price
	var totalPrice float64
	var intervalVat int
	count8 := 1
	count18 := 1
	discountCount := 0
	for _, p := range req.CartProducts {
		product, err := product.FindByID(s.proRepo, p.ProductID)
		if err != nil {
			return err
		}
		if float64(p.Quantity) > 3.0 && discountCount == 0 {
			totalPrice = product.UnitPrice * 0.92 * float64(p.Quantity)
			discountCount += discountCount
		} else {
			totalPrice += product.UnitPrice * float64(p.Quantity)
		}
		intervalVat = product.Vat
		if intervalVat == 8 {
			count8 += count8
		} else if intervalVat == 18 {
			count18 += count18
		}
		if count8 > 3 && discountCount == 0 {
			totalPrice = totalPrice * 0.9
			discountCount += discountCount
		}
		if count18 > 3 && discountCount == 0 {
			totalPrice = totalPrice * 0.85
			discountCount += discountCount
		}

	}

	newCart := &cart.Cart{
		CustomerID:   req.CustomerID,
		CartProducts: req.CartProducts,
		TotalPrice:   totalPrice,
	}
	return cart.Create(s.repository, newCart)
}

func (s *CartService) Update(req *UpdateCartRequest) error {
	if err := UpdateCartValidate(req); err != nil {
		return err
	}
	var totalPrice float64
	for _, p := range req.CartProducts {
		product, err := product.FindByID(s.proRepo, p.ProductID)
		if err != nil {
			return err
		}
		totalPrice += product.UnitPrice * float64(p.Quantity)
	}
	newCart := &cart.Cart{
		ID:           req.ID,
		CustomerID:   req.CustomerID,
		CartProducts: req.CartProducts,
		TotalPrice:   totalPrice,
	}
	return cart.Update(s.repository, newCart)
}

func (s *CartService) Delete(req *DeleteCartRequest) error {
	if err := DeleteCartValidate(req); err != nil {
		return err
	}
	return cart.Delete(s.repository, req.ID)
}

func (s *CartService) FindByID(req *FindByIDRequest) (*cart.Cart, error) {
	if err := FindByIDValidate(req); err != nil {
		return nil, err
	}
	return cart.FindByID(s.repository, req.ID)
}
