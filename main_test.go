package main

import (
	"github.com/perennial-go-lang/protobuf-vs-json/studentpb"
	"testing"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

var student1, student2, student3 studentpb.Student
var binaryData []byte
var jsonData bytes.Buffer

func init() {
	student1.Id = 1
	student1.FirstName = "Ajitem"
	student1.LastName = "Sahasrabuddhe"

	binaryData, _ = proto.Marshal(&student1)

	encoder := json.NewEncoder(&jsonData)
	encoder.Encode(&student1)
}

func BenchmarkWriteProtoBuf(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		proto.Marshal(&student1)
	}
}

func BenchmarkWriteJson(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(ioutil.Discard)
		encoder.Encode(&student1)
	}
}

func BenchmarkReadProtoBuf(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		proto.Unmarshal(binaryData, &student2)
	}
}

func BenchmarkReadJson(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		decoder := json.NewDecoder(&jsonData)
		decoder.Decode(&student3)
	}
}
