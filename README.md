# logging-for-otus

Implementing a small logger for Otus on Go.

## Installation

Run the following command from you terminal:


 ```bash
 go get github.com/koind/logging-for-otus
 ```

## Usage

Package usage example.

```go
package main

import (
	"github.com/koind/logging-for-otus"
	"bytes"
)

func main() {
	var b bytes.Buffer
    
	hwAccepted := logging_for_otus.NewHwAccepted(1, 25)
	logging_for_otus.LogOtusEvent(hwAccepted, &b)
	
	println(b.String()) // 2019-06-23 accepted 1 25 
}
```

## Available Methods

The following methods are available:

##### koind/logging-for-otus

```go
NewHwAccepted(id int, grade int) *HwAccepted
NewHwSubmitted(id int, code string, comment string) *HwSubmitted
LogOtusEvent(e OtusEvent, w io.Writer)
LogMessage() string
```

## Tests

Run the following command from you terminal:

```
go test -v .
```