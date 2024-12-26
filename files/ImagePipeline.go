package files

import (
	"os/exec"
	"strings"

	"hangout.com/core/storage-service/config"
	"hangout.com/core/storage-service/logger"
)

type image struct {
	filename string
}

func (i *image) processMedia(cfg *config.Config, log logger.Log) error {
	splittedFilename := strings.Split(i.filename, ".")
	inputFile := cfg.Path.Upload + "/" + i.filename
	outputFile := cfg.Path.Storage + "/" + splittedFilename[0] + "/" + splittedFilename[0]
	err := convertToAvif(inputFile, outputFile, log)
	return err

}

func convertToAvif(inputFile string, outputFile string, log logger.Log) error {
	log.Info("starting to convert Image to AVIF")
	log.Debug("resizing image to 144px width")
	cmd := exec.Command("convert", inputFile, "-resize", "144", outputFile+"_144.avif")
	_, err := cmd.Output()
	if err != nil {
		log.Error("error in converting image to avif")
	}
	log.Debug("resizing image to 360px width")
	cmd = exec.Command("convert", inputFile, "-resize", "360", outputFile+"_360.avif")
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in converting image to avif")
	}
	log.Debug("resizing image to 720px width")
	cmd = exec.Command("convert", inputFile, "-resize", "720", outputFile+"_720.avif")
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in converting image to avif")
	}
	log.Debug("resizing image to 1080px width")
	cmd = exec.Command("convert", inputFile, "-resize", "1080", outputFile+"_1080.avif")
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in converting image to avif")
	}
	return err
}
