package varint

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestVarInt(t *testing.T) {
	buf := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(buf, 300)
	fmt.Printf("300 % x \n", buf)
	b := []byte{0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x10, 0x00}
	b = []byte{0xac, 0x02, 0x00, 0x00}
	// b = []byte{0x08, 0x00}
	x, n := binary.Uvarint(b)
	fmt.Printf("%d % x %d \n", x, b, n)
}
