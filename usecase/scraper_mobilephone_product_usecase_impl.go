package usecase

import (
	"github.com/gocolly/colly/v2"
	"github.com/meifanora/web-scaper-go/config"
	"github.com/meifanora/web-scaper-go/entity"
	"strconv"
	"sync"
)

var (
	url             = "https://www.tokopedia.com/p/handphone-tablet/handphone" // URL to scrape
	maxProducts     = 100
	numberOfThreads = 10
)

type ScraperMobilePhoneProductUseCaseImpl struct {
	colly         *colly.Collector
	configuration config.Config
}

func NewScraperMobilePhoneProductUseCase(configuration config.Config) ScraperMobilePhoneProductUseCase {
	return &ScraperMobilePhoneProductUseCaseImpl{
		colly:         colly.NewCollector(),
		configuration: configuration,
	}
}

func (c ScraperMobilePhoneProductUseCaseImpl) ScrapeMobilePhoneProducts() ([]entity.Product, error) {
	var products []entity.Product

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a buffered channel to limit the number of simultaneous requests
	ch := make(chan entity.Product, numberOfThreads)

	// setting a valid User-Agent header
	c.colly.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	// Set up the collector to extract product information
	c.colly.OnHTML(".e1nlzfl2", func(e *colly.HTMLElement) {
		ch <- extractProduct(e)
	})

	// Start the scraping process
	for i := 1; i <= numberOfThreads; i++ {
		wg.Add(1) // // Increment the wait group counter
		go c.scrapeWorker(i, &wg)
	}

	// Close the result channel when all workers are done
	go func() {
		wg.Wait() // Wait for all workers to finish
		close(ch)
	}()

	for product := range ch {
		products = append(products, product)
		if len(products) >= maxProducts {
			break
		}
	}

	return products, nil
}

func (c ScraperMobilePhoneProductUseCaseImpl) scrapeWorker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	c.colly.Visit(url + "?page=" + strconv.Itoa(i))
}

func extractProduct(e *colly.HTMLElement) entity.Product {
	productName := e.ChildText("div.css-11s9vse > span.css-20kt3o")
	description := e.ChildText("div.css-11s9vse > span.css-20kt3o")
	imageUrl := e.ChildAttr("div.css-1g5og91 > img", "src") // base64 because generated from html
	price := e.ChildText("div.css-pp6b3e > span.css-o5uqvq")

	var rating int
	e.ForEach(".css-177n1u3", func(index int, j *colly.HTMLElement) {
		rating++
	})

	var merchantName string
	counterIndexMerchant := 1
	e.ForEach("div.css-vbihp9 > span.css-ywdpwd", func(index int, k *colly.HTMLElement) {
		if counterIndexMerchant%2 == 0 {
			merchantName = k.Text
		}
		counterIndexMerchant++
	})

	product := entity.Product{
		Name:        productName,
		Description: description,
		ImageUrl:    imageUrl,
		Price:       price,
		Rating:      rating,
		Merchant:    merchantName,
	}

	return product
}
