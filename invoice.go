package main

import "strconv"

// ParseCSV is parsing the CSV data from Clockify
func (inv *Invoice) ParseCSV(records [][]string) error {
	// Parse first line data
	inv.Currency = "USD"
	inv.InvoiceDate = ""
	inv.InvoiceUntil = ""
	inv.InvoiceNumber = 0 // AAAA MM NBR
	for i := 1; i < len(records); i++ {
		inv.Projects = append(inv.Projects, Project{
			Index:  strconv.Itoa(i),
			Name:   records[i][0],
			Time:   records[i][4],
			Amount: records[i][5],
		})
	}
	rate, err := GetRate(inv.Currency)
	if err != nil {
		// TODO: err managment
		return err
	}
	inv.Rate = rate
	return nil
}
