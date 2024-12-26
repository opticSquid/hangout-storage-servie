package vp9

import (
	"os/exec"

	"hangout.com/core/storage-service/logger"
)

func ProcessAudio(inputFile string, outputFile string, log logger.Log) error {
	var cmd *exec.Cmd
	var err error
	log.Info("pipeline checkpoint", "file", inputFile, "enocder", "libopus", "media-type", "audio", "status", "starting processing")
	cmd = exec.Command("ffmpeg", "-i", inputFile, "-vn", "-c:a", "libopus", outputFile+"_vp9_audio.m4a")
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing audio", "file", inputFile, "encoder", "libopus", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline checkpoint", "file", inputFile, "encoder", "libopus", "media-type", "audio", "status", "finished")
	}
	return nil
}
