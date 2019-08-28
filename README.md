# gogun

# Status:
 * Idea design prototype

# Information:
 Note this is prototype. It base on type struct build for golang. Javascript function are not same format in golang.

# Notes:
 * Upper case first character letter is public function call.
 * Lower case first character letter is private function call.
 * There no class and constructor which will be try to setup. That is different way to setup from infomration that I search around a bit. It work in progress test.


main.go
```go
package main

import (
	"fmt"
	"gogun"
)

func main() {

    fmt.Println("gun setup...")

    var Gun gogun.GunI = gogun.Gun{}
    Gun.Test()
    
}
```
Simple entry point to main