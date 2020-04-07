package cmd

import (
	"crit-go/gocrit"
	//"fmt"
	//"errors"
	//"flag"

	"github.com/spf13/cobra"
)

// xCmd represents the x command
var xCmd = &cobra.Command{
	Use:   "x",
	Short: "Explore image dir",
	Long:  `Explore image dir with option such as ps,fds,mems,rss`,
	Run: func(cmd *cobra.Command, args []string) {
		gocrit.Explore(dir, what)
	},
}

func init() {
	// rootCmd.AddCommand(xCmd)
	xCmd.Flags().StringVarP(&dir, "dir", "", "", "location/or the image")
	xCmd.MarkFlagRequired("dir")
	xCmd.Flags().StringVarP(&what, "what", "", "", "choose between {ps,fss,mems,rss}")
	xCmd.MarkFlagRequired("what")
	//xCmd.Flags().BoolP("fds", "", false, "file directory structure of image")
	//xCmd.Flags().BoolP("mems", "", false, "memory map of the image")
	/*
		rootCmd.AddCommand(xCmd)
		xCmd.Flags().Bool("dir", "", "")
		xCmd.Flags().Bool("ps", "", false, 'process structure of image')
		xCmd.Flags().Bool("fds", "", false, 'file directory structure of image')
		xCmd.Flags().Bool("mems", "", false, 'memory map of the image')
	*/
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// xCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// xCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
