/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	stopwatch "github.com/riku-smile/go-timer/pkg"

	"github.com/spf13/cobra"
)

// stopwatchCmd represents the stopwatch command
var stopwatchCmd = &cobra.Command{
	Use:   "stopwatch",
	Short: "ストップウォッチの機能を提供しています。",
	Long: `ストップウォッチの機能を提供しています。

	go-timer stopwatch で利用可能です。
	終了時はCtrl+Cを入力してください。`,
	Run: func(cmd *cobra.Command, args []string) {
		// シグナルを受け取るチャネルを作成
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		// ストップウォッチ生成
		watch := stopwatch.NewStart()
		println("計測開始>>>\nCtrl+Cで終了")
		// Ctrl + C が押下されたとき、以下の処理を実施
		<-quit
		// ストップウォッチ終了処理
		watch.Stop()
		fmt.Printf("\n<<<経過時間: %v\n", watch.Show())
	},
}

func init() {
	rootCmd.AddCommand(stopwatchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopwatchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopwatchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
