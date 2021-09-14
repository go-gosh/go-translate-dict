package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/go-gosh/go-youdao-dict/api"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := api.Translate(args[0], api.WithTo(api.LangTypeEN))
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", b)
		fmt.Printf("https://www.youdao.com/w/eng/%s/\n", url.QueryEscape(args[0]))

		return nil
	},
}

func main() {
	err := root.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
