// Copyright © 2017 hajimeni
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tmaher/aws-parameter-store-helper/aws"
	"os"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "load stored parameter then export formatted string",
	Run: func(cmd *cobra.Command, args []string) {
		aws.LoadParameterStore(&loadFlag)
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return aws.CheckRequiredFlags(&loadFlag)
	},
}

var loadFlag aws.LoadFlag

func init() {
	RootCmd.AddCommand(loadCmd)
	loadCmd.SetOutput(os.Stderr)

	loadCmd.Flags().StringVarP(&loadFlag.Path, "path", "p", "", "Parameter Store Path, must starts with '/' ")
	loadCmd.Flags().StringVar(&loadFlag.Prefix, "prefix","", "Parameter Store Prefix. export KEY is removed prefix")
	loadCmd.Flags().StringVarP(&loadFlag.Template, "template", "t", "export {{ .Name }}='{{ .Value }}'", "export format template(Go Template)")
	loadCmd.Flags().StringVarP(&loadFlag.Delimiter,"delimiter", "d", ";", "Delimiter each keys")
	loadCmd.Flags().StringVarP(&loadFlag.Region, "region", "r", "", "AWS SDK region")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
