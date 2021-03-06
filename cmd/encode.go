/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("encode called")

		cusor, err := os.Open("./gfwlist.haowei.txt")
		if err != nil {
			log.Fatal(err)
		}
		content, err := ioutil.ReadAll(cusor)
		out := make([]byte, base64.StdEncoding.EncodedLen(len(content)))
		base64.StdEncoding.Encode(out, content)
		cusor, err = os.Create("./gfwlist.haowei.txt")
		if err != nil {
			log.Fatal(err)
		}
		writer := bufio.NewWriter(cusor)

		i := 0
		for _, b := range out {
			writer.WriteByte(b)
			i += 1
			if i%64 == 0 {
				writer.WriteByte('\n')
				writer.Flush()
			}
		}
		writer.Flush()
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
