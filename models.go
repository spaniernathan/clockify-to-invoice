package main

// UserInfo struct
type UserInfo struct {
	CompagnyName string `json:"CompagnyName"`
	Name         string `json:"Name"`
	PhoneNumber  string `json:"PhoneNumber"`
	Email        string `json:"Email"`
}

// BanqueInfo struct
type BanqueInfo struct {
	BanqueName string `json:"BanqueName"`
	SwiftBic   string `json:"SwiftBic"`
	IBAN       string `json:"IBAN"`
}

// Address struct
type Address struct {
	PostCode string `json:"Postcode"`
	Country  string `json:"Country"`
	Address  string `json:"Address"`
	City     string `json:"City"`
}

// User struct
type User struct {
	UserInfo   UserInfo   `json:"UserInfo"`
	Address    Address    `json:"UserAddress"`
	BanqueInfo BanqueInfo `json:"BanqueInfo"`
	TVANumber  string     `json:"TVANumber"`
	SIRET      string     `json:"SIRET"`
}

// Client struct
type Client struct {
	Name    string  `json:"ClientName"`
	Address Address `json:"ClientAddress"`
}

// Settings struct
type Settings struct {
	User   User   `json:"User"`
	Client Client `json:"BilledClient"`
}

// Project struct
type Project struct {
	Index      string
	Name       string
	Time       string
	Amount     string
	UnitAmount string
	Currency   string
}

// Invoice struct
type Invoice struct {
	InvoiceNumber uint64
	InvoiceDate   string //date
	InvoiceUntil  string //date
	Currency      string
	Rate          string
	Client        Client
	Projects      []Project
	Settings      Settings
}
