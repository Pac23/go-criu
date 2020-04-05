package cmd

import (
	//"fmt"
	"crit-go/gocrit"
	//"errors"
	"github.com/spf13/cobra"
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Convert Json To Binary",
	Long:  ``,
	/*
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires atleast one -i argument with image")
			}
			return nil
		},
	*/
	Run: func(cmd *cobra.Command, args []string) {
		gocrit.Encode(args)
	},
}

//var inloc, outloc string

func init() {
	//rootCmd.AddCommand(encodeCmd)
	encodeCmd.Flags().StringVarP(&inloc, "in", "i", "", "criu image in binary format to be decoded (stdin by default)")
	encodeCmd.Flags().StringVarP(&outloc, "out", "o", "", "output loc of the file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
