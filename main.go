// Copyright 2014 Martin Hebnes Pedersen, LA5NTA. All rights reserved.

// fldigiattach is a tool to allow use of fldigi as a modem for Linux's AX.25-stack.
//
// This program creates a pty and uses kissattach(8) to attach the KISS
// interface of fldigi to an axport. After attachment, this program will
// act as a proxy between the AX.25-stack and fldigi to allow AX.25 over fldigi.
//
// Because kissattach daemonizes, you must kill it (as normal) after execution.
//
// fldigi 3.22 or later is required (KISS interface), as is kissattach (ax25-tools).
//
//   Usage of fldigiattach:
//     -mtu=0: Sets the mtu of the interface [default is paclen parameter in axports].
//     -port="": Name of a port given in the axports file
//     -rx-addr="127.0.0.1:7343": fldigi's rx address
//     -tx-addr="127.0.0.1:7342": fldigi's tx address
//
package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"os/user"
)

const DefTxAddr = "127.0.0.1:7342"
const DefRxAddr = "127.0.0.1:7343"

var (
	argMtu  int
	argPort string

	argTxAddr string
	argRxAddr string
)

func init() {
	flag.StringVar(&argTxAddr, "tx-addr", DefTxAddr, "fldigi's tx address")
	flag.StringVar(&argRxAddr, "rx-addr", DefRxAddr, "fldigi's rx address")
	flag.StringVar(&argPort, "port", "", "Name of a port given in the axports file")
	flag.IntVar(&argMtu, "mtu", 0, "Sets the mtu of the interface [default is paclen parameter in axports].")
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	if argPort == "" {
		log.Fatal("Argument '-port' is required.")
	} else if argMtu > 0 && argMtu < 128 {
		log.Println("Warning: low paclen (mtu) will cause system to hang. This is probably a bug in kissattach.")
	}

	// Give a hint to the user if not running as root, this is required by kissattach.
	if !IsRoot() {
		log.Println("It seems you're not root. You probably should be...")
	}

	// Attach new pty to AX.25 stack
	pty, err := KissAttach(argPort, argMtu)
	if err != nil {
		log.Fatal(err)
	}
	defer pty.Close()

	// Forward traffic between pty and fldigi
	log.Println("Forwarding...")
	startForwarding(argTxAddr, argRxAddr, pty)

	// Wait for interrupt/kill
	sig := WaitFor(os.Interrupt, os.Kill)
	log.Printf("Got %s, exiting. Remember to kill kissattach manually!", sig)
}

func WaitFor(sig ...os.Signal) os.Signal {
	c := make(chan os.Signal)
	signal.Notify(c, sig...)
	return <-c
}

func IsRoot() bool {
	user, err := user.Current()
	if err != nil {
		return false
	} else {
		return user.Uid == "0"
	}
}
