package abr

import (
	"os/exec"

	"hangout.com/core/storage-service/logger"
)

func CreatePlaylist(outputFilePath string, encoding string, log logger.Log) error {
	log.Info("pipeline status status", "segementation and playlist creation", "starting")
	videoFile := outputFilePath + "_" + encoding + "_"
	audioFile := outputFilePath + "_" + encoding + "_audio.mp4"
	var cmd *exec.Cmd
	var err error
	log.Debug("Input", "video file path", videoFile)
	log.Debug("Input", "audio file path", audioFile)
	log.Debug("Input", "output file path", outputFilePath)
	cmd = exec.Command("MP4Box", "-dash", "2000", "-frag", "2000", "-rap", "-segment-name", "segment_$RepresentationID$_", "-fps", "30", videoFile+"640p.mp4#video:id=640p", videoFile+"1280p.mp4#video:id=1280p", videoFile+"1920p.mp4#video:id=1920p", audioFile+"#audio:id=English:role=main", "-out", outputFilePath+".mpd")
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing segmentation and playlist creation", "error", err.Error())
		return err
	} else {
		log.Info("pipeline status status", "segementation and playlist creation", "finished")
	}
	return nil
}
