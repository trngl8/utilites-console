package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

const (
	RedColor   = "\x1b[31m%s\x1b[0m\n"
	BlueColor  = "\x1b[34m%s\x1b[0m\n"
	GrayColor  = "\x1b[37m%s\x1b[0m\n"
	GreenColor = "\x1b[32m%s\x1b[0m\n"
)

func main() {
	dirPath := "."
	PrintError("No .env file found")
	files, err := getFilesInDir(dirPath)

	if err != nil {
		PrintError(err.Error())
		return
	}

	printFiles(files)

	dirs, err := getDirsInDir(dirPath)

	if err != nil {
		PrintError(err.Error())
		return
	}

	printDirectories(dirs)

}

func repeatLine(line string, times int) string {
	var result = line
	for i := 0; i < times; i++ {
		result = result + line
	}
	return result
}

func PrintLine(str string, times int) {
	fmt.Printf(GrayColor, repeatLine(str, times))
}

func PrintSuccess(str string, value int) {
	var total = str + strconv.Itoa(value)
	fmt.Printf(GreenColor, total)
}

func PrintError(str string) {
	str = "Error: " + str
	fmt.Printf(RedColor, str)
}

func getFilesInDir(dirPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func getDirsInDir(dirPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func printDirectories(lines []string) {
	for _, line := range lines {
		fmt.Printf(GrayColor, line)
	}
	PrintLine("-", 20)
	PrintTotal(lines, "directories")
}

func printFiles(lines []string) {
	for _, line := range lines {
		fmt.Printf(BlueColor, line)
	}
	PrintLine("-", 20)
	PrintTotal(lines, "files")
}

func PrintTotal(lines []string, t string) {
	PrintSuccess("Total "+t+": ", len(lines))
	PrintLine("=", 20)
}

func generateValues(lines []string) []int {
	var strLengths []int
	for _, line := range lines {
		strLengths = append(strLengths, len(line))
	}
	return strLengths
}
