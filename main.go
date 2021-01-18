package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"encoding/csv"
	"encoding/json"

	"github.com/alecthomas/kong"
)

// Context struct
type Context struct {
	Debug bool
}

// CreateInvoiceCmd struct
type CreateInvoiceCmd struct {
	File string `arg name:"file" help:"Path to clokify csv file." type:"existingfile"`
	// Make a default value for settings file path
	Settings string `arg name:"settings" help:"Path to json config file." type:"existingfile"`
}

// Run is
func (cmd *CreateInvoiceCmd) Run(ctx *Context) error {
	// CSV File
	csvFile, err := ioutil.ReadFile(cmd.File)
	if err != nil {
		fmt.Println("Couldn't open CSV file")
		log.Fatal(err)
		return err
	}

	csvReader := csv.NewReader(strings.NewReader(string(csvFile)))
	csvData, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Couldn't read from CSV")
		log.Fatal(err)
		return err
	}

	var invoice Invoice
	// TODO: Get current date
	// Parse file name to get the period
	err = invoice.ParseCSV(csvData)
	if err != nil {
		fmt.Println("Couldn't parse CSV")
		log.Fatal(err)
		return err
	}

	// Settings
	settingsFile, err := ioutil.ReadFile(cmd.Settings)
	if err != nil {
		fmt.Println("Couldn't open Settings file")
		log.Fatal(err)
		return err
	}
	err = json.Unmarshal(settingsFile, &invoice.Settings)
	if err != nil {
		fmt.Println("Couldn't unmarshal Settings")
		log.Fatal(err)
		return err
	}

	// Template
	tmpl := template.Must(template.ParseFiles("templates/template.html"))
	f, err := os.OpenFile("./output/filled.html", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		fmt.Println("Couldn't open Template")
		log.Fatal(err)
		return err
	}
	err = tmpl.Execute(f, invoice)
	if err != nil {
		fmt.Println("Couldn't Execute Template")
		log.Fatal(err)
		return err
	}
	f.Close()

	// Build pdf
	path, err := exec.LookPath("node")
	if err != nil {
		log.Fatal("node is requiered")
	}
	err = exec.Command(path, "main.js").Run()
	if err != nil {
		return fmt.Errorf("Error exec: %v", err)
	}

	return nil
}

// CLI struct
var CLI struct {
	Debug bool `help:"Enable debug mode."`

	Create CreateInvoiceCmd `cmd help:"Create an invoice from a clockify csv report file."`
}

func main() {
	ctx := kong.Parse(&CLI)
	err := ctx.Run(&Context{
		Debug: CLI.Debug,
	})
	ctx.FatalIfErrorf(err)
}
