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
	GetOpt()
}

type OptConfig interface{
	Get() string
	Set(value string)
}

type InmemOpt struct{
	File string
}

func (o InmemOpt) Get() string{
	return o.File;
}

func (o InmemOpt) Set(value string){
	o.File = value;
}

type Store struct{
	Opt OptConfig
}

func (s Store) put(file string, data string, cb interface{}){
	filename := ""
	if len(file) > 0{
		fmt.Println("call")
		filename = file
	}else{
		fmt.Println("default")	
		filename = s.Opt.Get()
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
	filename := ""
	if len(file) > 0{
		fmt.Println("call")
		filename = file
	}else{
		fmt.Println("default")	
		filename = s.Opt.Get()
	}
	// read in the contents of the localfile.data
	data, err := ioutil.ReadFile(filename)
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
	s.Opt.Set(file)
	fmt.Println(file)
}

func (s Store) GetOpt(){
	fmt.Println("s.Opt.Get()")
	fmt.Println(s.Opt.Get())
}