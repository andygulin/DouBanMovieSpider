package cmd

import (
	"DouBanMovieSpider/spider"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "DouBanMovieSpider",
	Short: "Spider Douban movie/TV show information\n",
	Long:  "Spider Douban movie/TV show information\n",
}

func Execute() {
	rootCmd.AddCommand(infoCmd, storeCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print spider information.",
	Long:  "Print spider information.",
	Run: func(cmd *cobra.Command, args []string) {
		subjectId := args[0]
		request := spider.Request{SubjectId: subjectId}
		fmt.Println(request)
	},
}

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store spider information to MongoDB.",
	Long:  "Store spider information to MongoDB.",
	Run: func(cmd *cobra.Command, args []string) {
		subjectId := args[0]
		request := spider.Request{SubjectId: subjectId}
		fmt.Println(request)
	},
}
