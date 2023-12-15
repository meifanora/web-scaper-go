package entity

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	Price       string `json:"price"`
	Rating      int    `json:"rating"`
	Merchant    string `json:"merchant"`

	// nolint:structcheck,unused
	tableName struct{} `pg:",discard_unknown_columns"`
}
