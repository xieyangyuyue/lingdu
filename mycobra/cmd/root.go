package cmd // Package cmd 定义包名为cmd

import (
	"fmt"                    // 导入格式化输出包
	"github.com/spf13/cobra" // 导入cobra命令行工具库
	"github.com/spf13/viper"
	"os"
)

// rootCmd 是程序的根命令，当不带子命令直接运行时执行
var rootCmd = &cobra.Command{
	// 命令的使用方法提示
	Use: "root",

	// 命令的简短描述（通常在帮助列表中显示）
	Short: "short desc",

	// 命令的详细描述（通常在单独查看命令帮助时显示）
	Long: "Long desc",

	// 命令执行时调用的函数
	Run: func(cmd *cobra.Command, args []string) {
		// 打印命令开始执行的提示
		fmt.Println("root cmd run begin")

		// 这里应该添加打印flag参数的代码
		// 当前只是一个占位注释
		// 打印所有flag的当前值
		fmt.Println(
			// 获取并打印viper标志的值
			cmd.Flags().Lookup("viper").Value,
			// 获取并打印author标志的值
			cmd.PersistentFlags().Lookup("author").Value,
			// 获取并打印config标志的值
			cmd.PersistentFlags().Lookup("config").Value,
			// 获取并打印license标志的值
			cmd.PersistentFlags().Lookup("license").Value,
			// 获取并打印source标志的值
			cmd.Flags().Lookup("source").Value,
		)
		// 打印命令结束执行的提示
		// 打印从viper获取的配置值
		fmt.Println("----------------------------------")
		fmt.Println(viper.GetString("author"))  // 获取并打印author配置值
		fmt.Println(viper.GetString("license")) // 获取并打印license配置值
		fmt.Println("root cmd run end")
	},
	TraverseChildren: true,
}

// Execute 执行根命令
// 通常从main.main()调用，整个程序只需要调用一次rootCmd
func Execute() {
	// 执行根命令并处理可能的错误
	if err := rootCmd.Execute(); err != nil {
		// 这里通常会处理执行错误
		fmt.Println(err)
	}
}

// 定义全局变量：配置文件路径
var cfgFile string

// 定义全局变量：用户许可证信息
var userLicense string

// 初始化函数 - 用于设置命令行标志
func init() {
	cobra.OnInitialize(initConfig)
	// 添加持久化标志（布尔型）：是否使用viper配置
	// name: 标志名
	// value: 默认值
	// usage: 使用说明
	rootCmd.PersistentFlags().Bool("viper", true, "")

	// 添加持久化标志（字符串型）并设置缩写：
	// name: 标志全名
	// shorthand: 缩写字母
	// value: 默认值
	// usage: 使用说明
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "")

	// 添加持久化标志（字符串型）并将值绑定到变量：
	// &cfgFile: 绑定的变量指针
	// name: 标志名
	// value: 默认值
	// usage: 使用说明
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "")

	// 添加持久化标志（字符串型）并设置缩写，同时绑定到变量：
	// &userLicense: 绑定的变量指针
	// name: 标志全名
	// shorthand: 缩写字母
	// value: 默认值
	// usage: 使用说明
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "")

	// 添加本地标志（字符串型）并设置缩写：
	// name: 标志全名
	// shorthand: 缩写字母
	// value: 默认值
	// usage: 使用说明
	rootCmd.Flags().StringP("source", "s", "", "")

	// 将命令行flag绑定到viper配置
	err := viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	if err != nil {
		return
	} // 绑定author flag到viper的author键
	viper.BindPFlag("license", rootCmd.PersistentFlags().Lookup("license")) // 绑定license flag到viper的license键

	// 设置默认值
	viper.SetDefault("author", "default author")   // 设置author的默认值
	viper.SetDefault("license", "default license") // 设置license的默认值
}

// initConfig 函数用于初始化配置
func initConfig() {
	// 如果指定了配置文件路径
	if cfgFile != "" {
		// 设置使用指定的配置文件
		viper.SetConfigFile(cfgFile)
	} else {
		// 获取用户主目录
		home, err := os.UserHomeDir()
		// 检查错误
		cobra.CheckErr(err)

		// 添加配置文件搜索路径（用户主目录）
		viper.AddConfigPath(home)
		// 设置配置文件类型为yaml
		viper.SetConfigType("yaml")
		// 设置配置文件名为.cobra
		viper.SetConfigName(".cobra")
	}

	// 自动加载环境变量
	viper.AutomaticEnv()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		// 输出错误信息
		fmt.Println(err)
	} else {
		// 输出正在使用的配置文件路径
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
