package main

import (
	"fmt"
	"kv/store"
)

func fileStore() {
	fs := store.NewFileStore("my-fake-store-ID", ".fileStorey")
	err := fs.Set("demo_key", "This is my value from file")
	if err != nil {
		panic(err)
	}
	value, err := fs.Get("demo_key")
	if err != nil {
		panic(err)
	}
	fmt.Println("MS Value:", value)
}

func memoryStore() {
	ms := store.NewInMemoryStore("my-fake-store-ID")
	err := ms.Set("demo_key", "This is my value from memory")
	if err != nil {
		panic(err)
	}
	value, err := ms.Get("demo_key")
	if err != nil {
		panic(err)
	}
	fmt.Println("FS Value:", value)
	_ = ms.Delete("demo_key")
	fmt.Println("FS Value:", value)
}

func main() {
	fileStore()
	memoryStore()
}


//curl http://localhost:1024/kv/v1/store/set\?key\=D\&val\=55 // query-based request

//curl http://localhost:1024/kv/v1/store/set/A/11 // path-based request
