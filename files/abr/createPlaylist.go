package abr

import (
	"os/exec"

	"hangout.com/core/storage-service/logger"
)

func CreatePlaylist(outputFolder string, outputFile string, encoding string, log logger.Log) error {
	log.Debug("pipeline status status", "segementation and playlist creation", "starting")
	videoFile := outputFile + "_" + encoding + "_"
	audioFile := outputFile + "_audio.m4a"
	var cmd *exec.Cmd
	var err error
	cmd = exec.Command("MP4Box", "-dash", "4000", "-frag", "4000", "-rap", "-segment-name", "'segment_$RepresentationID$_'", videoFile+"360p.mp4#video:id=360p", videoFile+"720p.mp4#video:id=720p", videoFile+"1080p.mp4#video:id=1080p", audioFile+"#audio:id=English:role=main", "-out", outputFolder+"/index.mpd")
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing segmentation and playlist creation", "error", err.Error())
	} else {
		log.Debug("pipeline status status", "segementation and playlist creation", "finished")
	}
	return nil
}
