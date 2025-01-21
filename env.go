package env

import (
	"os"
	"strings"
)

// Load reads the .env file in the directory from which its called and loads the variables into ENV
// for this process.
//
// note: Load does not override an env variable that already exists.
func Load() error {
	return loadFile(".env")
}

// Read
func Read() (map[string]string, error) {
	return readFile(".env")
}

func loadFile(filename string) error {
	envMap, err := readFile(filename)
	if err != nil {
		return err
	}

	rawEnv := os.Environ()
	currentEnv := map[string]bool{}
	for _, line := range rawEnv {
		key := strings.Split(line, "=")[0]
		currentEnv[key] = true
	}

	for key, value := range envMap {
		if !currentEnv[key] {
			os.Setenv(key, value)
		}
	}

	return nil
}

func readFile(filename string) (envMap map[string]string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	return parse(file)
}
