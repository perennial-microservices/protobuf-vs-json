package main

import (
	"github.com/perennial-go-lang/protobuf-vs-json/studentpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
	"encoding/json"
	"log"
	"fmt"
)

func main() {
	var student1, student2, student3 studentpb.Student

	student1.Id = 1
	student1.FirstName = "Ajitem"
	student1.LastName = "Sahasrabuddhe"

	err := WriteProtoBuf("student.bin", &student1)
	if err != nil {
		log.Fatal(err)
	}
	err = WriteJson("student.json", &student1)
	if err != nil {
		log.Fatal(err)
	}

	err = ReadProtoBuf("student.bin", &student2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Read from Protobuf", student2)

	err = ReadJson("student.json", &student3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Read from JSON", student3)
}

func WriteProtoBuf(filename string, student *studentpb.Student) error {
	bytes, err := proto.Marshal(student)
	if err != nil {
		return err
	}
	ioutil.WriteFile(filename, bytes, 0644)
	return nil
}

func WriteJson(filename string, student *studentpb.Student) error {
	jsonFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(student)
	if err != nil {
		return err
	}
	return nil
}

func ReadProtoBuf(filename string, student *studentpb.Student) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(bytes, student)
	if err != nil {
		return err
	}

	return nil
}

func ReadJson(filename string, student *studentpb.Student) error {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(student)
	if err != nil {
		return err
	}
	return nil
}