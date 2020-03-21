package models

import (
	"context"

	"github.com/devarchi33/all-about-shoes-api/factory"
)

type ShoesCreateDto struct {
	BrandName   string  `json:"brand_name"`
	ProductName string  `json:"product_name"`
	StyleCode   string  `json:"style_code"`
	SizeFlag    string  `json:"size_flag"`
	Price       float64 `json:"price"`
	PriceUnit   string  `json:"price_unit"`
	ReleaseDate string  `json:"release_date"`
	CreatedBy   string  `json:"created_by"`
}

var (
	ShoesTableName = "cmm_shoes"
)

func (d *Shoes) TableName() string {
	return ShoesTableName
}

type Shoes struct {
	Id          int64   `json:"id"`
	BrandName   string  `json:"brand_name"`
	ProductName string  `json:"product_name"`
	StyleCode   string  `json:"style_code"`
	SizeFlag    string  `json:"size_flag"`
	Price       float64 `json:"price"`
	PriceUnit   string  `json:"price_unit"`
	ReleaseDate string  `json:"release_date"`
	Committed   `xorm:"extends"`
}

func (s *Shoes) Create(ctx context.Context) error {
	if _, err := factory.
		DB(ctx).
		Table(ShoesTableName).
		Insert(s); err != nil {
		return err
	}
	return nil
}

func (Shoes) TranslateForCreate(dto ShoesCreateDto) (Shoes, error) {
	return Shoes{
		BrandName:   dto.BrandName,
		ProductName: dto.ProductName,
		StyleCode:   dto.StyleCode,
		SizeFlag:    dto.SizeFlag,
		Price:       dto.Price,
		PriceUnit:   dto.PriceUnit,
		ReleaseDate: dto.ReleaseDate,
		Committed:   Committed{}.newCommitted(dto.CreatedBy),
	}, nil
}
