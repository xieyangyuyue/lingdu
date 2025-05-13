package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

// 定义自定义参数检查命令
var cusArgsCheckCmd = &cobra.Command{
	Use: "cusargs", // 命令使用名称
	// 参数验证函数
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 { // 检查最少参数数量
			return errors.New("至少输入一个参数") // 返回参数不足错误
		}
		if len(args) > 2 { // 检查最多参数数量
			return errors.New("最多输入两个参数") // 返回参数过多错误
		}
		return nil // 参数数量正确
	},
	// 命令执行函数
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run cusagrs cmd begin") // 命令开始执行
		fmt.Println(args)                    // 打印输入参数
		fmt.Println("run cusagrs cmd end")   // 命令结束执行
	},
}

// 定义参数检查命令，该命令只接受特定的有效参数
var argsCheckCmd = &cobra.Command{
	Use: "args", // 命令使用名称

	// 参数验证设置：
	Args: cobra.OnlyValidArgs, // 只接受ValidArgs列表中的参数

	// 该命令可接受的有效参数列表
	ValidArgs: []string{"xieyang", "yuyue", "yuyi"},

	// 命令执行函数
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("命令开始执行")        // 打印开始信息
		fmt.Println("接收到的参数:", args) // 打印接收到的参数
		fmt.Println("命令执行结束")        // 打印结束信息
	},
}

// 初始化函数，用于向根命令注册子命令
func init() {
	rootCmd.AddCommand(cusArgsCheckCmd) // 添加自定义参数检查命令
	rootCmd.AddCommand(argsCheckCmd)    // 添加预定义参数检查命令
}
