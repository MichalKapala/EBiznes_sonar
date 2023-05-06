package main

// type Product struct {
// 	Id uint `gorm:"primaryKey"`
// 	Name        string  `json:"name"`
// 	Description string  `json:"description"`
// 	Price       float64 `json:"price"`
// }

type Product struct {
	Id uint 	`gorm:"primaryKey"`
    Name        string
    Description string
    Price       float64
	Category	uint
}

type Category struct {
	Id uint 	`gorm:"primaryKey"`
    Name        string
}

type Bucket struct {
	ProductId 	uint 	`gorm:"primaryKey"`
    Quantity    uint
}

type Confirmation struct {
	Name 	string 	
    Address    string
    City    string
    State   string
    Zip     string
    Delivery string
    CardNo  string
    CardExp string
    Cvv     string

}