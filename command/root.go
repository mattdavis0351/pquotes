package command

import (
	"github.com/mattdavis0351/pquotes/responses"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(responses.GetRandomQuote)
}

// RootCmd needs a comment so you leave me alone
var RootCmd = &cobra.Command{
	Use:   "tech-quote",
	Short: "tech-quote inspires you",
	Long:  "tech-quote returns quotes to your terminal to inspire you to code better",
}
