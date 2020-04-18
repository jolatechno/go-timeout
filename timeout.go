package timeout

import (
  "time"
  "errors"
)

var (
  TimeOut = errors.New("Timed-out")
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
    return nil, TimeOut
  }
}

func MakeSimpleTimeout(f func() error, timeout time.Duration) error {
  errChan := make(chan error)

  go func(){
    errChan <- f()
    close(errChan)
  }()

  select {
  case err := <- errChan:
    return err

  case <- time.After(timeout):
    return TimeOut
  }
}

func MakeCheckerTimeout(f func() (interface{}, error), timeout time.Duration, checker func() error, ticking time.Duration) (interface{}, error) {
  dataChan := make(chan interface{})
  errChan := make(chan error)

  go func(){
    data, err := f()
    if err != nil {
      errChan <- err
    } else {
      dataChan <- data
    }
  }()

  ticker := time.NewTicker(ticking)

  for {
    select {
    case data := <- dataChan:
      ticker.Stop()
      return data, nil

    case err := <- errChan:
      ticker.Stop()
      return nil, err

    case <- time.After(timeout):
      ticker.Stop()
      return nil, TimeOut

    case <- ticker.C:
      err := checker()
      if err != nil {
        ticker.Stop()
        return nil, err
      }
    }
  }
}
