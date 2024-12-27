package files

import (
	"os"
	"strings"

	"hangout.com/core/storage-service/config"
	"hangout.com/core/storage-service/files/abr"
	"hangout.com/core/storage-service/files/vp9"
	"hangout.com/core/storage-service/logger"
)

type video struct {
	filename string
}

func (v *video) processMedia(cfg *config.Config, log logger.Log) error {
	splittedFilename := strings.Split(v.filename, ".")
	inputFile := cfg.Path.Upload + "/" + v.filename
	outputFolder := cfg.Path.Storage + "/" + splittedFilename[0]
	filename := splittedFilename[0]
	var err error
	err = os.Mkdir(outputFolder, 0755)
	if err != nil {
		log.Error("could not create base output folder", "err", err.Error())
	}
	err = processVp9(inputFile, outputFolder, filename, log)
	if err != nil {
		log.Error("error in vp9 pipeline", "error", err.Error())
	}
	return nil
}

func processVp9(inputFilePath string, outputFolder string, filename string, log logger.Log) error {
	log.Info("pipeline checkpoint", "file", inputFilePath, "enocder", "vp9", "status", "starting processing")
	outputFilePath := outputFolder + "/" + filename
	vp9.ProcessSDRResolutions(inputFilePath, outputFilePath, log)
	vp9.ProcessAudio(inputFilePath, outputFilePath, log)
	abr.CreatePlaylist(outputFilePath, "vp9", log)
	log.Info("pipeline checkpoint", "file", inputFilePath, "enocder", "vp9", "status", "finished processing")
	return nil
}
