package main

import (
	"bufio"
	"errors"
	"flag"
	"log"
	"net"
	"time"

	"github.com/vitamin-nn/tcp_pow/helper"
	"github.com/vitamin-nn/tcp_pow/pow"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	isVerified, err := verify(conn)
	if err != nil {
		log.Printf("error occured while verification: %v\n", err)
	}
	if isVerified {
		log.Println("sending phrase")
		helper.SendResp(conn, []byte(helper.GetQuote()))
	} else {
		helper.SendResp(conn, []byte("error: fail verificaton"))
	}

	log.Printf("successful closing connection with %s\n", conn.RemoteAddr())
}

func verify(conn net.Conn) (bool, error) {
	// sending pow algorhytm
	helper.SendResp(conn, []byte("hashcash"))

	scanner := bufio.NewScanner(conn)
	// getting response if client support algorythm
	support, err := helper.ScanReq(scanner)
	if err != nil {
		return false, err
	}

	if string(support) != "OK" {
		return false, errors.New("client does not support algorhytm")
	}

	// sending hash
	hash := getHash()
	helper.SendResp(conn, hash)

	// receiving nonce
	n, err := helper.ScanReq(scanner)
	if err != nil {
		return false, err
	}

	// checking solution
	return pow.Check(append(hash, n...)), nil
}

func getHash() []byte {
	return []byte(time.Now().String())
}

func main() {
	serverUrl := flag.String("server", ":3032", "server url")
	flag.Parse()
	l, err := net.Listen("tcp", *serverUrl)
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}
	defer l.Close()
	log.Printf("running server on: %s", *serverUrl)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("cannot accept: %v", err)
		}

		go handleConnection(conn)
	}
}
