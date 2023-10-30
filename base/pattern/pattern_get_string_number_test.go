package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetStringNumber(t *testing.T) {
	input := "Plan: 1 to add, 0 to change, 10 to destroy."

	// 创建一个正则表达式来匹配数字
	regex := regexp.MustCompile(`\d+`)

	// 查找所有匹配的数字
	matches := regex.FindAllString(input, -1)

	// 打印匹配的数字
	for _, match := range matches {
		fmt.Println(match)
	}
}

func TestGetStringNumberForText(t *testing.T) {
	input := `This is a long text with multiple patterns.
    Plan: 123 to add, 456 to change, 789 to destroy.
    Another plan: 42 to add, 99 to change, 1234 to destroy.
    Some more text.
    Plan: 9876 to add, 5432 to change, 101 to destroy.`

	// 创建一个正则表达式来匹配类似 "Plan: 123 to add, 456 to change, 789 to destroy." 的模式
	regex := regexp.MustCompile(`Plan: \d+ to add, \d+ to change, \d+ to destroy`)

	// 查找所有匹配的字符串
	matches := regex.FindAllString(input, -1)

	// 打印匹配的字符串
	for _, match := range matches {
		fmt.Println(match)
	}
}

func TestSubTerraformPlan(t *testing.T) {
	input := `Some text before.古代诗歌
Terraform will perform the following actions:
gsajlgjsalkgjlkdsjglkjlkasjgl
sdgjl
sdgjlkdsjg
sdgjlkdsjgldsjg
sdgj
sd
\u001b[0m\u001b[1mPlan:\u001b[0m 1 to add, 0 to change, 0 to destroy.
大概是佳丽就是大概佳丽的世界观
`
	re := regexp.MustCompile("\x1b\\[[0-9;]*[a-zA-Z]")
	input = re.ReplaceAllString(input, "")
	fmt.Println(input)
}
