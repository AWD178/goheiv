package goheiv

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
	"fmt"
	"bufio"
	"strings"
)


type HEIVFile struct {
	FilePath string
	OutputPath string
	Width int
	Height int
	Crop bool
}


func (FileHeiv *HEIVFile) SetFile(filepath string) *HEIVFile {

	if _, err := os.Stat(filepath); err != nil {
		panic(err)
	}

	FileHeiv.FilePath = filepath

	return FileHeiv
}


func (FileHeiv *HEIVFile) SetOutput(filepath string) *HEIVFile {
	FileHeiv.OutputPath = filepath
	return FileHeiv
}

func (FileHeiv *HEIVFile) SetWidth(width int) *HEIVFile {
	FileHeiv.Width = width
	return FileHeiv
}

func (FileHeiv *HEIVFile) SetHeight(height int) *HEIVFile {
	FileHeiv.Height = height
	return FileHeiv
}

func (FileHeiv *HEIVFile) UseCrop() *HEIVFile {
	FileHeiv.Crop = true
	return FileHeiv
}


func (FileHeiv *HEIVFile) Convert() (string, error) {
	cmdName, err := exec.LookPath("tifig")

	if err != nil {
		return "", errors.New("tifig is not installed")
	}

	var args []string

	args = append(args,"-v")
	args = append(args, "-p")

	if FileHeiv.Crop {
		args = append(args, "--crop")
	}

	if FileHeiv.Width > 0 {
		args = append(args, "--width")
		args = append(args, strconv.Itoa(FileHeiv.Width))
	}

	if FileHeiv.Height > 0 {
		args = append(args, "--height")
		args = append(args, strconv.Itoa(FileHeiv.Height))
	}

	if FileHeiv.FilePath == "" {
		return "", errors.New("File path is required")
	}


	if FileHeiv.OutputPath == "" {
		return "", errors.New("Output file path is required")
	}

	args = append(args, FileHeiv.FilePath)
	args = append(args, FileHeiv.OutputPath)



	cmd := exec.Command(cmdName, args...)

	cmdReader, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println("this 1")
		return "", err
	}

	scanner := bufio.NewScanner(cmdReader)

	var outputCmd []string

	go func() {
		for scanner.Scan() {
			outputCmd = append(outputCmd, scanner.Text())
		}
	}()

	err = cmd.Start()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return "", err

	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		return "", err
	}

	return strings.Join(outputCmd,"\n"),nil

}

