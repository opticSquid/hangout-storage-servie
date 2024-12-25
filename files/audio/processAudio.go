package audio

import (
	"os/exec"

	"hangout.com/core/storage-service/logger"
)

func ProcessAudio(inputFile string, outputFile string, log logger.Log) error {
	log.Debug("pipeline status status", "track", "audio", "encoder", "aac", "status", "starting")
	var cmd *exec.Cmd
	var err error
	cmd = exec.Command("ffmpeg", "-y", "-i", inputFile, "-vn", "-c:a", "aac", "-b:a", "200k", "-ar", "48000", "-ac", "2", outputFile+"_audio.m4a")
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing auio workflow", "error", err.Error())
	} else {
		log.Debug("pipeline status", "track", "audio", "encoder", "aac", "status", "finished")
	}
	return nil
}
