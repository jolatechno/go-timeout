# go-timeout

Library use to create timeout from any function

## Usage

### _Example code :_

```
package main

import (
  "time"
  "fmt"
  "errors"

  "github.com/jolatechno/go-timeout"
)

func testFunc() (interface{}, error) {
  time.Sleep(2 * time.Second)
  return "returned", nil
}

func testFuncError() (interface{}, error) {
  time.Sleep(2 * time.Second)
  return nil, errors.New("Error")
}

func main() {
  fmt.Println(timeout.MakeTimeout(testFunc, 3 * time.Second))
  fmt.Println(timeout.MakeTimeout(testFunc, 1 * time.Second))
  fmt.Println(timeout.MakeTimeout(testFuncError, 3 * time.Second))
  fmt.Println(timeout.MakeTimeout(testFuncError, 1 * time.Second))
}
```

### _Output :_

```
returned <nil>
<nil> Timed-out
<nil> Error
<nil> Timed-out

```

## License

MIT
