/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"bufio"
	"k8scopilot/utils"
	"os"

	"github.com/spf13/cobra"
)

const prePrompt = "你现在是一个 K8S Copilot，你要帮用户完成 K8S 相关的任务，当需要输出YAML文件的时候，只输出YAML内容即可，并且不要把YAML内容放在```代码块里。"

// chatgptCmd represents the chatgpt command
var chatgptCmd = &cobra.Command{
	Use:   "chatgpt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		startChat()
	},
}

func startChat() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("我是 k8s copilot，有什么需要帮助？")

	for {
		fmt.Print("> ")
		if scanner.Scan() {
			input := scanner.Text()
			if input == "exit" {
				fmt.Println("再见！")
				break
			}
			if input == "" {
				continue
			}
			response := processInput(input)
			fmt.Println(response)
		}
	}
}

func processInput(input string) {
	client, err := utils.NewOpenAIClient()
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}
	defer client.Close()

	response, err := client.SendMessage(prePrompt, input)
	if err != nil {
		fmt.Println("Error generating completion:", err)
		return
	}

	return response
}

func init() {
	askCmd.AddCommand(chatgptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatgptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatgptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
