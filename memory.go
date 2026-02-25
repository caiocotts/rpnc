package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var ErrUnableToSaveStack = errors.New("error: unable to save stack")
var ErrUnableToLoadStack = errors.New("error: unable to load stack")

const stackMemoryFileName = "stack.rpnc"

func pathToMemoryDir() (string, error) {
	var memoryDirName string
	if runtime.GOOS == "windows" {
		memoryDirName = "rpnc"
	} else {
		memoryDirName = ".rpnc"
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		errors.Join()
		return "", errors.Join(ErrUnableToSaveStack, err)
	}
	return filepath.Join(homeDir, memoryDirName), nil
}

func SaveStack(c Calculator) error {
	p, err := pathToMemoryDir()
	if err != nil {
		return err
	}
	err = os.MkdirAll(p, 0755)
	if err != nil {
		return errors.Join(ErrUnableToSaveStack, err)
	}
	f, err := os.Create(filepath.Join(p, stackMemoryFileName))
	if err != nil {
		return errors.Join(ErrUnableToSaveStack, err)
	}
	defer f.Close()

	for _, element := range c.Stack.ToSlice() {
		fmt.Fprintln(f, element)
	}
	return nil
}

func LoadStack(c *Calculator) error {
	p, err := pathToMemoryDir()
	if err != nil {
		return err
	}
	f, err := os.Open(filepath.Join(p, stackMemoryFileName))
	if err != nil {
		return errors.Join(ErrUnableToLoadStack, err)
	}
	defer f.Close()

	scn := bufio.NewScanner(f)
	for scn.Scan() {
		c.Stack.Push(scn.Text())
	}
	return nil
}
