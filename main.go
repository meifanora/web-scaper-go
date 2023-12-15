package main

import (
	"github.com/meifanora/web-scaper-go/command"
	"github.com/meifanora/web-scaper-go/config"
	"github.com/meifanora/web-scaper-go/db"
	"github.com/meifanora/web-scaper-go/repository"
	"github.com/meifanora/web-scaper-go/usecase"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	config := config.New()
	database := db.NewDatabase(config)

	productRepository := repository.NewProductRepository(database)
	productUseCase := usecase.NewProductUseCase(productRepository, database)
	scraperUseCase := usecase.NewScraperMobilePhoneProductUseCase(config)
	csvUseCase := usecase.NewCSVWriteUseCase()

	rootCmd := &cobra.Command{
		Use:   "command-cli",
		Short: "Example CLI command in Go using cobra.",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	rootCmd.AddCommand(
		command.ScrapeMobilePhoneProducts(scraperUseCase, productUseCase, csvUseCase),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
