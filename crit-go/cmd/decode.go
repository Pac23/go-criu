package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Convert Binary to Json",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) error {
    	if len(args) < 1 {
      		return errors.New("requires a -i argument with image")
    	}
		Decode(args []string)
	}
}


func init() {
	rootCmd.AddCommand(decodeCmd)
	decodeCmd.Flags().String("in", "i", "", 'criu image in binary format to be decoded (stdin by default)')
	decodeCmd.Flags().Bool("pretty", "p", false, 'Multiline with indents and some numerical fields in field-specific format')



	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

