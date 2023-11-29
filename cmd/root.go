package cmd

import (
	. "DouBanMovieSpider/service"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var rootCmd = &cobra.Command{
	Use:   "DouBanMovieSpider",
	Short: "Spider Douban movie/TV show information",
	Long:  "Spider Douban movie/TV show information",
}

func Execute() {
	rootCmd.AddCommand(subjectCmd, commentCmd, reviewCmd, photoCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const (
	cmdInfo  = "info"
	cmdFile  = "file"
	cmdStore = "store"
	cmdPhoto = "photo"
	cmdQuery = "query"
)

var subCommands = []string{cmdInfo, cmdFile, cmdStore, cmdPhoto, cmdQuery}

func inArray(subCommand string) bool {
	exist := false
	for _, command := range subCommands {
		if command == subCommand {
			exist = true
			break
		}
	}
	return exist
}

var infoHandle Handle
var fileHandle Handle
var storeHandle Handle
var request Request

var storeQuery Query

func init() {
	infoHandle = new(InfoHandle)
	fileHandle = new(FileHandle)
	storeHandle = new(StoreHandle)
}

var subjectCmd = &cobra.Command{
	Use:   "subject",
	Short: "Subject Info.",
	Long:  "Subject Info.",
	Run: func(cmd *cobra.Command, args []string) {
		subCommand := args[0]
		subjectId := args[1]
		if !inArray(subCommand) {
			panic(fmt.Sprintf("Command %s is incorrect. Please use %v", subCommand, subCommands))
		}

		var output Result
		var err error
		if subCommand != cmdQuery {
			request = Request{SubjectId: subjectId}
			response, err := request.SpiderSubject()
			if err != nil {
				panic(err)
			}
			if subCommand == cmdInfo {
				output, err = infoHandle.HandleSubject(response)
			}
			if subCommand == cmdFile {
				output, err = fileHandle.HandleSubject(response)
			}
			if subCommand == cmdStore {
				output, err = storeHandle.HandleSubject(response)
			}
		} else {
			pageNo, _ := strconv.ParseInt(args[2], 10, 64)
			pageSize, _ := strconv.ParseInt(args[3], 10, 64)

			storeQuery = &StoreQuery{
				SubjectId: subjectId,
				PageNo:    pageNo,
				PageSize:  pageSize,
			}
			output, err = storeQuery.QuerySubject()
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(output)
	},
}

var commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "Comment Info.",
	Long:  "Comment Info.",
	Run: func(cmd *cobra.Command, args []string) {
		subCommand := args[0]
		subjectId := args[1]
		if !inArray(subCommand) {
			panic(fmt.Sprintf("Command %s is incorrect. Please use %v", subCommand, subCommands))
		}

		var output Result
		var err error
		if subCommand != cmdQuery {
			request = Request{SubjectId: subjectId}
			response, err := request.SpiderComment()
			if err != nil {
				panic(err)
			}
			if subCommand == cmdInfo {
				output, err = infoHandle.HandleComment(response)
			}
			if subCommand == cmdFile {
				output, err = fileHandle.HandleComment(response)
			}
			if subCommand == cmdStore {
				output, err = storeHandle.HandleComment(response)
			}
		} else {
			pageNo, _ := strconv.ParseInt(args[2], 10, 64)
			pageSize, _ := strconv.ParseInt(args[3], 10, 64)

			storeQuery = &StoreQuery{
				SubjectId: subjectId,
				PageNo:    pageNo,
				PageSize:  pageSize,
			}
			output, err = storeQuery.QueryComment()
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(output)
	},
}

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Review Info.",
	Long:  "Review Info.",
	Run: func(cmd *cobra.Command, args []string) {
		subCommand := args[0]
		subjectId := args[1]
		if !inArray(subCommand) {
			panic(fmt.Sprintf("Command %s is incorrect. Please use %v", subCommand, subCommands))
		}

		var output Result
		var err error
		if subCommand != cmdQuery {
			request = Request{SubjectId: subjectId}
			response, err := request.SpiderReview()
			if err != nil {
				panic(err)
			}
			if subCommand == cmdInfo {
				output, err = infoHandle.HandleReview(response)
			}
			if subCommand == cmdFile {
				output, err = fileHandle.HandleReview(response)
			}
			if subCommand == cmdStore {
				output, err = storeHandle.HandleReview(response)
			}
		} else {
			pageNo, _ := strconv.ParseInt(args[2], 10, 64)
			pageSize, _ := strconv.ParseInt(args[3], 10, 64)

			storeQuery = &StoreQuery{
				SubjectId: subjectId,
				PageNo:    pageNo,
				PageSize:  pageSize,
			}
			output, err = storeQuery.QueryReview()
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(output)
	},
}

var photoCmd = &cobra.Command{
	Use:   "photo",
	Short: "Photo Info.",
	Long:  "Photo Info.",
	Run: func(cmd *cobra.Command, args []string) {
		subCommand := args[0]
		subjectId := args[1]
		if !inArray(subCommand) {
			panic(fmt.Sprintf("Command %s is incorrect. Please use %v", subCommand, subCommands))
		}

		var output Result
		var err error
		if subCommand != cmdQuery {
			request = Request{SubjectId: subjectId}
			response, err := request.SpiderPhoto()
			if err != nil {
				panic(err)
			}
			if subCommand == cmdInfo {
				output, err = infoHandle.HandlePhoto(response)
			}
			if subCommand == cmdFile {
				output, err = fileHandle.HandlePhoto(response)
			}
			if subCommand == cmdStore {
				output, err = storeHandle.HandlePhoto(response)
			}
		} else {
			pageNo, _ := strconv.ParseInt(args[2], 10, 64)
			pageSize, _ := strconv.ParseInt(args[3], 10, 64)

			storeQuery = &StoreQuery{
				SubjectId: subjectId,
				PageNo:    pageNo,
				PageSize:  pageSize,
			}
			output, err = storeQuery.QueryPhoto()
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(output)
	},
}
