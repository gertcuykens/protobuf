package main

import (
	"encoding/json"
	testing "testing"

	"./task"
	proto2 "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
)

// const UINT64 = uint64(unsafe.Sizeof(uint64(0)))

//go:generate bash -c ./proto.sh

func BenchmarkMarshalJson(b *testing.B) {
	msg := &task.Task{
		Text: "test",
		Done: false,
	}

	for n := 0; n < b.N; n++ {
		_, err := json.Marshal(msg)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkUnMarshalJson(b *testing.B) {
	msg := &task.Task{
		Text: "test",
		Done: false,
	}

	bin, err := json.Marshal(msg)
	if err != nil {
		b.Error(err)
		return
	}

	var task task.Task
	for n := 0; n < b.N; n++ {
		err := json.Unmarshal(bin, &task)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkMarshalProto(b *testing.B) {
	msg := &task.Task{
		Text: "test",
		Done: false,
	}

	for n := 0; n < b.N; n++ {
		_, err := proto.Marshal(msg)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkUnMarshalProto(b *testing.B) {
	msg := &task.Task{
		Text: "test",
		Done: false,
	}

	bin, err := proto.Marshal(msg)
	if err != nil {
		b.Error(err)
		return
	}

	var t task.Task
	for n := 0; n < b.N; n++ {
		err := proto.Unmarshal(bin, &t)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkMarshalProto2(b *testing.B) {
	msg := &task.Task{
		Text: "test",
		Done: false,
	}

	for n := 0; n < b.N; n++ {
		_, err := proto2.Marshal(msg)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkUnMarshalProto2(b *testing.B) {
	msg := &task.Task{
		Text: "test",
		Done: false,
	}

	bin, err := proto2.Marshal(msg)
	if err != nil {
		b.Error(err)
		return
	}

	var t task.Task
	for n := 0; n < b.N; n++ {
		err := proto2.Unmarshal(bin, &t)
		if err != nil {
			b.Error(err)
			return
		}
	}
}
