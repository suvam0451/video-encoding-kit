package gdrive

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"

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

		for _, _file := range r.Files {
			data = append(data, FileData{
				Name:      _file.Name,
				ID:        _file.Id,
				Extension: _file.FileExtension,
			})

			fs = append(fs, _file)
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

// UploadFolder -- upload a single file
func UploadFolder(_service *drive.Service, folderid string, folderdir string) {
	var files []string
	guard := make(chan struct{}, 4) // Allows 3 workers at a time
	var wg sync.WaitGroup

	if err := filepath.Walk(folderdir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	}); err == nil {
		// Channel for progress bar
		done := make(chan bool, len(files))

		for _, file := range files {
			wg.Add(1)
			go func(filename string) {
				guard <- struct{}{}
				UploadFile(_service, folderid, filename)
				done <- true // send success
				wg.Done()
				<-guard
			}(file)

		}
		counter := 0
		for i := range done {
			counter++
			if i {
				fmt.Println("\r" + fmt.Sprintf("Uploaded %d/%d files.", counter, len(files)))
			}
			if counter == len(files) {
				close(done)
			}
		}

		wg.Wait()
		fmt.Println("Done...")
	}
}

// DownloadFolder : Be skeptical
func DownloadFolder(_service *drive.Service, _id, outputpath string) {
	guard := make(chan struct{}, 4) // Allows 3 workers at a time
	var wg sync.WaitGroup

	if arr, err := ListFilesInFolder(_service, _id); err == nil {
		fmt.Println(fmt.Sprintf("Downloading %d files...", len(arr)))

		// Channel for progress bar
		done := make(chan bool, len(arr))
		// Worker loop
		for _, file := range arr {
			wg.Add(1)
			go func(fileid, filename string) {
				guard <- struct{}{}
				downloadFile(_service, fileid, filename)
				done <- true // send success
				wg.Done()
				<-guard
			}(file.Id, path.Join(outputpath, file.Name))
		}

		counter := 0
		for i := range done {
			counter++
			if i {
				fmt.Println("\r" + fmt.Sprintf("Downloaded %d/%d files.", counter, len(arr)))
			}
			if counter == len(arr) {
				close(done)
			}
		}
	}

	wg.Wait()
	fmt.Println("Done...")
}

func downloadFile(_service *drive.Service, _id, outpath string) {
	// test, err1 := _service.Files.Get(_id).Do()
	// if err1 == nil {
	// 	fmt.Println(test.Name)
	// }
	res := _service.Files.Get(_id)
	res2, err := res.Download()
	defer res2.Body.Close()
	out, err := os.Create(outpath)
	if err != nil {
		// panic?
	}
	defer out.Close()
	io.Copy(out, res2.Body)
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
// 	name: name as which to upload
// 	folderid : id of folder to send to
// 	filename : location of file
func UploadFile(_service *drive.Service, folderid string, filename string) (*drive.File, error) {

	if filecontent, err := os.Open(filename); err == nil {
		f := &drive.File{
			Name:    filepath.Base(filename),
			Parents: []string{folderid},
		}

		file, err := _service.Files.Create(f).Media(filecontent).Do()
		if err != nil {
			log.Println("Failed to create file:" + err.Error())
		}

		return file, nil
	} else {
		log.Fatalln("Input file not found. Verify your -d parameter. Is it pointing to an existing file ?")
	}

	return nil, nil
}
