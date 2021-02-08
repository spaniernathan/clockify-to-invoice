package main

import (
	"fmt"
	"strconv"
	"time"
)

// ParseCSV is parsing the CSV data from Clockify
func (inv *Invoice) ParseCSV(records [][]string, from time.Time, to time.Time) error {
	// Parse first line data
	inv.Currency = "USD" // make configurable
	inv.InvoiceDate = fmt.Sprintf("%02d-%02d-%d", to.Day(), int(to.Month()), to.Year())
	inv.InvoiceNumber = fmt.Sprintf("%d%02d%d", to.Year(), int(to.Month()), 1) // TODO: change 1 to invoice number / setup local jsondb or something
	to = to.AddDate(0, 1, 0)
	inv.InvoiceUntil = fmt.Sprintf("%02d-%02d-%d", to.Day(), int(to.Month()), to.Year())
	totalAmount := 0.0
	for i := 1; i < len(records); i++ {
		amount, err := strconv.ParseFloat(records[i][4], 64)
		totalAmount += amount * 25
		if err != nil {
			return fmt.Errorf("Coudln't parse amount: %v", err)
		}
		inv.Projects = append(inv.Projects, Project{
			Index:      strconv.Itoa(i),
			Name:       fmt.Sprintf("%s (%s)", records[i][0], records[i][1]),
			Time:       fmt.Sprintf("%.2f", amount),
			Amount:     strconv.FormatFloat(amount*25, 'f', 2, 64),
			UnitAmount: "25.0",
		})
	}
	inv.TotalAmount = strconv.FormatFloat(totalAmount, 'f', 2, 64)
	rate, err := GetRate(inv.Currency)
	if err != nil {
		// TODO: err managment
		return fmt.Errorf("Coudln't fetch exchange rate: %v", err)
	}
	inv.Rate = rate
	return nil
}
