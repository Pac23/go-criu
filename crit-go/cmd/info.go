package cmd

import (
	//"fmt"
	"crit-go/gocrit"
	//"errors"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Gives info of the image",
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
		gocrit.Info(args)
	},
}

//var inloc string

func init() {
	//rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().StringVarP(&inloc, "in", "i", "", "show info about criu image (stdin by default)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
