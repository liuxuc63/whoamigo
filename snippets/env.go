package snippets

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TestEnv() {
	// Open the file
	file, err := os.Open("conf/env.cfg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a scanner
	scanner := bufio.NewScanner(file)

	// Read each line
	for scanner.Scan() {
		// Print the line
		line := strings.TrimSpace(scanner.Text())

		fmt.Println(line)
		if len(line) > 0 && line[0] != '#' {
			ll := strings.SplitN(line, "=", 2)
			if len(ll) == 2 {
				v1 := strings.Trim(strings.Trim(ll[1], "'"), "\"")
				os.Setenv(ll[0], v1)
			}
		}
	}
	fmt.Println(">> env ----")

	for _, str := range os.Environ() {
		fmt.Println(str)
	}

	//if line.startswith('#') or not line.strip():
	//os.Setenv("hell")
}
