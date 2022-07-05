package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	directoryName string
	// editorOrExecutablePath string
	removeAfter string
)

func init() {
	flag.StringVar(&directoryName, "d", "", "Directory to create and enter for the given amount of time")
	flag.StringVar(&removeAfter, "after", "1 min", "Specification for how long to keep the directory around for")
	// flag.StringVar(&editorOrExecutablePath, "e", "code", "Editor or Executable to open the directory in/with")
}

func removeDirectory() {
	err := os.RemoveAll(directoryName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to remove the directory: %v", err)
		os.Exit(69)
	}
}

func main() {
	flag.Parse()

	if removeAfter == "" {
		fmt.Fprintf(os.Stderr, "Time spec must be provided")
		os.Exit(69)
	}

	duration, err := time.ParseDuration(removeAfter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "VALID time spec must be provided")
		os.Exit(69)
	}

	if directoryName == "" {
		fmt.Fprintf(os.Stderr, "directory must be provided")
		os.Exit(69)
	}

	fi, err := os.Stat(directoryName)
	if os.IsNotExist(err) {
		err = os.Mkdir(directoryName, os.ModeAppend)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create new directory: %v", err)
			os.Exit(69)
		}
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to access directory: %s got %v", directoryName, err)
		os.Exit(69)
	}

	fi, err = os.Stat(directoryName)
	if !fi.IsDir() {
		fmt.Fprintf(os.Stderr, "Cannot remove file, must be a directory")
		os.Exit(69)
	}

	t := time.After(duration)
	<-t
	removeDirectory()
	fmt.Println("Removed directory successfully")
	os.Exit(1)
}
