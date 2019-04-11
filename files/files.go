package files

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
)

type osDeleteParams struct {
	fullPath   string
	deleteTree bool
}

var oneGB = int64(1024 * 1024 * 1024)
var fourGB = int64(4 * 1024 * 1024 * 1024)
var defaultBufSize = int64(64 * 1024)

func OSPathExists(fullPath string) bool {
	if _, errMsg := os.Stat(fullPath); os.IsNotExist(errMsg) {
		return false
	}
	return true
}
func IsError(errObj error) bool {
	return (errObj != nil)
}

func TempFileName(fileSuffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return hex.EncodeToString(randBytes) + fileSuffix
}

func OSDelete(osDelSettings osDeleteParams) (string, bool) {
	if errMsg := os.Remove(osDelSettings.fullPath); IsError(errMsg) {
		return errMsg.Error(), false
	}
	return "", true
}
func OSRename(fileFrom string, fileTo string) (string, bool) {
	if OSPathExists(fileFrom) {
		if errMsg := os.Rename(fileFrom, fileTo); IsError(errMsg) {
			return errMsg.Error(), false
		}
		return "", true
	}
	return "File Not Found", false
}
func foundFile(path string, fileInfo os.FileInfo, errMsg error) error {
	if fileInfo.IsDir() {
		return nil
	}
	fmt.Printf("Found %s\n", path)
	return nil
}
func FindFiles(startDir string) {
	filepath.Walk(startDir, foundFile)
}
func ListDir(dirName string) ([]os.FileInfo, error) {
	fileList, errMsg := ioutil.ReadDir(dirName)
	if errMsg != nil {
		fmt.Println(errMsg)
		return nil, errMsg
	}
	return fileList, nil
}
func FileSize(filePath string) int64 {
	var fileInBytes int64
	fileHandle, errMsg := os.Open(filePath)
	if errMsg != nil {
		fileInBytes = -1
	} else {
		defer fileHandle.Close()

		fileStat, errMsg := fileHandle.Stat()
		if errMsg != nil {
			fileInBytes = -1
		} else {
			fileInBytes = fileStat.Size()
		}
	}
	return fileInBytes
}
func CopyFileMultiWrite(fileName, srcDir, dstDir string) bool {
	var numWorkers int
	var blockSize int64

	fileSize := FileSize(path.Join(srcDir, fileName))
	if fileSize > 0 {
		switch {
		case fileSize > fourGB:
			blockSize = 4096
			numWorkers = int(fileSize/(256*1024*1024)) + 1

		case fileSize > oneGB:
			blockSize = 2048
			numWorkers = 1
		}
		fmt.Println("Size", fileSize, blockSize, numWorkers)
		if !OSPathExists(path.Join(dstDir, fileName)) {
			dstFile, errMsg := os.Create(path.Join(dstDir, fileName))
			if errMsg != nil {
				return false
			}
			if errMsg := dstFile.Truncate(fileSize); errMsg != nil {
				return false
			}
			dstFile.Sync()
			dstFile.Close()
		}
		for indivWorker := 0; indivWorker < numWorkers; indivWorker++ {
			fmt.Println("Worker", indivWorker)
			ReadWriteFile(path.Join(srcDir, fileName), path.Join(dstDir, fileName), int64(indivWorker)*blockSize, blockSize)
		}
	}
	return false
}
func ReadWriteFile(fromPath string, dstPath string, startPoint int64, bCount int64) bool {
	srcFile, errMsg := os.Open(fromPath)
	if errMsg != nil {
		fmt.Println("Failed To Open File", fromPath, errMsg)
		return false
	}
	defer srcFile.Close()

	dstFile, errMsg := os.OpenFile(dstPath, os.O_WRONLY, 0755)
	if errMsg != nil {
		fmt.Println("Failed To Open File", dstPath, errMsg)
		return false
	}
	defer dstFile.Close()

	readBuffer := make([]byte, defaultBufSize)
	newPos, errMsg := srcFile.Seek(startPoint, 0)
	if errMsg != nil {
		fmt.Println("Failed To Seek srcFile")
		return false
	}
	fmt.Println("Moved Seek To ", newPos)

	newPos, errMsg = dstFile.Seek(startPoint, 0)
	if errMsg != nil {
		fmt.Println("Failed To Seek In dstFile")
		return false
	}
	fmt.Println("Reading Until", startPoint)
	var loopCount int64
	for loopCount = 0; loopCount < bCount+1; loopCount++ {
		fmt.Println(loopCount)
		bytesRead, errMsg := srcFile.Read(readBuffer)
		if errMsg != nil {
			fmt.Println("Encountered Error Reading From Source", errMsg)
			return false
		}
		bytesWritten, errMsg := dstFile.Write(readBuffer)
		if errMsg != nil {
			return false
		}
		if bytesRead != bytesWritten {
			fmt.Println("Bytes Written Don't Equal Bytes Read")
		}
	}
	return true

}
func CopyDir(srcDir, dstDir string) {
	fileList, _ := ListDir(srcDir)
	for _, indivfile := range fileList {
		fmt.Println(indivfile.Name())
		CopyFileMultiWrite(indivfile.Name(), srcDir, dstDir)
	}
}
