# gogun

## Status / Features:
 * Idea design prototype
 * write/read file (simple test)
 * check dup (simple one part function)
 * radisk (not yet build)
 * radix (not yet build)
 * ham (not yet build)
 * store (simple test)
 * websocket (simple send / receive )

## Information:
 This is gun.go prototype port from gun.js.
 It is base on video on how to build gun database.

 Link:
 * https://github.com/gundb/port
 * https://www.youtube.com/watch?v=5fCPRY-9hkc
 * https://gun.eco/docs/Porting-GUN
 * https://github.com/amark/gun/wiki/porting-gun

 Golang is base on type struct and interface build as well different way to setup packages. Javascript and golang code language have different format. It coded in different way a bit. It will take a while to copy or clone gunjs in simple format.

main.go
```go
package main

import (
	"fmt"
	"gogun"
)

func main() {

    fmt.Println("gun setup...")
    //this is just a test package.
    var Gun gogun.GunI = gogun.Gun{}
    Gun.Test()
    
}
```
Simple entry point to main

## setup: 
```
go get github.com/lightnet/gogun
go get github.com/lightnet/gohttpgunjs
```

```go
msg["#"] // tricky to get '#' golang format
msg["_"] //
```
Json format will take a while to get part of the gun work in golang.

callback function setup differently a bit of challenge.

## Notes:
 * Upper case first character letter is public function call.
 * Lower case first character letter is private function call.
 * There no class and constructor which will be hard to setup. That is different way to setup from information that I search around a bit. It work in progress test.
 * pacakge in group in files will able to call function to other ones that share if package is same location folder.
 * There is another gun golang version which seem be better. Link below.

## @cretz github (WARNING: This is an early proof-of-concept alpha version. Many pieces are not implemented or don't work.)
 * https://github.com/amark/gun/issues/610
 * https://github.com/cretz/esgopeta
 * https://godoc.org/github.com/cretz/esgopeta/gun