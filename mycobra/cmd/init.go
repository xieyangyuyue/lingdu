package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// 定义initCmd命令
var initCmd = &cobra.Command{
	Use:   "add",        // 命令使用名称
	Short: "short init", // 简短描述
	Long:  "Long init",  // 详细描述
	Run: func(cmd *cobra.Command, args []string) {
		// 命令开始执行
		fmt.Println("run init cmd begin")

		// 打印所有flag的值
		fmt.Println(
			cmd.Flags().Lookup("viper").Value,           // viper标志值
			cmd.Flags().Lookup("author").Value,          // author标志值
			cmd.Flags().Lookup("config").Value,          // config标志值
			viper.GetString("author"),                   // 从viper获取author值
			cmd.Flags().Lookup("license").Value,         // license标志值
			cmd.Parent().Flags().Lookup("source").Value, // 父命令的source标志值
		)

		// 命令结束执行
		fmt.Println("run init cmd end")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
