package h264

import (
	"os/exec"

	"hangout.com/core/storage-service/logger"
)

func ProcessCRFmode(inputFile string, outputFile string, log logger.Log) error {
	log.Debug("pipeline status status", "encoder", "h264", "method", "crf", "status", "starting")
	var cmd *exec.Cmd
	var err error
	// creating 360p, 720p, 1080p using h264 crf mode
	cmd = exec.Command("ffmpeg", "-i", inputFile,
		"-c:v", "libx264", "-crf", "23", "-maxrate", "500k", "-bufsize", "1M", "-vf", "scale=-2:360", "-an", outputFile+"_h264_360p.mp4",
		"-c:v", "libx264", "-crf", "23", "-maxrate", "1M", "-bufsize", "3M", "-vf", "scale=-2:720", "-an", outputFile+"_h264_720p.mp4",
	)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing h264 crf workflow", "error", err.Error())
	} else {
		log.Debug("pipeline status", "encoder", "h264", "method", "crf", "status", "finished")
	}
	return nil
}
