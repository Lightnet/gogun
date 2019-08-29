/*
 * Prototype build
*/

package gogun

import (
	"fmt"
)

type DupI interface{
	check(id string)
	track()
	random()
	optdefault()
}

type Dup struct{
	s  map[string]interface{}
	opt struct {
		max int //`default:1000` //1000
		age int //`default:9000` //1000 * 9
	}
}
//https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
func (d Dup) check(id string) {
	fmt.Println("check")
	fmt.Println(id)
	if value, ok := d.s[id]; ok {
		fmt.Println("value: ", value)
	} else {
		fmt.Println("key not found")
	}
}

func (d Dup) track() {
	
}

func (d Dup) random() {
	
}

func (d Dup) optdefault(){
	d.opt.max = 1000;
	d.opt.age = 9000;
	fmt.Println(d.opt.max)
}