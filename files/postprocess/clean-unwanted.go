package postprocess

import (
	"os"
	"strings"

	"hangout.com/core/storage-service/logger"
)

func CleanUp(workerId int, encoding string, filename string, log logger.Log) {

	// delete the original file from temp directory
	storageDir := "/tmp"
	sourceFilepath := storageDir + "/" + filename
	err := os.Remove(sourceFilepath)
	if err != nil {
		log.Debug("could not delete the original file", "error", err, "path", sourceFilepath, "worker-id", workerId)
	}
	log.Debug("removed source file", "path", sourceFilepath, "worker-id", workerId)
	// remvoe transcoded files
	baseFilename := strings.Split(filename, ".")[0]
	transcodedVideoFileBaseName := storageDir + "/" + baseFilename + "/" + baseFilename + "_" + encoding + "_"
	resolutions := []string{"640p", "1280p", "1920p", "audio"}
	for _, res := range resolutions {
		finalFileName := transcodedVideoFileBaseName + res + ".mp4"
		log.Debug("removing transcoded video files", "file", finalFileName, "worker-id", workerId)
		err = os.Remove(finalFileName)
		if err != nil {
			log.Error("could not remove trnascoded video file", "error", err, "file", finalFileName, "worker-id", workerId)
		}
	}
}
