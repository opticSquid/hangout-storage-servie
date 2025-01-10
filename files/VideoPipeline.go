package files

import (
	"os"
	"strings"

	"hangout.com/core/storage-service/config"
	"hangout.com/core/storage-service/files/abr"
	"hangout.com/core/storage-service/files/h264"
	"hangout.com/core/storage-service/files/vp9"
	"hangout.com/core/storage-service/logger"
)

type video struct {
	filename string
}

func (v *video) processMedia(cfg *config.Config, log logger.Log) error {
	splittedFilename := strings.Split(v.filename, ".")
	inputFile := cfg.Minio.UploadBucket + "/" + v.filename
	outputFolder := cfg.Minio.StorageBucket + "/" + splittedFilename[0]
	filename := splittedFilename[0]
	var err error
	err = os.Mkdir(outputFolder, 0755)
	if err != nil {
		log.Error("could not create base output folder", "err", err.Error())
	}
	err = processH264(inputFile, outputFolder, filename, log)
	if err != nil {
		log.Error("error in video processing pipeline", "error", err.Error())
	}
	return nil
}

func processVp9(inputFilePath string, outputFolder string, filename string, log logger.Log) error {
	log.Info("pipeline checkpoint", "file", inputFilePath, "enocder", "vp9", "status", "starting processing")
	outputFilePath := outputFolder + "/" + filename
	log.Debug("Input", "Input file path", inputFilePath)
	log.Debug("Input", "output file path", outputFilePath)
	vp9.ProcessSDRResolutions(inputFilePath, outputFilePath, log)
	vp9.ProcessAudio(inputFilePath, outputFilePath, log)
	abr.CreatePlaylist(outputFilePath, "vp9", log)
	log.Info("pipeline checkpoint", "file", inputFilePath, "enocder", "vp9", "status", "finished processing")
	return nil
}
func processH264(inputFilePath string, outputFolder string, filename string, log logger.Log) error {
	log.Info("pipeline checkpoint", "file", inputFilePath, "enocder", "h264", "status", "starting processing")
	outputFilePath := outputFolder + "/" + filename
	log.Debug("Check", "Input file path", inputFilePath)
	log.Debug("Check", "Output file path", outputFilePath)
	h264.ProcessSDRResolutions(inputFilePath, outputFilePath, log)
	h264.ProcessAudio(inputFilePath, outputFilePath, log)
	abr.CreatePlaylist(outputFilePath, "h264", log)
	log.Info("pipeline checkpoint", "file", inputFilePath, "enocder", "h264", "status", "finished processing")
	return nil
}
