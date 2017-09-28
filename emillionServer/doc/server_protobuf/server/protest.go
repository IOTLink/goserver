package main

import (
	"github.com/golang/protobuf/proto"
	"log"
	//"emillionServer/doc/server_protobuf/protest"
	"../protest"
	"fmt"
)

func main() {
	test := &protest.Test {
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &protest.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
	fmt.Println(newTest)
	// etc.
}