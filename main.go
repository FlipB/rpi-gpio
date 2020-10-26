package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/flipb/rpi-gpio/rpi"
)

var verbose bool

func main() {

	pinNum := -1
	state := false
	sleepTime := time.Duration(2 * time.Second)
	dontUnexport := false
	flag.IntVar(&pinNum, "pin", pinNum, "GPIO Pin to set")
	flag.BoolVar(&state, "high", state, "set to high")
	flag.DurationVar(&sleepTime, "wait", sleepTime, "wait time after setting pin state, before close")
	flag.BoolVar(&dontUnexport, "keep", dontUnexport, "dont reset pin state on close")
	flag.BoolVar(&verbose, "verbose", verbose, "enable printouts")
	flag.Parse()

	if pinNum == -1 {
		println("Missing arguments!\n\n", filepath.Base(os.Args[0]), "usage:\n")
		flag.PrintDefaults()
		os.Exit(2)
	}

	if err := run(pinNum, state, sleepTime, dontUnexport); err != nil {
		fmt.Fprintf(os.Stderr, "error opening GPIO Pin %d: %v", pinNum, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run(pinNum int, state bool, sleep time.Duration, keep bool) error {
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

	time.Sleep(sleep)

	return pin.Close()
}

func log(msg string) {
	if verbose {
		println(msg)
	}
}
