package ffmpeg

import (
	"fmt"
	"github.com/pkg/errors"
	"os/exec"
)

func (f *Ffmpeg) mergeAndDownload(vodUrl, name, sockFileName string) error {
	args := []string{"-i", fmt.Sprintf(`%s`, vodUrl)}
	if len(f.ffmpegParams) != 0 {
		args = append(args, f.ffmpegParams...)
	}
	args = append(args, "-progress", fmt.Sprintf(`tcp://%s`, sockFileName))
	args = append(args, name, "-y")
	cmd := exec.Command(f.ffmpegExec, args...)
	err := cmd.Run()
	if err != nil {
		return errors.Wrap(err, "exec.Run")
	}

	return nil
}

func (f *Ffmpeg) checkFfmpeg() (err error) {
	cmd := exec.Command(f.ffmpegExec, "-version")
	err = cmd.Run()
	if err != nil {
		return errors.Wrap(err, "exec.Run")
	}
	return nil
}
