# gogun

# Status:
 * Idea design prototype
 * write/read file
 * check dup

# Information:
 Note this is prototype test. Golang is base on type struct and interface build. Javascript function are not same format in golang. It coded in different way a bit. It will take a while to copy or clone gunjs in simple format.

# Notes:
 * Upper case first character letter is public function call.
 * Lower case first character letter is private function call.
 * There no class and constructor which will be hard to setup. That is different way to setup from information that I search around a bit. It work in progress test.
 * pacakge in group in files will able to call function to other ones that share if package is same location.


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