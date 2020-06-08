package gdrive

// go get -u golang.org/x/oauth2/google
// go get -u google.golang.org/api/drive/v3

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

/**
*	CREDENTIAL_FILE_LOCATION : location of credentials.json
 */

// Authenticate : Authenticates you and returns handle to client
func Authenticate(credpath, tokenpath string) (*drive.Service, error) {
	if b, err := ioutil.ReadFile(credpath); err == nil {
		// Proper scope goes here
		if config, err := google.ConfigFromJSON(b, drive.DriveScope); err == nil {
			client := getClient(config, tokenpath)

			if srv, err := drive.New(client); err == nil {
				return srv, nil
				// ListFilesInFolder(srv, "1qV-5YmODxtDVJGhXFq3RTC-E2NkLmL0b")
				// DownloadFile(srv, "11YQA_nVRYvOby3Xuoi30cflhR4zTtOAV")
				// UploadFile(srv, "sugoi.json", "1qV-5YmODxtDVJGhXFq3RTC-E2NkLmL0b", "test.json")
			}
			log.Fatalf("Unable to create drive client.")
		} else {
			log.Fatalf("Unable to parse client secret file.")
		}
	} else {
		log.Fatalf("Unable to locate credentials.json file.")
	}
	return nil, nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config, tokenpath string) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := tokenpath
	tok, err := tokenFromFile(tokFile) // modified to accomodate multiple drive sources
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}
