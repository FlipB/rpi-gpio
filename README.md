# rpi-gpio

`rpi-gpio` is used to control GPIO output pin states from the command line.

## Usage

```sh
Usage of rpi-gpio:
  -high
        set to high
  -keep
        dont reset pin state on close
  -pin int
        GPIO Pin to set (required) (default -1)
  -verbose
        enable printouts
  -wait duration
        wait time after setting pin state, before close (default -1)
```

If no `wait` duration is specified then the program has to be stopped by signalling SIGINT or SIGTERM.

## Example

```sh
[nix-shell:~/rpi-gpio]$ # Set GPIO pin 14 to High, and keep it high even after rpi-gpio has terminated
[nix-shell:~/rpi-gpio]$ rpi-gpio -pin 14 -high -keep
[nix-shell:~/rpi-gpio]$ # Set GPIO pin 14 to Low
[nix-shell:~/rpi-gpio]$ rpi-gpio -pin 14
[nix-shell:~/rpi-gpio]$ # Set GPIO pin 15 to High for 10 seconds
[nix-shell:~/rpi-gpio]$ rpi-gpio -pin 15 -high -wait 15
```

# Nix files

This repository mainly acts as a playground for packaging with Nix and using Nix shell.
You can use the Nix package manager [https://nixos.org/] to install development dependencies
and build the contents of this repo.

## Nix development shell

```sh
[~/rpi-gpio]$ nix-shell

[nix-shell:~/rpi-gpio]$ # Check that go is installed
[nix-shell:~/rpi-gpio]$ go version
go version go1.15.2 linux/arm64 
[nix-shell:~/rpi-gpio]$ # Even the project was built and made available in PATH of nix-shell
[nix-shell:~/rpi-gpio]$ which rpi-gpio
/nix/store/8hd8vid6hjnmyssv40v6g4hwy9smsmqb-rpi-gpio/bin/rpi-gpio

[nix-shell:~/rpi-gpio]$ exit
[~/rpi-gpio]$ # nix shell exited, programs no longer available:
[~/rpi-gpio]$ which rpi-gpio
which: no rpi-gpio in ($PATH) 
```

## Building with Nix

```sh
$ nix-build . # this is same as nix-build ./default.nix
```
