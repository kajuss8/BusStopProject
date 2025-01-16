package handleFiles

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"path/filepath"

)

const gtfsUrlPath = "http://www.stops.lt/kaunas/kaunas/gtfs.zip"
const gtfsFolderPath = "C:/Users/Kajus.Sciaponis/Desktop/BusStopProject/gtfsFolder/gtfs.zip"
const destination = "C:/Users/Kajus.Sciaponis/Desktop/BusStopProject/gtfsFolder"

func DownloadGtfs()  error {
	out, err := os.Create(gtfsFolderPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(gtfsUrlPath)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	

	// Write the body to file

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func Unzip() error {
	archive, err := zip.OpenReader(gtfsFolderPath)
	if err != nil {
		return err
	}
	defer archive.Close()
	for _, file := range archive.Reader.File {
		reader, err := file.Open()
		if err != nil {
			return err
		}
		defer reader.Close()
		path := filepath.Join(destination, file.Name)
		// Remove file if it already exists; no problem if it doesn't; other cases can error out below
		_ = os.Remove(path)
		// Create a directory at path, including parents
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
		// If file is _supposed_ to be a directory, we're done
		if file.FileInfo().IsDir() {
			continue
		}
		// otherwise, remove that directory (_not_ including parents)
		err = os.Remove(path)
		if err != nil {
			return err
		}
		// and create the actual file.  This ensures that the parent directories exist!
		// An archive may have a single file with a nested path, rather than a file for each parent dir
		writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer writer.Close()
		_, err = io.Copy(writer, reader)
		if err != nil {
			return err
		}
	}
	return nil
}
