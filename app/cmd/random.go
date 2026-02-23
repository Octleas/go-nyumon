/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type ApiRespArr struct {
	Message []string `json:"message"`
	Status string `json:"status"`
}

var dogg int

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "DogAPIからランダムに画像URLを取得し、URLを出力するよ。",
	Run: func(cmd *cobra.Command, args []string) {
			url := fmt.Sprintf("https://dog.ceo/api/breeds/image/random/%d",dogg)
			resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
			return
		}
		defer resp.Body.Close()
		
		message, _ := io.ReadAll(resp.Body)
		var responseObject ApiRespArr
		json.Unmarshal(message, &responseObject)
		for _, url := range responseObject.Message {
		fmt.Println(url)
		}
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
	randomCmd.Flags().IntVarP(&dogg, "images", "i", 1, "持ってくるいぬの写真の数")
}
