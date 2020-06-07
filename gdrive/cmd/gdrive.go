/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	gdrive "github.com/suvam0451/drivekit/gdrive"
)

// gdriveCmd represents the gdrive command
var gdriveCmd = &cobra.Command{
	Use:   "gdrive",
	Short: "Use to access google drive.",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get named arguments
		tokenpath, _ := cmd.Flags().GetString("tokenpath")
		credpath, _ := cmd.Flags().GetString("credentialspath")
		jsonfile, _ := cmd.Flags().GetString("jsonquery")
		query, _ := cmd.Flags().GetString("query")
		id, _ := cmd.Flags().GetString("id")

		if _endpoint, err := gdrive.Authenticate(credpath, tokenpath); err == nil {
			fmt.Println("Authentication successful.")
			if jsonfile == "" {
				switch query {
				case "getfile":
					gdrive.DownloadFile(_endpoint, id)

				case "listdir":
					gdrive.ListFilesInFolder(_endpoint, id)
				default:
					// Do nothing
				}
			} else {

			}
		} else {
			log.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(gdriveCmd)

	// Flags
	gdriveCmd.PersistentFlags().StringP("tokenpath", "t", "./token.json", "fullpath location to the token file for this run.")
	gdriveCmd.PersistentFlags().StringP("credentialspath", "c", "./credentials.json", "fullpath location to your credentials file.")
	gdriveCmd.PersistentFlags().StringP("query", "q", "none", "Query argument. View the README for possible values.")
	gdriveCmd.PersistentFlags().StringP("jsonquery", "j", "", "Formatted JSON file as query. Use to carry out multiple in-order operations.")
	gdriveCmd.PersistentFlags().StringP("id", "i", "", "ID field for single queries. See README for usecases. Use JSON files for multiple concurrect queries.")
}
