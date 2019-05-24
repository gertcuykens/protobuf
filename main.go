package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
)

// const UINT64 = uint64(unsafe.Sizeof(uint64(0)))

//go:generate bash -c ./proto.sh
func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}
	var err error
	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list()
	case "add":
		err = add(strings.Join(flag.Args()[1:], " "))
	default:
		err = fmt.Errorf("unknown subcommand %s", cmd)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func add(text string) error {
	task := &Task{
		Text: text,
		Done: false,
	}

	b, err := proto.Marshal(task)
	if err != nil {
		return fmt.Errorf("marshal task %v", err)
	}

	f, err := os.OpenFile("tasks.pbf", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("open file %s", err)
	}

	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, uint64(len(b)))

	for i := 0; i < n; i++ {
		if err := binary.Write(f, binary.BigEndian, buf[i]); err != nil {
			return fmt.Errorf("write varint %s", err)
		}
	}

	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("write task %s", err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("close file %s", err)
	}
	return nil
}

func list() error {
	file, err := os.Open("tasks.pbf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return stream(bufio.NewReader(file))
}

func stream(buffer io.ByteReader) error {
	l, err := binary.ReadUvarint(buffer)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return fmt.Errorf("%v ----", err)
	}

	if l == 0 {
		return nil
	}

	b := make([]byte, l)

	for i := range b {
		b[i], err = buffer.ReadByte()
		if err != nil {
			return fmt.Errorf("%v", err)
		}
	}

	if err := task(b); err != nil {
		return fmt.Errorf("%v", err)
	}

	return stream(buffer)
}

func task(b []byte) error {
	var task Task
	if err := proto.Unmarshal(b, &task); err != nil {
		return fmt.Errorf("read task %v", err)
	}
	if task.Done {
		fmt.Printf("👍")
	} else {
		fmt.Printf("😱")
	}
	fmt.Printf(" %s\n", task.Text)
	return nil
}
