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
	"os"

	"github.com/spf13/cobra"
	gdrive "github.com/suvam0451/drivekit/gdrive"
)

/**
*	Example queries -->
*	go run main.go gdrive upload -i 1Ud4MyIu5fnpJqnB2epovCP4SJqnX2i5Q -d ./uplink
 */
// gdriveCmd represents the gdrive command
var gdriveUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Use to upload files to google drive.",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get named arguments
		tokenpath, _ := cmd.Flags().GetString("tokenpath")
		credpath, _ := cmd.Flags().GetString("credentialspath")
		// jsonfile, _ := cmd.Flags().GetString("jsonquery")
		// query, _ := cmd.Flags().GetString("query")
		id, _ := cmd.Flags().GetString("id")
		uploadpath, _ := cmd.Flags().GetString("uploadpath")

		if _endpoint, err := gdrive.Authenticate(credpath, tokenpath); err == nil {
			fmt.Println("Authentication successful.")
			if fi, err := os.Stat(uploadpath); err == nil {
				switch mode := fi.Mode(); {
				case mode.IsDir():
					gdrive.UploadFolder(_endpoint, id, uploadpath)
				case mode.IsRegular():
					gdrive.UploadFile(_endpoint, id, uploadpath)
				default:
					log.Fatalf("Path provided is neither a file, nor a directory.")
				}
			} else {
				log.Fatalf("The input file/directory doesn't exist.")
			}
		} else {
			log.Println(err.Error())
		}
	},
}

func init() {
	gdriveCmd.AddCommand(gdriveUploadCmd)

	// Flags
	gdriveUploadCmd.PersistentFlags().StringP("tokenpath", "t", "./token.json", "fullpath location to the token file for this run.")
	gdriveUploadCmd.PersistentFlags().StringP("credentialspath", "c", "./credentials.json", "fullpath location to your credentials file.")
	gdriveUploadCmd.PersistentFlags().StringP("query", "q", "none", "Query argument. View the README for possible values.")
	gdriveUploadCmd.PersistentFlags().StringP("jsonquery", "j", "", "Formatted JSON file as query. Use to carry out multiple in-order operations.")
	gdriveUploadCmd.PersistentFlags().StringP("id", "i", "", "ID field for single queries. See README for usecases. Use JSON files for multiple concurrect queries.")
	gdriveUploadCmd.PersistentFlags().StringP("uploadpath", "d", "", "Location of object to upload. File for single file upload. Folder for multiple file upload")
}
