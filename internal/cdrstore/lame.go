package cdrstore

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/fdully/calljournal/internal/logging"
)

// convert wav to mp3 using lame executable and temp files. Temp files are removed after converting.
func Wav2Mp3(ctx context.Context, wav *bytes.Buffer) (*bytes.Buffer, error) {
	logger := logging.FromContext(ctx)

	// temporary file for wav
	w, err := ioutil.TempFile(os.TempDir(), "cj-")
	if err != nil {
		return nil, fmt.Errorf("failed to create tmp file: %w", err)
	}

	defer func() {
		err := os.Remove(w.Name())
		if err != nil {
			logger.Errorf("failed to remove temp file: %v", err)
		}
	}()

	_, err = wav.WriteTo(w)
	if err != nil {
		return nil, fmt.Errorf("failed to write wav to %s file: %w", w.Name(), err)
	}

	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("failed to close file: %w", err)
	}

	// temporary file for mp3
	m, err := ioutil.TempFile(os.TempDir(), "cj-")
	if err != nil {
		return nil, fmt.Errorf("failed to create tmp file: %w", err)
	}

	defer func() {
		err := os.Remove(m.Name())
		if err != nil {
			logger.Errorf("failed to remove temp file: %v", err)
		}
	}()

	if err := m.Close(); err != nil {
		return nil, fmt.Errorf("failed to close file: %w", err)
	}

	// execute lame to convert wav to mp3
	l, err := exec.LookPath("lame")
	if err != nil {
		return nil, fmt.Errorf("didn't find 'lame' executable: %w", err)
	}

	command := exec.Command(l, "-S", "--preset", "insane", w.Name(), m.Name())
	if err := command.Run(); err != nil {
		return nil, fmt.Errorf("failed to execute lame: %w, got shell exitcode: %d",
			err, command.ProcessState.ExitCode())
	}

	// return mp3 in []byte
	mp3, err := ioutil.ReadFile(m.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to read mp3 file after conversion: %w", err)
	}

	return bytes.NewBuffer(mp3), nil
}

func WavFileToMp3File(ctx context.Context, wavFile, mp3File string) error {
	l, err := exec.LookPath("lame")
	if err != nil {
		return fmt.Errorf("didn't find 'lame' executable: %w", err)
	}

	c := exec.Command(l, "-S", "--preset", "insane", wavFile, mp3File)
	if err := c.Run(); err != nil {
		return fmt.Errorf("failed to execute lame: %w, got shell exitcode: %d",
			err, c.ProcessState.ExitCode())
	}

	return nil
}

func PingLame() error {
	l, err := exec.LookPath("lame")
	if err != nil {
		return fmt.Errorf("didn't find 'lame' executable: %w", err)
	}

	command := exec.Command(l, "--version")
	if err := command.Run(); err != nil {
		return fmt.Errorf("failed to ping lame command: %w", err)
	}

	return nil
}
