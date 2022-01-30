package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func retry(attempts int, sleep time.Duration, f func() ([]int, error)) ([]int, error) {
	var err error
	for i := 0; i < attempts; i++ {
		if i > 0 {
			log.Println("retrying after error:", err)
			time.Sleep(sleep)
			sleep *= 2
		}
		res, err := f()
		if err == nil {
			return res, nil
		}
	}
	return nil, fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

func works() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func fails() ([]int, error) {
	return nil, errors.New("this function had an error")
}

func main() {

	res, err := retry(3, time.Second, works)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	res, err = retry(3, time.Second, fails)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}
