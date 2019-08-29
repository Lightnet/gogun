/*
 * Prototype build
*/
//https://medium.com/rungo/error-handling-in-go-f0125de052f0

package gogun

import (
	"fmt"
	"io/ioutil"
	"log"
)

type StoreI interface{
	put(file string, data string, cb interface{})
	get(file string, cb interface{})
	list(cb interface{})
	setFile(file string)
}

type Store struct{
	opt struct{
		file string
	}
}

func (s Store) put(file string, data string, cb interface{}){
	filename := ""
	if len(file) > 0{
		fmt.Println("call")
		filename = file
	}else{
		fmt.Println("default")	
		filename = s.opt.file
	}
	fmt.Println(filename)
	mydata := []byte(data)
	// the WriteFile method returns an error if unsuccessful
	err := ioutil.WriteFile(filename, mydata, 0777)
	// handle this error
	if err != nil {
  		// print it out
  		fmt.Println(err)
	}
}

func (s Store) get(file string, cb interface{}){
	// read in the contents of the localfile.data
	data, err := ioutil.ReadFile(s.opt.file)
	// if our program was unable to read the file
	// print out the reason why it can't
	if err != nil {
		fmt.Println(err)
	}
	// if it was successful in reading the file then
	// print out the contents as a string
	fmt.Print(string(data))
}
//https://stackoverflow.com/questions/14668850/list-directory-in-go
func (s Store) list(cb interface{}){
	files, err := ioutil.ReadDir("./")
	if err != nil {
        log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func (s Store) setFile(file string){
	s.opt.file = file
	fmt.Println(file)
}