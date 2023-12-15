package usecase

import (
	"github.com/gocolly/colly/v2"
	"github.com/meifanora/web-scaper-go/config"
	"github.com/meifanora/web-scaper-go/entity"
	"strconv"
)

var (
	maxProducts     = 100
	numberOfThreads = 5
)

type ScraperMobilePhoneProductUseCaseImpl struct {
	colly         *colly.Collector
	configuration config.Config
}

func NewScraperMobilePhoneProductUseCase(configuration config.Config) ScraperMobilePhoneProductUseCase {
	return &ScraperMobilePhoneProductUseCaseImpl{
		colly: colly.NewCollector(
			colly.Async(true),
		),
		configuration: configuration,
	}
}

func (c ScraperMobilePhoneProductUseCaseImpl) ScrapeMobilePhoneProducts() ([]entity.Product, error) {
	var products []entity.Product
	// URL to scrape
	url := "https://www.tokopedia.com/p/handphone-tablet/handphone"

	c.colly.Limit(&colly.LimitRule{
		// limit the parallel requests to 4 request at a time
		Parallelism: 2,
	})

	// Create a wait group to wait for all goroutines to finish
	//var wg sync.WaitGroup

	// setting a valid User-Agent header
	c.colly.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	// Set up callback to extract product information
	c.colly.OnHTML(".e1nlzfl2", func(e *colly.HTMLElement) {
		product := extractProduct(e)
		products = append(products, product)
	})

	for i := 1; i <= 5; i++ {
		c.colly.Visit(url + "?page=" + strconv.Itoa(i))
	}

	// wait for tColly to visit all pages
	c.colly.Wait()

	return products, nil
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
