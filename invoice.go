package main

import "fmt"

// Project is
type Project struct {
	Name   string
	Time   string
	Amount string
}

// Client is
type Client struct {
	Name string
}

// Invoice is
type Invoice struct {
	Client  Client
	Project []Project
}

// Parse is
func (inv *Invoice) Parse(records [][]string) error {
	for i := 1; i < len(records); i++ {
		inv.Project = append(inv.Project, Project{
			Name:   records[i][0],
			Time:   records[i][4],
			Amount: records[i][5],
		})
	}
	return nil
}

// Debug is
func (inv *Invoice) Debug() error {
	for i := 0; i < len(inv.Project); i++ {
		fmt.Printf(`
			Name: %s
			Time: %s
			Amount: %s`,
			inv.Project[i].Name, inv.Project[i].Time, inv.Project[i].Amount)
	}
	return nil
}
