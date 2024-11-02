package files

import (
	"errors"
	"regexp"

	"hangout.com/core/storage-service/config"
	"hangout.com/core/storage-service/logger"
)

type File struct {
	ContentType string
	Filename    string
}

func (f *File) Process(cfg *config.Config, log logger.Log) error {
	isImage, _ := regexp.MatchString(`^image/`, f.ContentType)
	isVideo, _ := regexp.MatchString(`^video/`, f.ContentType)
	var media media
	if isImage {
		media = &image{filename: f.Filename}
	} else if isVideo {
		media = &video{filename: f.Filename}
	} else {
		log.Debug("unsupported content type. can not process file", "contentType", f.ContentType)
		return errors.New("unsupported file type received, contentType is: " + f.ContentType)
	}
	err := media.processMedia(cfg, log)
	if err != nil {
		return err
	}
	return nil
}
