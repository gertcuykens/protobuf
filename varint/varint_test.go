package varint

import (
	"encoding/binary"
	"fmt"
)

func ExampleVarInt() {
	var d uint64 = 300
	x := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(x, 300)
	fmt.Printf("%d / % x / %d\n", d, x[:n], n)
	// x = []byte{0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x10, 0x00}
	// x = []byte{0x08, 0x00}
	// x = []byte{0xac, 0x02, 0x00, 0x00}
	d, n = binary.Uvarint(x[:n])
	fmt.Printf("%d / % x / %d\n", d, x[:n], n)
	// Output:
	// 300 / ac 02 / 2
	// 300 / ac 02 / 2
}
