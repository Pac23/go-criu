package cmd

import (
	"crit-go/gocrit"
	//"errors"
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	//"os"
)

// xCmd represents the x command
var xCmd = &cobra.Command{
	Use:   "x",
	Short: "Explore image dir",
	Long:  `Explore image dir with option such as ps,fds,mems,rss`,
	/*
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("requires a x argument with image and another option such as ps,fds,mems,rss")
			}
			return nil
		},
	*/
	Run: func(cmd *cobra.Command, args []string) {
		err := gocrit.Explore(dir, what)
		if err != nil
		gocrit.check(err)
	},
}

var dir, what string

func init() {
	// rootCmd.AddCommand(xCmd)
	xCmd.Flags().StringVarP(&dir, "dir", "", "", "location/or the image")
	xCmd.MarkFlagRequired("dir")
	/*
		lastVal := &LastFlag{}
		flag.Var(lastVal, "flag", "takes last value given")
		//os.Args = []string{"exename", "-flag", "ps", "-flag", "fds", "-flag", "mems", "-flag", "-rss"}

		//fmt.Println("Final value:", lastVal.String())

		//args := os.Args[1:]
	*/
	xCmd.Flags().BoolVarP(&what, "what", "", "", "choose between {ps,fss,mems,rss}")
	xCmd.MarkFlagRequired("what")
	//xCmd.Flags().BoolP("fds", "", false, "file directory structure of image")
	//xCmd.Flags().BoolP("mems", "", false, "memory map of the image")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// xCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// xCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
