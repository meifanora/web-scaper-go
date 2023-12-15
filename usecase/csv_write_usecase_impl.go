package usecase

import (
	"encoding/csv"
	"github.com/meifanora/web-scaper-go/entity"
	"os"
	"strconv"
)

type CSVWriteUseCaseImpl struct {
}

func NewCSVWriteUseCase() CSVWriteUseCase {
	return &CSVWriteUseCaseImpl{}
}

func (c CSVWriteUseCaseImpl) WriteMobilePhoneProductCSV(products []entity.Product) error {
	file, err := os.Create("products.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Product Name", "Description", "Image Url", "Price", "Rating", "Merchant"}
	writer.Write(headers)

	for _, product := range products {
		record := []string{
			product.Name,
			product.Description,
			product.ImageUrl,
			product.Price,
			strconv.Itoa(product.Rating),
			product.Merchant,
		}

		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
