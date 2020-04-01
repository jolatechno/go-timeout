package timeout

import (
  "time"
  "errors"

  "fmt"
)

func MakeTimeout(f func() (interface{}, error), timeout time.Duration) (interface{}, error) {
  dataChan := make(chan interface{})
  errChan := make(chan error)

  go func(){
    data, err := f()
    if err != nil {
      errChan <- err
    } else {
      dataChan <- data
    }

    close(dataChan)
    close(errChan)
  }()

  select {
  case data := <- dataChan:
    return data, nil

  case err := <- errChan:
    return nil, err

  case <- time.After(timeout):
    return nil, errors.New("Timed-out")
  }
}
