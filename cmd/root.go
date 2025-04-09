package cmd

import (
	. "DouBanMovieSpider/service"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

var rootCmd = &cobra.Command{
	Use:   "DouBanMovieSpider",
	Short: "抓取豆瓣电影",
	Long:  "抓取豆瓣电影信息（详情、图片、评论、影评），打印信息或将信息保存至文件和MongoDB.",
}

func Execute() {
	rootCmd.AddCommand(subjectCmd, commentCmd, reviewCmd, photoCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

const (
	// 打印信息
	cmdInfo = "info"
	// 保存至文件
	cmdFile = "file"
	// 保存至数据库
	cmdStore = "store"
	// 从数据库中查询
	cmdQuery = "query"
)

var subCommands = []string{cmdInfo, cmdFile, cmdStore, cmdQuery}

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

func toInt(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func checkSubCommands(subCommand string) {
	if !inArray(subCommand) {
		fmt.Println(fmt.Sprintf("命令 %s 不存在. 请使用 %v", subCommand, subCommands))
		os.Exit(1)
	}
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
	Short: "电影详情信息",
	Long:  "电影详情信息",
	Run: func(cmd *cobra.Command, args []string) {
		subCommand := args[0]
		checkSubCommands(subCommand)
		subjectId := args[1]

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
			storeQuery = &StoreQuery{
				SubjectId: subjectId,
			}
			output, err = storeQuery.QuerySubject()
		}
		if err != nil {
			panic(err)
		}
		log.Println(output)
	},
}

var commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "电影评论信息",
	Long:  "电影评论信息",
	Run: func(cmd *cobra.Command, args []string) {
		subCommand := args[0]
		checkSubCommands(subCommand)
		subjectId := args[1]

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
			pageNo, _ := toInt(args[2])
			pageSize, _ := toInt(args[3])

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
		log.Println(output)
	},
}

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "电影影评信息",
	Long:  "电影影评信息",
	Run: func(cmd *cobra.Command, args []string) {
		subCommand := args[0]
		checkSubCommands(subCommand)
		subjectId := args[1]

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
			pageNo, _ := toInt(args[2])
			pageSize, _ := toInt(args[3])

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
		log.Println(output)
	},
}

var photoCmd = &cobra.Command{
	Use:   "photo",
	Short: "电影图片信息",
	Long:  "电影图片信息",
	Run: func(cmd *cobra.Command, args []string) {
		subCommand := args[0]
		checkSubCommands(subCommand)
		subjectId := args[1]

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
			pageNo, _ := toInt(args[2])
			pageSize, _ := toInt(args[3])

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
		log.Println(output)
	},
}
