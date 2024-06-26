package selectp

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

/*
In our case, we don't care what type is sent to the channel,
we just want to signal we are done and closing the
channel works perfectly!

Why struct{} and not another type like a bool? Well,
a chan struct{} is the smallest data type available
from a memory perspective so we get no allocation versus a bool.
Since we are closing and not sending anything on the chan,
why allocate anything?
*/
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
