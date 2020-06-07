package gdrive

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/api/drive/v3"
)

// FileData : Used to offload file data into JSON
type FileData struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	Extension string `json:"ext"`
}

// Syntax memo --> List().Q("ID").Fields("nextPageToken, files(id,name)").Spaces("drive")

// ListFilesInFolder : Lists files under folder (Provide folder ID)
func ListFilesInFolder(_service *drive.Service, folderid string) ([]*drive.File, error) {
	fmt.Println("Yeet")
	var fs []*drive.File
	var data []FileData

	pageToken := ""
	for {
		q := _service.Files.List().Q("'" + folderid + "' in parents and trashed=false").Spaces("drive").Fields("nextPageToken, files(id,name)")
		if pageToken != "" {
			q = q.PageToken(pageToken)
		}
		r, err := q.Do()
		if err != nil {
			fmt.Printf("An error occured...")
			fmt.Printf(err.Error())
			return fs, err
		}
		fmt.Println("Navigated one page...")
		for _, _file := range r.Files {
			data = append(data, FileData{
				Name:      _file.Name,
				ID:        _file.Id,
				Extension: _file.FileExtension,
			})
		}

		pageToken = r.NextPageToken
		if pageToken == "" { // No more pages
			break
		}
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)

	return fs, nil
}

// DownloadFile -- downloads file
func DownloadFile(_service *drive.Service, _id string) {
	test, err1 := _service.Files.Get(_id).Do()
	if err1 == nil {
		fmt.Println(test.Name)
	}
	res := _service.Files.Get(_id)
	res2, err := res.Download()
	defer res2.Body.Close()
	out, err := os.Create(test.Name)
	if err != nil {
		// panic?
	}
	defer out.Close()
	io.Copy(out, res2.Body)
}

// UploadFile : Upload a file to cloud
func UploadFile(_service *drive.Service, name string, folderid string, filename string) (*drive.File, error) {

	if filecontent, err := os.Open(filename); err == nil {
		f := &drive.File{
			Name:    name,
			Parents: []string{folderid},
		}

		file, err := _service.Files.Create(f).Media(filecontent).Do()
		if err != nil {
			log.Println("Failed to create file:" + err.Error())
		}

		return file, nil
	}

	return nil, nil
}
