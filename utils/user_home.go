package utils

import (
	"os"
	"path/filepath"
	"runtime"
)
/**
 * @brief: Get the home directory of the current user
 */
func HomeDir() string {

	if runtime.GOOS == "windows" {
		home := os.Getenv("HOME")
		homeDriveHomePath := ""
		if homeDrive, homePath := os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"); len(homeDrive) > 0 && len(homePath) > 0 {
			homeDriveHomePath = homeDrive + homePath
		}
		userProfile := os.Getenv("USERPROFILE")

		for _, p := range []string{home, homeDriveHomePath, userProfile} {
			if len(p) == 0 {
				continue
			}
			if _, err := os.Stat(filepath.Join(p, ".spmanager", "spm.conf")); err != nil {
				continue
			}
			return p
		}

		firstSetPath := ""
		firstExistingPath := ""

		for _, p := range []string{home, userProfile, homeDriveHomePath} {
			if len(p) == 0 {
				continue
			}
			if len(firstSetPath) == 0 {
				// remember the first path that is set
				firstSetPath = p
			}
			info, err := os.Stat(p)
			if err != nil {
				continue
			}
			if len(firstExistingPath) == 0 {
				// remember the first path that exists
				firstExistingPath = p
			}
			if info.IsDir() && info.Mode().Perm()&(1<<(uint(7))) != 0 {
				// return first path that is writeable
				return p
			}
		}

		// If none are writeable, return first location that exists
		if len(firstExistingPath) > 0 {
			return firstExistingPath
		}

		// If none exist, return first location that is set
		if len(firstSetPath) > 0 {
			return firstSetPath
		}

		// We've got nothing
		return ""
	}
	return os.Getenv("HOME")
}

/**
 * @brief: Create a directory if it does not exist
 */
func CreateDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}
}

/**
 * @brief: Create a file if it does not exist
 */
func CreateFile(file string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		os.Create(file)
	}
}

func CheckFile(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return false
	} else {
		return true
	}
}
