package main

import (
	"github.com/spf13/cobra"
	"go-tour/internal/word"
	"log"
	"strings"
)
const (
	MODE_UPPER = iota + 1
	MODE_LOWER
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE
	MODE_CAMELCASE_TO_UNDERSCORE
)

var mode int8
var str string
var desc = strings.Join([]string{
	"改子命令支持各种单词格式转换，模式如下:",
	"1:全部单词转为大写",
	"2:全部单词转为小写",
	"3:下划线单词转为大写驼峰单词",
	"4:下划线单词转为小写驼峰单词",
	"5:驼峰单词转为下划线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",   //子命令命令标示
	Short: "单词格式转换", //简短说明
	Long:  desc,     //help 完整说明
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.TOUpper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = word.UnderScoreToUpperCamelCase(str)
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = word.UnderSCoreToLowerCamelCase(str)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式,请执行 help word 查看帮助文档")
		}
		log.Printf("输出结果: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
	//varP 系列方法中, 第一个参数为需要绑定的变量, 第二个参数为接受该参数的完整命令标志
	//第三个参数为对应的短标识, 第四个为默认值, 第五个参数为使用说明
}
