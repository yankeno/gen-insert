/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	output string
	table  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gen-insert",
	Short: "An application that generates SQL INSERT statements from a specific file format.",
	Long: `gen-insert is an application that generates SQL INSERT statements 
from files of a specific format.

Supported file formats are as follows:

- CSV
- TSV

Supported file encodings are as follows:

- UTF-8`,
	Args: cobra.MatchAll(cobra.ExactArgs(1), func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			return fmt.Errorf("file not exists: %s", args[0])
		}

		if table != "" {
			if matched, err := regexp.MatchString(`^[a-zA-Z_][a-zA-Z0-9_]*$`, table); !matched || err != nil {
				return fmt.Errorf("table name is invalid: %s", table)
			}
		}

		if ext := filepath.Ext(args[0]); !(ext == ".csv" || ext == ".tsv") {
			return fmt.Errorf(`file format is not supported.
Supported file formats: 

- CSV
- TSV`)
		}

		return nil
	}),
	Run: func(cmd *cobra.Command, args []string) {
		fullpath := args[0]
		f, err := os.Open(fullpath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		ext := filepath.Ext(fullpath)
		filename := filepath.Base(fullpath)
		tablename := strings.Replace(filename, ext, "", 1)
		if table != "" {
			tablename = table
		}

		dirname := filepath.Dir(output)
		if _, err := os.Stat(dirname); os.IsNotExist(err) {
			os.MkdirAll(dirname, 0755)
		}

		outputFilename := tablename + ".sql"
		if output != "" {
			outputFilename = filepath.Base(output)
		}
		outputBasename := strings.Replace(outputFilename, filepath.Ext(outputFilename), "", 1)

		sn := 0
		for {
			_, err := os.Stat(filepath.Join(dirname, outputFilename))
			if os.IsNotExist(err) {
				break
			}
			sn++
			outputFilename = fmt.Sprintf("%s(%d).sql", outputBasename, sn)
		}

		r := csv.NewReader(f)
		r.LazyQuotes = true
		if ext == ".tsv" {
			r.Comma = '\t'
		}

		headers, err := r.Read()
		if err != nil {
			fmt.Println("Error reading headers: ", err)
			return
		}

		o, err := os.Create(filepath.Join(dirname, outputFilename))
		if err != nil {
			log.Fatal(err)
		}
		defer o.Close()

		for {
			values := make([]string, 0, len(headers))
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			for i := 0; i < len(headers); i++ {
				if strings.EqualFold(record[i], "null") {
					values = append(values, "NULL")
					continue
				}
				if _, err := strconv.ParseFloat(record[i], 64); err != nil {
					values = append(values, fmt.Sprintf("'%s'", record[i]))
					continue
				}
				values = append(values, record[i])
			}
			_, err = fmt.Fprintf(o, "INSERT INTO `%s` (%s) VALUES (%s);\n", tablename, "`"+strings.Join(headers, "`, `")+"`", strings.Join(values, ", "))
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Printf("INSERT statement '%s' was generated successfully!\n", outputFilename)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.src.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "Output file name")
	rootCmd.Flags().StringVarP(&table, "table", "t", "", "Table name")
}
