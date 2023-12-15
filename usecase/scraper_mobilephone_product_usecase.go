package usecase

import "github.com/meifanora/web-scaper-go/entity"

type ScraperMobilePhoneProductUseCase interface {
	ScrapeMobilePhoneProducts() ([]entity.Product, error)
}