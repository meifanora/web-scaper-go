package command

import (
	"context"
	"fmt"
	"github.com/meifanora/web-scaper-go/usecase"
	"github.com/spf13/cobra"
)

func ScrapeMobilePhoneProducts(scraperService usecase.ScraperMobilePhoneProductUseCase, productService usecase.ProductUseCase, csvService usecase.CSVWriteUseCase) *cobra.Command {
	return &cobra.Command{
		Use:   "scrape-mobile-phone-products",
		Short: "Scrape mobile phone products",
		Long:  "Scrape mobile phone products from Tokopedia",
		RunE: func(cmd *cobra.Command, args []string) error {
			return scrapeMobilePhoneProducts(scraperService, productService, csvService)
		},
	}
}

func scrapeMobilePhoneProducts(scraperService usecase.ScraperMobilePhoneProductUseCase, productService usecase.ProductUseCase, csvService usecase.CSVWriteUseCase) error {
	products, err := scraperService.ScrapeMobilePhoneProducts()
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = productService.InsertProducts(ctx, products)
	if err != nil {
		return err
	}

	err = csvService.WriteMobilePhoneProductCSV(products)
	if err != nil {
		return err
	}

	fmt.Println("Scraping mobile phone products has been successfully completed!")

	return nil
}
