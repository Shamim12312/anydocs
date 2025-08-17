package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/AstraBert/anydocs/ai"
	"github.com/AstraBert/anydocs/docs"
	"github.com/briandowns/spinner"
	"github.com/common-nighthawk/go-figure"
	"github.com/rvfet/rich-go"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "anydocs",
	Short: "anydocs is a CLI tool for fetching documentations from any endpoint returning plain text.",
	Long:  "anydocs is a CLI tool for fetching documentations from any endpoint returning plain text: it is designed to save the content of these documents into instructions files for local coding agents (like CLAUDE.md or AGENTS.md)",
	Run: func(cmd *cobra.Command, args []string) {
		logo := figure.NewColorFigure("anydocs", "larry3d", "red", true)
		logo.Print()
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing anydocs '%s'\n", err)
		os.Exit(1)
	}
}

var path string
var urls string
var summary bool

var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Aliases: []string{"f"},
	Short:   "Fetch documentation content",
	Long:    "Fetch documentation contant by passing the endpoint URLs (comma-separated, flag -u,--urls) and the path to which you would like to save this documentation (flag -p, --path). Optionally, you can also decide to produce an AI summary of the documentation (flag -s, --summary).",
	Run: func(cmd *cobra.Command, args []string) {
		logo := figure.NewColorFigure("anydocs", "larry3d", "red", true)
		logo.Print()

		if urls == "" {
			fmt.Fprintf(os.Stderr, "Error: URLs are required. Use -u or --urls flag.\n")
			os.Exit(1)
		}
		if path == "" {
			fmt.Fprintf(os.Stderr, "Error: Path is required. Use -p or --path flag.\n")
			os.Exit(1)
		}

		urls := strings.Split(urls, ",")
		// Clean up URLs (remove whitespace)
		for i, url := range urls {
			urls[i] = strings.TrimSpace(url)
		}
		s := spinner.New(spinner.CharSets[11], 1*time.Millisecond)
		s.Start()
		content := docs.FetchMany(urls)
		if summary {
			summary, err := ai.AnthropicResponse(content)
			if err != nil {
				rich.Error("There was an error while producing the AI summary, defaulting to writing the whole fetched content.\nError: " + err.Error() + "❌")
			} else {
				rich.Info("AI summary successfully produced✅")
				content = summary
			}
		}
		fileError := docs.WriteFileContent(path, content)
		if fileError != nil {
			s.Stop()
			os.Exit(1)
		} else {
			s.Stop()
			os.Exit(0)
		}
	},
}

func init() {
	fetchCmd.Flags().StringVarP(&urls, "urls", "u", "", "Pass a set of llms.txt endpoints, comma separated (e.g. 'https://docs.llamaindex.ai/en/latest/llms.txt,https://raw.githubusercontent.com/AstraBert/anydocs/main/README.md')")
	fetchCmd.Flags().StringVarP(&path, "path", "p", "", "Pass the path you want to save your files at")
	fetchCmd.Flags().BoolVarP(&summary, "summary", "s", false, "Use this flag if you want to enable AI summary of fetched documentation.")

	fetchCmd.MarkFlagRequired("urls")
	fetchCmd.MarkFlagRequired("path")

	rootCmd.AddCommand(fetchCmd)
}
