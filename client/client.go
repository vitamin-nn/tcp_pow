package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"net"

	"github.com/vitamin-nn/tcp_pow/helper"
	"github.com/vitamin-nn/tcp_pow/pow"
)

func interractRoutine(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	// receiving pow algorithm
	_, err := helper.ScanReq(scanner)
	if err != nil {
		log.Printf("error occured while request PoW algorhytm: %v", err)

		return
	}

	helper.SendResp(conn, []byte("OK"))

	// receiving hash
	hash, err := helper.ScanReq(scanner)
	if err != nil {
		log.Printf("error occured while request hash: %v", err)

		return
	}

	// resolve puzzle
	nonce, err := pow.ResolveHashcash(hash)
	if err != nil {
		log.Printf("error occured while resolve hashcash: %v", err)
	}

	// sending nonce
	helper.SendResp(conn, nonce)

	// receiving phrase
	quote, err := helper.ScanReq(scanner)
	if err != nil {
		log.Printf("error occured while request quote: %v", err)

		return
	}
	log.Printf("recived quote: %s", quote)

	log.Println("success")
}

func main() {
	serverUrl := flag.String("server", "0.0.0.0:3032", "server url")
	flag.Parse()
	dialer := &net.Dialer{}
	conn, err := dialer.DialContext(context.Background(), "tcp", *serverUrl)
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}
	defer conn.Close()

	interractRoutine(conn)
}
