package pow

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"log"
	"math"
)

const zeroCountTarget = 10
const maxNonceIterations = math.MaxUint32

func ResolveHashcash(hash []byte) ([]byte, error) {
	nonceSize := 4

	x := make([]byte, len(hash)+nonceSize)
	copy(x, hash)
	nonce := x[len(hash):]
	var i uint32

	for i < maxNonceIterations {
		binary.LittleEndian.PutUint32(nonce, i)
		if Check(x) {
			log.Printf("found nonce in %d iterations\n", i)

			return nonce, nil
		}
		i++
	}

	return nil, errors.New("unable to resolve hashcash")
}

func Check(b []byte) bool {
	zeroBitCount := getZeroBitCount(sha256.Sum256(b))

	return zeroBitCount > zeroCountTarget
}

func getZeroBitCount(x [32]byte) int {
	var result int
	for i := 0; i <= len(x); i++ {
		if x[i] != 0 {
			if x[i]&64 != 0x00 {
				result++
			} else if x[i]&32 != 0x00 {
				result += 2
				break
			} else if x[i]&16 != 0x00 {
				result += 3
				break
			} else if x[i]&8 != 0x00 {
				result += 4
				break
			} else if x[i]&4 != 0x00 {
				result += 5
				break
			} else if x[i]&2 != 0x00 {
				result += 6
				break
			} else if x[i]&1 != 0x00 {
				result += 7
			}

			break
		}
		result += 8
	}

	return result
}
