/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wudaoluo/golog"
	"github.com/wudaoluo/sonic/common"
	v1 "github.com/wudaoluo/sonic/endpoint/v1"
	"github.com/wudaoluo/sonic/service"

	"github.com/spf13/cobra"
)

// logicCmd represents the logic command
var logicCmd = &cobra.Command{
	Use:   "logic",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		golog.Info("logic service start...")
		service.NewLogicService().Start(context.Background())

		router := gin.Default()

		v1.LogicV1Router(router)

		go func() {
			err := router.Run(common.GetConf().Logic.Addr)
			if err != nil {
				golog.Error("auth service start faild","err",err)
				panic(err)
			}
		}()
	},
}

func init() {
	rootCmd.AddCommand(logicCmd)
}
