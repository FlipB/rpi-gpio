package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/flipb/rpi-gpio/internal/rpi"
)

var verbose bool

func main() {

	pinNum := -1
	state := false
	sleepTime := time.Duration(-1)
	dontUnexport := false
	flag.IntVar(&pinNum, "pin", pinNum, "GPIO Pin to set (required)")
	flag.BoolVar(&state, "high", state, "set to high")
	flag.DurationVar(&sleepTime, "wait", sleepTime, "wait time after setting pin state, before close")
	flag.BoolVar(&dontUnexport, "keep", dontUnexport, "dont reset pin state on close")
	flag.BoolVar(&verbose, "verbose", verbose, "enable printouts")
	flag.Parse()

	if pinNum == -1 {
		fmt.Fprintf(os.Stderr, "Missing arguments!\n\n%s usage:\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(2)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	defer signal.Stop(sigChan)

	stopChan := make(chan struct{}, 0)
	go func() {
		var stopTimeChan <-chan time.Time
		if sleepTime > 0 {
			stopTimeChan = time.After(sleepTime)
		}
		select {
		case <-stopTimeChan:
			close(stopChan)
			log("Stopping due to timeout")
		case sig := <-sigChan:
			close(stopChan)
			log(fmt.Sprintf("Stopping due to signal %s", sig))
		}
	}()

	if err := run(pinNum, state, stopChan, dontUnexport); err != nil {
		fmt.Fprintf(os.Stderr, "error opening GPIO Pin %d: %v", pinNum, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run(pinNum int, state bool, stop chan struct{}, keep bool) error {
	opts := []rpi.Opt{}
	if keep {
		opts = append(opts, rpi.NoExportOnClose())
	}
	pin, err := rpi.OpenPin(pinNum, rpi.OUT, opts...)
	if err != nil {
		return err
	}

	if state {
		log("Setting High")
		err = pin.Write(rpi.HIGH)
	} else {
		log("Setting Low")
		err = pin.Write(rpi.LOW)
	}
	if err != nil {
		return fmt.Errorf("error setting pin state: %v", err)
	}

	<-stop
	if !keep {
		log("Resetting pin")
	} else {
		log("Stopping")
	}
	return pin.Close()
}

func log(msg string) {
	if verbose {
		println(msg)
	}
}
