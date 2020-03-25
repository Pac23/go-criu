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
	"fmt"

	"github.com/spf13/cobra"
)

// xCmd represents the x command
var xCmd = &cobra.Command{
	Use:   "x",
	Short: "Explore image dir" ,
	Long: `Explore image dir with option such as ps,fds,mems,rss`,
	Run: func(cmd *cobra.Command, args []string) error {
    	if len(args) < 2 {
      		return errors.New("requires a x argument with image and another option such as ps,fds,mems,rss")
    	}
		Explore(args []string)
	}
}

func init() {
	rootCmd.AddCommand(xCmd)
	xCmd.Flags().Bool("dir", "", "")
	xCmd.Flags().Bool("ps", "", false, 'process structure of image')
	xCmd.Flags().Bool("fds", "", false, 'file directory structure of image')
	xCmd.Flags().Bool("mems", "", false, 'memory map of the image')	
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// xCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// xCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
