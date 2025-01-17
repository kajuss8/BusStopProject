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

	resp, err := http.Get(gtfsUrlPath)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
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

		_ = os.Remove(path)

		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}

		if file.FileInfo().IsDir() {
			continue
		}

		err = os.Remove(path)
		if err != nil {
			return err
		}

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
