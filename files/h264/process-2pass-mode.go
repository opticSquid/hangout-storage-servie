package h264

import (
	"os"
	"os/exec"

	"hangout.com/core/storage-service/logger"
)

func Process2PassMode(inputFile string, outputFile string, log logger.Log) error {
	var cmd *exec.Cmd
	var err error
	// creating 1080p using h264 2 pass mode
	log.Debug("pipeline status", "encoder", "h264", "method", "2 pass", "status", "starting")
	// doing 1st pass
	cmd = exec.Command("ffmpeg", "-i", inputFile,
		"-c:v", "libx264", "-pass", "1", "-passlogfile", outputFile, "-b:v", "2M", "-fps_mode", "cfr", "-vf", "scale=-2:1080", "-an", "-f", "null", "/dev/null",
	)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing h264 2 pass workflow, error in 1st pass", "current resolution", "1080p", "error", err.Error())
	} else {
		log.Debug("pipeline status", "encoder", "h264", "method", "2 pass", "current pass", 1, "current resolution", "1080p", "status", "finished")
	}
	// doing 2nd pass
	// creating 1080p video in 2nd pass out of 1st pass log and mbtree files
	cmd = exec.Command("ffmpeg", "-i", inputFile,
		"-c:v", "libx264", "-pass:", "2", "-passlogfile", outputFile, "-b:v", "2M", "-fps_mode", "cfr", "-vf", "scale=-2:1080", "-an", outputFile+"_h264_1080p.mp4",
	)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing h264 2 pass workflow, error in 2nd pass", "current resolution", "1080p", "error", err.Error())
	} else {
		log.Debug("pipeline status", "encoder", "h264", "method", "2 pass", "current pass", 2, "current resolution", "1080p", "status", "finished")
	}

	// deleting ffmpeg generated log file
	err = os.Remove(outputFile + "-0.log")
	if err != nil {
		log.Error("error in deleting ffmpeg  h264 log file", "current resolution", "1080p", "error", err.Error())
		return err
	} else {
		log.Debug("deleted ffmpeg h264 log file", "current resolution", "1080p")
	}
	// deleting ffmpeg generated mbtree file
	err = os.Remove(outputFile + "-0.log.mbtree")
	if err != nil {
		log.Error("error in deleting ffmpeg h264 mbtree file", "current resolution", "1080p", "error", err.Error())
		return err
	} else {
		log.Debug("deleted ffmpeg h264 mbtree file", "current resolution", "1080p")
	}
	log.Debug("pipeline status", "encoder", "h264", "method", "2 pass", "status", "finished")
	return nil
}
