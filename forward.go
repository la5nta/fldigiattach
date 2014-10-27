package main

import (
	"io"
	"log"
	"net"
	"os"
)

func startForwarding(txAddr string, rxAddr string, pty *os.File) {
	go func() {
		if _, err := txFrom(txAddr, pty); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		if _, err := rxTo(rxAddr, pty); err != nil {
			log.Fatal(err)
		}
	}()
}

func txFrom(txAddr string, pty *os.File) (int64, error) {
	conn, err := net.Dial("udp", txAddr)
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	return io.Copy(conn, pty)
}

func rxTo(rxAddr string, pty *os.File) (int64, error) {
	addr, err := net.ResolveUDPAddr("udp", rxAddr)
	if err != nil {
		return 0, err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	return io.Copy(pty, conn)
}
