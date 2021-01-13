package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"encoding/csv"

	"github.com/alecthomas/kong"
)

// Context in
type Context struct {
	Debug bool
}

// CreateInvoiceCmd in
type CreateInvoiceCmd struct {
	File string `arg name:"file" help:"Paths to clokify csv file." type:"existingfile"`
}

// Run in
func (cmd *CreateInvoiceCmd) Run(ctx *Context) error {
	data, err := ioutil.ReadFile(cmd.File)
	csvReader := csv.NewReader(strings.NewReader(string(data)))
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return err
	}
	var invoice Invoice
	fmt.Print(records)
	err = invoice.Parse(records)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_ = invoice.Debug()
	return nil
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Create CreateInvoiceCmd `cmd help:"Create an invoice from a clockify csv report file."`
}

func main() {
	ctx := kong.Parse(&cli)
	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&Context{
		Debug: cli.Debug,
	})
	ctx.FatalIfErrorf(err)
}
