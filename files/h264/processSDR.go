package h264

import (
	"os/exec"

	"hangout.com/core/storage-service/logger"
)

func ProcessSDRResolutions(inputFilePath string, outputFilePath string, log logger.Log) error {
	log.Info("pipeline checkpoint", "file", inputFilePath, "enocder", "h264", "media-type", "video-sdr", "status", "starting processing")
	var err error
	err = process640p(inputFilePath, outputFilePath, log)
	if err != nil {
		return err
	}
	err = process1280p(inputFilePath, outputFilePath, log)
	if err != nil {
		return err
	}
	err = process1920p(inputFilePath, outputFilePath, log)
	if err != nil {
		return err
	}
	log.Info("pipeline checkpoint", "file", inputFilePath, "enocder", "h264", "media-type", "video-sdr", "status", "finished")
	return nil
}
func process640p(inputFilePath string, outputFilePath string, log logger.Log) error {
	log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "h264", "media-type", "video-sdr", "resolution", "360x640", "status", "starting processing")
	var cmd *exec.Cmd
	var err error
	outputFilePath = outputFilePath + "_h264_640p.mp4"
	cmd = exec.Command("ffmpeg", "-i", inputFilePath, "-c:v", "libx264", "-crf", "25", "-g", "30", "-vf", "scale=320x640", "-preset", "slow", "-an", outputFilePath)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing video", "file", inputFilePath, "encoder", "h264", "resolution", "360x640", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "h264", "media-type", "video-sdr", "resolution", "360x640", "status", "finished")
	}
	return nil
}

func process1280p(inputFilePath string, outputFilePath string, log logger.Log) error {
	log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "h264", "media-type", "video-sdr", "resolution", "720x1280", "status", "starting processing")
	var cmd *exec.Cmd
	var err error
	outputFilePath = outputFilePath + "_h264_1280p.mp4"
	cmd = exec.Command("ffmpeg", "-i", inputFilePath, "-c:v", "libx264", "-crf", "25", "-g", "30", "-vf", "scale=720x1280", "-preset", "slow", "-an", outputFilePath)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing video", "file", inputFilePath, "encoder", "h264", "resolution", "720x1280", "pass", "1", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "h264", "media-type", "video-sdr", "resolution", "720x1280", "pass", "1", "status", "finished")
	}
	return nil
}

func process1920p(inputFilePath string, outputFilePath string, log logger.Log) error {
	log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "h264", "media-type", "video-sdr", "resolution", "1080x1920", "status", "starting processing")
	var cmd *exec.Cmd
	var err error
	outputFilePath = outputFilePath + "_h264_1920p.mp4"
	cmd = exec.Command("ffmpeg", "-i", inputFilePath, "-c:v", "libx264", "-crf", "25", "-g", "30", "-vf", "scale=1080x1920", "-preset", "slow", "-an", outputFilePath)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing video", "file", inputFilePath, "encoder", "h264", "resolution", "1080x1920", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "h264", "media-type", "video-sdr", "resolution", "1080x1920", "status", "finished")
	}
	return nil
}
