package helper

import (
	"bufio"
	"errors"
	"math/rand"
	"net"
	"time"
)

func ScanReq(scanner *bufio.Scanner) ([]byte, error) {
	if !scanner.Scan() {
		return nil, errors.New("error happened on scan client response")
	}

	b := scanner.Bytes()
	//log.Printf("received: %s", string(b))
	return b, nil
}

func SendResp(conn net.Conn, resp []byte) {
	conn.Write(resp)
	conn.Write([]byte("\n"))
}

func GetQuote() string {
	quotes := []string{
		"The best way out is always through",
		"Always Do What You Are Afraid To Do",
		"Believe and act as if it were impossible to fail",
		"Keep steadily before you the fact that all true success depends at last upon yourself",
		"The journey of a thousand miles begins with one step",
	}

	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(quotes))

	return quotes[n]
}
