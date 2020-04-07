package cmd

import (
	//"fmt"
	"bufio"
	"crit-go/gocrit"
	"os"
	//"errors"
	"github.com/spf13/cobra"
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Convert Json To Binary",
	Long: `Converts the Json file to Binary and writes to given location for
	more info visit https://www.criu.org/CRIT#Functionality`,
	Run: func(cmd *cobra.Command, args []string) {
		if inloc == "" {
			reader := bufio.NewReader(os.Stdin)
			stdininp, err := reader.ReadString('\n')
			gocrit.Check(err)
			gocrit.Encode(stdininp, outloc)
		} else {
			gocrit.Encode(inloc, outloc)
		}
	},
}

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
