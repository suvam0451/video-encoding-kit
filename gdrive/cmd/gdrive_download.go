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
*	go run main.go gdrive download -i 1Ud4MyIu5fnpJqnB2epovCP4SJqnX2i5Q -d ./downlink
 */
// gdriveCmd represents the gdrive command
var gdriveDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Use to download files from google drive.",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get named arguments
		tokenpath, _ := cmd.Flags().GetString("tokenpath")
		credpath, _ := cmd.Flags().GetString("credentialspath")
		id, _ := cmd.Flags().GetString("id")
		_path, _ := cmd.Flags().GetString("uploadpath")

		if _endpoint, err := gdrive.Authenticate(credpath, tokenpath); err == nil {
			fmt.Println("Authentication successful.")
			if fi, err := os.Stat(_path); err == nil {
				switch mode := fi.Mode(); {
				case mode.IsDir():
					fmt.Println("Folder detected")
					gdrive.DownloadFolder(_endpoint, id, _path)
				case mode.IsRegular():
					gdrive.DownloadFile(_endpoint, id)
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
	gdriveCmd.AddCommand(gdriveDownloadCmd)

	// Flags
	gdriveDownloadCmd.PersistentFlags().StringP("tokenpath", "t", "./token.json", "fullpath location to the token file for this run.")
	gdriveDownloadCmd.PersistentFlags().StringP("credentialspath", "c", "./credentials.json", "fullpath location to your credentials file.")
	gdriveDownloadCmd.PersistentFlags().StringP("query", "q", "none", "Query argument. View the README for possible values.")
	gdriveDownloadCmd.PersistentFlags().StringP("jsonquery", "j", "", "Formatted JSON file as query. Use to carry out multiple in-order operations.")
	gdriveDownloadCmd.PersistentFlags().StringP("id", "i", "", "ID field for single queries. See README for usecases. Use JSON files for multiple concurrect queries.")
	gdriveDownloadCmd.PersistentFlags().StringP("uploadpath", "d", "", "Location of object to upload. File for single file upload. Folder for multiple file upload")

	// go run main.go gdrive -q uploadfile -i 1Ud4MyIu5fnpJqnB2epovCP4SJqnX2i5Q -d ./g-taste-3.mp4
	// go run main.go gdrive -q uploadfolder -i 1Ud4MyIu5fnpJqnB2epovCP4SJqnX2i5Q -d ./uplink
	// go run main.go gdrive -q uploadfolder -i 1Ud4MyIu5fnpJqnB2epovCP4SJqnX2i5Q -d ./uplink
}
