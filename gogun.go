//===============================================
// Prototype

//https://gowebexamples.com/websockets/
//https://flaviocopes.com/go-date-time-format/
//https://gowebexamples.com/json/
//https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/
//https://stackoverflow.com/questions/14289256/cannot-convert-data-type-interface-to-type-string-need-type-assertion
//https://stackoverflow.com/questions/37011799/how-to-pass-a-child-struct-into-a-function-accepting-parent-struct
package gogun

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/websocket"
	"time"
)
//===============================================
// Web Sokcet Config
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

//===============================================
// Test
func readfiledata(){
	// read in the contents of the localfile.data
	data, err := ioutil.ReadFile("localfile.data")
	// if our program was unable to read the file
	// print out the reason why it can't
	if err != nil {
		fmt.Println(err)
	}
 
	// if it was successful in reading the file then
	// print out the contents as a string
	fmt.Print(string(data))
}
func writefiledata(){
	mydata := []byte("all my data I want to write to a file")
	// the WriteFile method returns an error if unsuccessful
	err := ioutil.WriteFile("localfile.data", mydata, 0777)
	// handle this error
	if err != nil {
  		// print it out
  		fmt.Println(err)
	}
}

var dup DupI
var store StoreI

func init(){
	fmt.Println("func init")
	dup := Dup{}
	//fmt.Println(dup)
	//fmt.Println(dup.opt.max) // nope
	dup.optdefault()

	store := Store{}//work
	//store = Store{opt:{file:"radata"}} //nope
	store.setFile("radata")
	store.opt.file = "radata"
	//store.put("radataa","test",nil)
	store.put("","test",nil)

	start := time.Now()
	fmt.Println(start)
}


//===============================================
// GUN FUNC's

var graph interface{}

type GunI interface {
	Get()
	Map()
	Put()
	Once()
	Test() string
}

var peers []string

//https://forum.golangbridge.org/t/solved-map-string-interface-best-practice/5457
type Gun []map[string]interface{}
//func (csvd CSVDatum ) ToString() string  {...}

func (g Gun) Get() {
	fmt.Println("func Get")
}

func (g Gun) Map() {
	fmt.Println("func Map")
}

func (g Gun) Put() {
	fmt.Println("func Put")
}

func (g Gun) Once() {
	fmt.Println("func Once")
}

func (g Gun) Test() string{
	fmt.Println("func Test")
  	return "test"
}

func Test(a, b int) int {
	return a - b
}

func Add(a, b int) int {
	return a + b
}
//===============================================
// Web Server

func WSEndpoint(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	
	var myMsg map[string]interface{}
	//var myJson interface{}//nope

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		
		json.Unmarshal([]byte(msg), &myMsg)
		//fmt.Println( myMsg["#"])//works
		fmt.Println( msgType)//works
		//id := myMsg['#']
		//https://stackoverflow.com/questions/14289256/cannot-convert-data-type-interface-to-type-string-need-type-assertion
		dup.check(fmt.Sprint(myMsg["#"]))

		/*
		for k, v := range myMsg {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case float64:
				fmt.Println(k, "is float64", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}
		*/
		// Print the message to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

//func web(h http){
	//fmt.Println("init htpp...",h)
	//h.HandleFunc("/gun",wsEndpoint)
//}
