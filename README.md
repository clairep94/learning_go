# learning_go

commands
```bash
go env GOPATH #find go environment

go mod init hello #before running go test
go test -bench=.
go test -cover

go run hello.go
go build # make executable
go install # make executable in a bin file
```

Basic:
``` go
package main

import ( //no commas in list
  "fmt"
)

func main() { // entrypoint
  var x int // variable init with type 0

  a := [5]int{5, 4, 3, 2, 1}
  b := []int{}
  b = append(b, 5) //makes a copy

  vertices := make(map[string]int)
  vertices["square"] = 4
  delete(vertices, "square")


  // for loop:
  for i := 0; i < 5; i++ {
    fmt.Print("Looping")
  }

  // while loop:
  i := 0
  for i < 5 {
    fmt.Print("Looping")
    i++
  }

  // for i,v in range:
  for index, value := range a {
    fmt.Println("index:", index, "value:", value)
  }

  // for k,v in map:
  m := make(map[string]string)
  m["a"] = "alpha"
  m["b"] = "beta"

  for key, value := range m {
    fmt.Println("key:", key, "value:", value)
  }

  // Error handling -- go doesn't have exceptions
  if result, err := sqrt(16); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(result)
  }

  // structs
  p := person{name: "Jake", age: 23}

}

func sqrt(x float64) (float64, error) {
  if x < 0 {
    return 0, errors.New("Undefined")
  }

  return math.Sqrt(x), nil
}

type person struct {
  name string
  age int
}

```


Special to go:
```go
package main

import (
  "fmt"
)

func main() {
  i := 7
  fmt.Println(i) //7
  fmt.Println(&i) //memory loc --> pointer to i

  inc(i) // this returns 8, but doesn't change i, should be  i=inc(i), bc i here is a copy
  fmt.Println(i) //7

  inc(*i) // changes i directly
  fmt.Println(i) //8
  
}

func inc(x *int) {
  x++ //*x++
}
```


Concurrency:
- splitting sequences up into packages that can be done at the same time

```go
package main
import (
  "fmt"
  "time"
)

func main() {
  go count("sheep") //goroutine -- runs both at same time vs. sequence
  count("fish")
}

func count(thing string) {
  for i := 1; true; i++ {
    fmt.Println(i, thing)
    time.Sleep(time.Milliseconds * 500)
  }
}
```