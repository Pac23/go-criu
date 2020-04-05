package cmd

import (
	"bufio"
	"crit-go/gocrit"
	//"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Convert Binary to Json",
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
		if inloc == "" {
			reader := bufio.NewReader(os.Stdin)
			stdininp, err := reader.ReadString('\n')
			gocrit.check(err)
			/*
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			*/
			err = gocrit.Decode(stdininp, outloc, pretty)
			/*
				if err != nil {
					fmt.Println(err)
				}
			*/
			gocrit.check(err)
		} else {
			err := gocrit.Decode(inloc, outloc, pretty)
			gocrit.check(err)
		}
	},
}

var inloc, outloc string
var pretty bool

func init() {

	//rootCmd.AddCommand(decodeCmd)
	decodeCmd.Flags().StringVarP(&inloc, "in", "i", "", "criu image in binary format to be decoded (stdin by default)")
	decodeCmd.Flags().StringVarP(&outloc, "out", "o", "", "where to put the image in json format(Stdout by default)")
	decodeCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Multiline with indents and some numerical fields in field-specific format")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

