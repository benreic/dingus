package utils

import (
	"os"
)

/**
 * Determines if a file exists on disk
 *
 * @author Ben Reichelt <ben.reichelt@gmail.com>
 *
 * @param   string    The full path to the file to test
 * @return  bool
**/

func fileExists(fullPath string) bool {

	if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
		return true
	}

	return false

}
