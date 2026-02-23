/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"net/http"

	"github.com/spf13/cobra"
)

// type ApiRespArr struct {
// 	Message []string `json:"message"`
// 	Status string `json:"status"`
// }

var breedListCmd = &cobra.Command{
	Use:   "breed-list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://dog.ceo/api/breeds/list")
		if err != nil {
			fmt.Printf("BADなHTTPリクエスト:%s\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		data, _ := io.ReadAll(resp.Body)
		var responseObject ApiRespArr
		json.Unmarshal(data, &responseObject)
		for _, specie := range responseObject.Message {
		fmt.Println(specie)
		}
	},
}
func init() {
	rootCmd.AddCommand(breedListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// breedListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// breedListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
