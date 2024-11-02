package files

import (
	"os"
	"os/exec"
	"strings"

	"hangout.com/core/storage-service/config"
	"hangout.com/core/storage-service/logger"
)

type video struct {
	filename string
}

func (v *video) processMedia(cfg *config.Config, log logger.Log) error {
	splittedFilename := strings.Split(v.filename, ".")
	inputFile := cfg.Hangout.Media.UploadPath + "/" + v.filename
	outputFile := cfg.Hangout.Media.ProcessedPath + "/" + splittedFilename[0]
	var err error
	err = processH264(inputFile, outputFile, log)
	if err != nil {
		log.Error("error in deleting original file", "error", err.Error())
	}
	err = processH265(inputFile, outputFile, log)
	if err != nil {
		log.Error("error in deleting original file", "error", err.Error())
	}
	// deleting original uploaded file
	err = os.Remove(inputFile)
	if err != nil {
		log.Error("error in deleting original file", "error", err.Error())
		return err
	} else {
		log.Debug("deleted original file")
	}
	return err
}

// generte h264 (crf mode and 2 pass mode) encoded versions of the uploaded file in 360p, 720p, 1080p
func processH264(inputFile string, outputFile string, log logger.Log) error {
	log.Info("h264 pipeline", "status", "starting")
	var cmd *exec.Cmd
	var err error

	log.Debug("pipeline status status", "encoder", "h264", "method", "crf", "status", "starting")
	// creating 360p, 720p, 1080p using h264 crf mode
	cmd = exec.Command("ffmpeg", "-i", inputFile,
		"-c:v", "libx264", "-crf", "23", "-b:v", "500k", "-vf", "scale=-2:360", outputFile+"_h264_360p_crf.mp4",
		"-c:v", "libx264", "-crf", "23", "-b:v", "1000k", "-vf", "scale=-2:720", outputFile+"_h264_720p_crf.mp4",
		"-c:v", "libx264", "-crf", "23", "-b:v", "2000k", "-vf", "scale=-2:1080", outputFile+"_h264_1080p_crf.mp4",
	)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing h264 crf workflow", "error", err.Error())
	} else {
		log.Debug("pipeline status", "encoder", "h264", "method", "crf", "status", "finished")
	}

	// creating 360p, 720p, 1080p using h264 2 pass mode
	log.Debug("pipeline status", "encoder", "h264", "method", "2 pass", "status", "starting")
	// doing 1st pass
	cmd = exec.Command("ffmpeg", "-i", inputFile,
		"-c:v", "libx264", "-pass", "1", "-passlogfile", outputFile, "-fps_mode", "cfr", "-f", "null", "/dev/null",
	)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing h264 2 pass workflow, error in 1st pass", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline status", "encoder", "h264", "method", "2 pass", "current pass", 1, "status", "finished")
	}
	// doing 2nd pass
	// creating 360p video in 2nd pass out of 1st pass log and mbtree files
	cmd = exec.Command("ffmpeg", "-i", inputFile,
		"-c:v", "libx264", "-pass:", "2", "-passlogfile", outputFile, "-b:v", "500k", "-fps_mode", "cfr", "-vf", "scale=-2:360", outputFile+"_h264_360p_2pass.mp4",
	)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing h264 2 pass workflow, error in 2nd pass", "current resolution", "360p", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline status", "encoder", "h264", "method", "2 pass", "current pass", 2, "current resolution", "360p", "status", "finished")
	}
	// doing 2nd pass
	// creating 720p video in 2nd pass out of 1st pass log and mbtree files
	cmd = exec.Command("ffmpeg", "-i", inputFile,
		"-c:v", "libx264", "-pass:", "2", "-passlogfile", outputFile, "-b:v", "1000k", "-fps_mode", "cfr", "-vf", "scale=-2:720", outputFile+"_h264_720p_2pass.mp4",
	)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing h264 2 pass workflow, error in 2nd pass", "current resolution", "720p", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline status", "encoder", "h264", "method", "2 pass", "current pass", 2, "current resolution", "720p", "status", "finished")
	}
	// doing 2nd pass
	// creating 1080p video in 2nd pass out of 1st pass log and mbtree files
	cmd = exec.Command("ffmpeg", "-i", inputFile,
		"-c:v", "libx264", "-pass:", "2", "-passlogfile", outputFile, "-b:v", "2000k", "-fps_mode", "cfr", "-vf", "scale=-2:1080", outputFile+"_h264_1080p_2pass.mp4",
	)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing h264 2 pass workflow, error in 2nd pass", "current resolution", "1080p", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline status", "encoder", "h264", "method", "2 pass", "current pass", 2, "current resolution", "1080p", "status", "finished")
	}
	log.Debug("pipeline status", "encoder", "h264", "method", "2 pass", "status", "finished")

	// deleting ffmpeg generated log file
	err = os.Remove(outputFile + "-0.log")
	if err != nil {
		log.Error("error in deleting ffmpeg  h264 log file", "error", err.Error())
		return err
	} else {
		log.Debug("deleted ffmpeg h264 log file")
	}

	// deleting ffmpeg generated mbtree file
	err = os.Remove(outputFile + "-0.log.mbtree")
	if err != nil {
		log.Error("error in deleting ffmpeg h264 mbtree file", "error", err.Error())
		return err
	} else {
		log.Debug("deleted ffmpeg h264 mbtree file")
	}

	log.Info("h264 pipeline", "status", "finished")
	return nil
}

// generte h265 (crf mode) encoded versions of the uploaded file in 360p, 720p, 1080p
func processH265(inputFile string, outputFile string, log logger.Log) error {
	log.Info("h265 pipeline", "status", "starting")
	var cmd *exec.Cmd
	var err error

	log.Debug("pipeline status", "encoder", "h265", "method", "crf", "status", "starting")
	// creating 360p, 720p, 1080p using h265 crf mode
	cmd = exec.Command("ffmpeg", "-i", inputFile,
		"-c:v", "libx265", "-crf", "28", "-b:v", "500k", "-vf", "scale=-2:360", outputFile+"_h265_360p_crf.mp4",
		"-c:v", "libx265", "-crf", "28", "-b:v", "1000k", "-vf", "scale=-2:720", outputFile+"_h265_720p_crf.mp4",
		"-c:v", "libx265", "-crf", "28", "-b:v", "2000k", "-vf", "scale=-2:1080", outputFile+"_h265_1080p_crf.mp4",
	)
	_, err = cmd.Output()
	if err != nil {
		log.Error("error in processing h265 crf workflow", "error", err.Error())
		return err
	} else {
		log.Debug("pipeline status", "encoder", "h265", "method", "crf", "status", "finished")
	}

	log.Info("h265 pipeline", "status", "finished")
	return nil
}
