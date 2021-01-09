package downloader

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

const (
	contentLength = "Content-Length"

	contentDisposition = "Content-Disposition"
	contentSeparator   = "."

	minimalFileExtensionSize = 2
)

type Downloader struct {
	allowedExtensions map[string]struct{}
	// dir/%s:{random value}/%s:{stored filename from url}
	saveDirTemplate string
	// "attachment; filename=\"%s\""
	filenameFormat string
}

func (d *Downloader) DownloadFile(url string) (path, extension string, err error) {
	// head data first to get data's length and check for it's extension
	head, err := http.Head(url)
	if err != nil {
		return path, extension, err
	}

	var filename string

	if filename, extension, err = d.checkFileName(head.Header); err != nil {
		return path, extension, err
	}

	// TODO : add parallel download
	// starting file storing process
	randomFiledir, err := uuid.NewV4()
	if err != nil {
		return path, extension, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return path, extension, err
	}
	defer resp.Body.Close()

	path = fmt.Sprintf(d.saveDirTemplate, randomFiledir.String(), filename)
	out, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0x777)
	if err != nil {
		return path, extension, err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return
}

func (d *Downloader) checkContentLength(header http.Header) (len int64, err error) {
	return strconv.ParseInt(header[contentLength][0], 10, 64)
}

func (d *Downloader) checkFileName(header http.Header) (filename, extension string, err error) {
	if _, err = fmt.Sscanf(
		header[contentDisposition][0],
		d.filenameFormat, &filename,
	); err != nil {
		return filename, extension, err
	}

	filename = filename[:len(filename)-1]
	ext := strings.Split(filename, contentSeparator)
	if len(ext) != minimalFileExtensionSize {
		return filename, extension, errors.New("file does not have any extension")
	}
	extension = ext[len(ext)-1]

	if _, exist := d.allowedExtensions[extension]; !exist {
		return filename, extension, errors.New("extension mismatch")
	}

	return filename, extension, nil
}

func NewDownloader(
	allowedExtensions map[string]struct{},
	saveDirTemplate string,
	filenameFormat string,
) *Downloader {
	return &Downloader{
		allowedExtensions: allowedExtensions,
		saveDirTemplate:   saveDirTemplate,
		filenameFormat:    filenameFormat,
	}
}
