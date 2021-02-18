package main

import (
	"fmt"
	"kv/store"
)

func fileStore()  {
	fs := store.NewFileStore("my-fake-store-ID", ".fileStore")
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

func memoryStore()  {
	ms := store.NewMemoryStore("my-fake-store-ID")
	err := ms.Set("demo_key", "This is my value from memory")
	if err != nil {
		panic(err)
	}
	value, err := ms.Get("demo_key")
	if err != nil {
		panic(err)
	}
	fmt.Println("FS Value:", value)
}

func main(){
	memoryStore()
	fileStore()
}
