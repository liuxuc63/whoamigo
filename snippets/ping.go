package snippets

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"time"
)

func Ping() {
	// Create a mutex to synchronize access to the stdout pipe.
	var mutex sync.Mutex

	// Create a command to run the shell command.
	cmd := exec.Command("ping", "zyb.bur.ink")
	//cmd := exec.Command("bash", "-c", "echo 'sleep 2' && sleep 2 && ls /tmp/error1234 && ping zyb.bur.ink")
	//cmd.Process.Signal(syscall.SIGINT)

	cmd.Stderr = os.Stderr // will print error to terminal
	// Create a pipe to capture the stdout of the command.
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start the command.
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		time.Sleep(time.Second * 5)
		cmd.Process.Kill()
		//cmd.Process.Signal(syscall.SIGINT)
	}()

	// Goroutine to read the stdout of the command and print it synchronized.
	go func() {
		//Wait will close the pipe after seeing the command exit, so most callers need not close the pipe themselves.
		//defer stdoutPipe.Close()

		// Read the stdout of the command.
		reader := bufio.NewReader(stdoutPipe)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}

			// Lock the mutex to prevent concurrent access to the stdout.
			mutex.Lock()

			// Print the line of output.
			fmt.Println(line)

			// Unlock the mutex.
			mutex.Unlock()
		}
	}()

	// Wait for the command to finish.
	err = cmd.Wait()
	if err != nil {
		fmt.Println("cmd error:", err)
		return
	}
}
