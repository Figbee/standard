package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gin后台管理系统",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s/n", `欢迎使用gin后台管理系统,可以使用-h查看命令`)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
