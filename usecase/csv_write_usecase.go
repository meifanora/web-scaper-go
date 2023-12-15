package usecase

import "github.com/meifanora/web-scaper-go/entity"

type CSVWriteUseCase interface {
	WriteMobilePhoneProductCSV(products []entity.Product) error
}
