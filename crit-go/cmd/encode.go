package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Convert Json To Binary",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) error {
    	if len(args) < 1 {
      		return errors.New("requires a -i argument with image")
    	}
		Encode(args []string)
	}
}

func init() {
	rootCmd.AddCommand(encodeCmd)
	encodeCmd.Flags().StringP("in", "i", "", 'criu image in binary format to be decoded (stdin by default)')
	encodeCmd.Flags().StringP("out", "o", "", 'output loc of the file')

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
