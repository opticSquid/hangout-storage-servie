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
	cmd = exec.Command("MP4Box", "-dash", "4000", "-frag", "4000", "-rap", "-segment-name", "segment_$RepresentationID$_", "-fps", "30", videoFile+"320p.webm#video:id=320p", videoFile+"640p.webm#video:id=640p", videoFile+"1280p.webm#video:id=1280p", videoFile+"1920p.webm#video:id=1920p", audioFile+"#audio:id=English:role=main", "-out", "index.mpd")
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing segmentation and playlist creation", "error", err.Error())
		return err
	} else {
		log.Info("pipeline status status", "segementation and playlist creation", "finished")
	}
	return nil
}
