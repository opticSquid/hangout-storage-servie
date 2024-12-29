package abr

import (
	"os/exec"

	"hangout.com/core/storage-service/logger"
)

func CreatePlaylist(outputFilePath string, encoding string, log logger.Log) error {
	log.Info("pipeline status status", "segementation and playlist creation", "starting")
	videoFile := outputFilePath + "_" + encoding + "_"
	audioFile := outputFilePath + "_" + encoding + "_audio.opus"
	var cmd *exec.Cmd
	var err error
	cmd = exec.Command("ffmpeg", "-i", videoFile+"320.webm", "-i", videoFile+"640p.webm", "-i", videoFile+"1280p.webm", "-i", videoFile+"1920p.webm", "-i", audioFile, "-map", "0:v", "-c:v", "copy", "-map", "1:v", "-c:v", "copy", "-map", "2:v", "-c:v", "copy", "-map 3:v", "-c:v", "copy", "-map 4:a", "-c:a", "copy", "-keyint_min", "60", "-g", "60", "-use_timeline", "1", "-use_template", "1", outputFilePath+".mpd")
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing segmentation and playlist creation", "error", err.Error())
		return err
	} else {
		log.Info("pipeline status status", "segementation and playlist creation", "finished")
	}
	return nil
}
