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
	"time"

	"github.com/riku-smile/go-timer/pkg"
	"github.com/spf13/cobra"
)

// timerCmd represents the timer command
var timerCmd = &cobra.Command{
	Use:   "timer",
	Short: "コマンドライン上で動作するタイマー機能です。",
	Long: `コマンドライン上で動作するタイマー機能です。

	半角数字を用いて、()h<>m[]s の形で利用してください。
	それぞれ、()時間<>分[]秒 という意味になります。
	例： 1h30m45s -> １時間３０分４５秒のタイマーを設定する。`,
	Run: func(cmd *cobra.Command, args []string) {
		timer := time.NewTimer(pkg.SetSeconds(args[0]) * time.Second)
		fmt.Printf("%vのタイマーを開始しました", pkg.SetTimer(args[0]))
		<-timer.C
		fmt.Printf("%vのタイマーが終了しました", pkg.SetTimer(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(timerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
