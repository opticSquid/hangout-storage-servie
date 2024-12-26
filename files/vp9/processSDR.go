package vp9

import (
	"os/exec"

	"hangout.com/core/storage-service/logger"
)

func ProcessSDRResolutions(inputFilePath string, outputFilePath string, log logger.Log) error {
	log.Info("pipeline checkpoint", "file", inputFilePath, "enocder", "vp9", "media-type", "video-sdr", "status", "starting processing")
	var err error
	// processing 320p videos
	err = process320p(inputFilePath, outputFilePath, log)
	if err != nil {
		return err
	}
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
	log.Info("pipeline checkpoint", "file", inputFilePath, "enocder", "vp9", "media-type", "video-sdr", "status", "finished")
	return nil
}

func process320p(inputFilePath string, outputFilePath string, log logger.Log) error {
	log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "vp9", "media-type", "video-sdr", "resolution", "240x320", "status", "starting processing")
	var cmd *exec.Cmd
	var err error
	outputFilePath = outputFilePath + "_vp9_320p.webm"
	log.Debug("Input", "Input file path", inputFilePath)
	log.Debug("Input", "output file path", outputFilePath)
	cmd = exec.Command("ffmpeg", "-i", inputFilePath, "-vf", "scale=240x320", "-b:v", "150k", "-minrate", "75k", "-maxrate", "218k", "-tile-columns", "0", "-g", "240", "-threads", "2", "-quality", "good", "-crf", "37", "-c:v", "libvpx-vp9", "-an", "-pass", "1", "-speed", "4", outputFilePath)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing video", "file", inputFilePath, "encoder", "vp9", "resolution", "240x320", "pass", "1", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "vp9", "media-type", "video-sdr", "resolution", "240x320", "pass", "1", "status", "finished")
	}
	cmd = exec.Command("ffmpeg", "-y", "-i", inputFilePath, "-vf", "scale=240x320", "-b:v", "150k", "-minrate", "75k", "-maxrate", "218k", "-tile-columns", "0", "-g", "240", "-threads", "2", "-quality", "good", "-crf", "37", "-c:v", "libvpx-vp9", "-an", "-pass", "2", "-speed", "1", outputFilePath)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing video", "file", inputFilePath, "encoder", "vp9", "resolution", "240x320", "pass", "2", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "vp9", "media-type", "video-sdr", "resolution", "240x320", "pass", "2", "status", "finished")
	}
	return nil
}

func process640p(inputFilePath string, outputFilePath string, log logger.Log) error {
	log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "vp9", "media-type", "video-sdr", "resolution", "480x640", "status", "starting processing")
	var cmd *exec.Cmd
	var err error
	outputFilePath = outputFilePath + "_vp9_640p.webm"
	log.Debug("Input", "Input file path", inputFilePath)
	log.Debug("Input", "output file path", outputFilePath)
	cmd = exec.Command("ffmpeg", "-i", inputFilePath, "-vf", "scale=480x640", "-b:v", "750k", "-minrate", "375k", "-maxrate", "1088k", "-tile-columns", "1", "-g", "240", "-threads", "4", "-quality", "good", "-crf", "33", "-c:v", "libvpx-vp9", "-an", "-pass", "1", "-speed", "4", outputFilePath)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing video", "file", inputFilePath, "encoder", "vp9", "resolution", "480x640", "pass", "1", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "vp9", "media-type", "video-sdr", "resolution", "480x640", "pass", "1", "status", "finished")
	}
	cmd = exec.Command("ffmpeg", "-y", "-i", inputFilePath, "-vf", "scale=480x640", "-b:v", "750k", "-minrate", "375k", "-maxrate", "1088k", "-tile-columns", "1", "-g", "240", "-threads", "4", "-quality", "good", "-crf", "33", "-c:v", "libvpx-vp9", "-an", "-pass", "2", "-speed", "4", outputFilePath)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing video", "file", inputFilePath, "encoder", "vp9", "resolution", "480x640", "pass", "2", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "vp9", "media-type", "video-sdr", "resolution", "480x640", "pass", "2", "status", "finished")
	}
	return nil
}

func process1280p(inputFilePath string, outputFilePath string, log logger.Log) error {
	log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "vp9", "media-type", "video-sdr", "resolution", "720x1280", "status", "starting processing")
	var cmd *exec.Cmd
	var err error
	outputFilePath = outputFilePath + "_vp9_1280p.webm"
	log.Debug("Input", "Input file path", inputFilePath)
	log.Debug("Input", "output file path", outputFilePath)
	cmd = exec.Command("ffmpeg", "-i", inputFilePath, "-vf", "scale=720x1280", "-b:v", "1024k", "-minrate", "512k", "-maxrate", "1485k", "-tile-columns", "2", "-g", "240", "-threads", "8", "-quality", "good", "-crf", "32", "-c:v", "libvpx-vp9", "-an", "-pass", "1", "-speed", "4", outputFilePath)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing video", "file", inputFilePath, "encoder", "vp9", "resolution", "720x1280", "pass", "1", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "vp9", "media-type", "video-sdr", "resolution", "720x1280", "pass", "1", "status", "finished")
	}
	cmd = exec.Command("ffmpeg", "-y", "-i", inputFilePath, "-vf", "scale=720x1280", "-b:v", "1024k", "-minrate", "512k", "-maxrate", "1485k", "-tile-columns", "2", "-g", "240", "-threads", "8", "-quality", "good", "-crf", "32", "-c:v", "libvpx-vp9", "-an", "-pass", "2", "-speed", "4", outputFilePath)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing video", "file", inputFilePath, "encoder", "vp9", "resolution", "720x1280", "pass", "2", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "vp9", "media-type", "video-sdr", "resolution", "720x1280", "pass", "2", "status", "finished")
	}
	return nil
}

func process1920p(inputFilePath string, outputFilePath string, log logger.Log) error {
	log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "vp9", "media-type", "video-sdr", "resolution", "1080x1920", "status", "starting processing")
	var cmd *exec.Cmd
	var err error
	outputFilePath = outputFilePath + "_vp9_1920p.webm"
	log.Debug("Input", "Input file path", inputFilePath)
	log.Debug("Input", "output file path", outputFilePath)
	cmd = exec.Command("ffmpeg", "-i", inputFilePath, "-vf", "scale=1080x1920", "-b:v", "1800k", "-minrate", "900k", "-maxrate", "2610k", "-tile-columns", "3", "-g", "240", "-threads", "8", "-quality", "good", "-crf", "31", "-c:v", "libvpx-vp9", "-an", "-pass", "1", "-speed", "4", outputFilePath)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing video", "file", inputFilePath, "encoder", "vp9", "resolution", "1080x1920", "pass", "1", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "vp9", "media-type", "video-sdr", "resolution", "1080x1920", "pass", "1", "status", "finished")
	}
	cmd = exec.Command("ffmpeg", "-y", "-i", inputFilePath, "-vf", "scale=1080x1920", "-b:v", "1800k", "-minrate", "900k", "-maxrate", "2610k", "-tile-columns", "3", "-g", "240", "-threads", "8", "-quality", "good", "-crf", "31", "-c:v", "libvpx-vp9", "-an", "-pass", "2", "-speed", "4", outputFilePath)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing video", "file", inputFilePath, "encoder", "vp9", "resolution", "1080x1920", "pass", "2", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline checkpoint", "file", inputFilePath, "encoder", "vp9", "media-type", "video-sdr", "resolution", "1080x1920", "pass", "2", "status", "finished")
	}
	return nil
}
