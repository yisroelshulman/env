package env

import (
	"os"
	"strings"
)

// Load reads the .env file in the directory from which its called and loads the variables into ENV
// for this process.
//
// note: Load does not override any env variable that already exists.
func Load() error {
	return loadFile(".env")
}

// Read reads the .env file in the directory from which it is called from and returns a map of key,
// value pairs where key is the variable and value is the value to be assigned to the variable.
//
// This is if the user needs to read a .env file and doesnt want to store it in ENV
func Read() (map[string]string, error) {
	return readFile(".env")
}

// reads and loads variables in ENV
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

// reads key, value pairs from a file into a map
func readFile(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parse(file)
}
