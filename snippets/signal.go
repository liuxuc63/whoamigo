package snippets

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func HandlerDemo() {
	/*
		// how to test, put the follow 2 lines on func main(){ }
		go snippets.HandlerSignal()
		snippets.Ping()
	*/
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	// Passing no signals to Notify means that
	// all signals will be sent to the channel.
	// signal.Notify(c)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
}

func handler(signal os.Signal) {
	if signal == syscall.SIGTERM {
		fmt.Println("Got kill signal. ")
		fmt.Println("Program will terminate now.")
		os.Exit(0)
	} else if signal == syscall.SIGINT {
		fmt.Println("Got CTRL+C signal")
		fmt.Println("Closing.")
		os.Exit(0)
	} else {
		fmt.Println("Ignoring signal: ", signal)
	}
}

func TestHandler() {
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl)
	exitchnl := make(chan int)

	go func() {
		for {
			s := <-sigchnl
			handler(s)
		}
	}()

	exitcode := <-exitchnl
	os.Exit(exitcode)
}
