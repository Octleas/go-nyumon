/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"

	"go-nyumon/app/img"

	"github.com/spf13/cobra"
)

// dogCmd represents the dog command
var dogCliCmd = &cobra.Command{
	Use:   "dog",
	Short: "イヌと出てくる",
	Long: `イヌと表示されるってわけ`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := img.Open("dog.txt")
		if err != nil {
			fmt.Println(err)
			return 
		}
		defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)

	fmt.Print(buf.String())
		},
}

func init() {
	rootCmd.AddCommand(dogCliCmd)
}
