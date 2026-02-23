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

type ApiResp struct {
	Message string `json:"message"`
	Status string `json:"status"`
	ErrCode int `json:"code"`
}

var dogSpecies string

var breedCmd = &cobra.Command{
	Use:   "breed",
	Short: "./dog cli breed {犬種名} で犬種を指定し, いぬのURLをとってきます.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("Error: 犬種が指定されていません\n")
			os.Exit(1)
		}
		dogSpecies := args[0]
		url := fmt.Sprintf("https://dog.ceo/api/breed/%s/images/random", dogSpecies)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("BADなHTTPリクエスト:%s\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		data, _ := io.ReadAll(resp.Body)
		var responseObject ApiResp
		json.Unmarshal(data, &responseObject)
		switch responseObject.ErrCode {
		case 404:
			fmt.Printf("Error: 無効な犬種です\n")
			os.Exit(1)
		case 0:
			fmt.Println(responseObject.Message)
		}
	},
}

func init() {
	rootCmd.AddCommand(breedCmd)
}
