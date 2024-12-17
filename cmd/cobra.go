package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "A simple CLI app",
	Long:  "This is a simple CLI app built using Cobra.",
	Run: func(cmd *cobra.Command, args []string) {
		// 默认执行命令
		fmt.Println("Welcome to mycli!")
	},
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Print a hello message",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			name = "World"
		}
		fmt.Printf("Hello, %s!\n", name)
	},
}

var goodbyeCmd = &cobra.Command{
	Use:   "goodbye",
	Short: "Print a goodbye message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Goodbye!")
	},
}

func main() {

	// 设置子命令
	helloCmd.Flags().String("name", "", "Name to greet")
	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(goodbyeCmd)

	// 执行根命令
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
