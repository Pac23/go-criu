package cmd

import (
	"bufio"
	"crit-go/gocrit"
	//"errors"
	//"fmt"

	"github.com/spf13/cobra"
	"os"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Convert Binary to Json",
	Long: `Converts the Binary file to Json and writes to given location
	pretty printing of json is available aswell using the pretty flag for
	more info visit https://www.criu.org/CRIT#Functionality `,
	Run: func(cmd *cobra.Command, args []string) {
		if inloc == "" {
			reader := bufio.NewReader(os.Stdin)
			stdininp, err := reader.ReadString('\n')
			gocrit.Check(err)
			gocrit.Decode(stdininp, outloc, pretty)
		} else {
			gocrit.Decode(inloc, outloc, pretty)
		}
	},
}

func init() {
	decodeCmd.Flags().StringVarP(&inloc, "in", "i", "", "criu image in binary format to be decoded (stdin by default)")
	decodeCmd.Flags().StringVarP(&outloc, "out", "o", "", "where to put the image in json format(Stdout by default)")
	decodeCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Multiline with indents and some numerical fields in field-specific format")
}
