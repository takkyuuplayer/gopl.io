package test

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

func waitForServer(url string) error {
	const timeout = 5 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%ss); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s faild to respond after %s", url, timeout)
}

func TestWaitForServer(t *testing.T) {
	err := waitForServer("https://404.takkyuuplayer.com")

	if err == nil {
		t.Errorf(`err = %#v, do NOT want %#v`, err, nil)
	}

	if fmt.Sprintf("%T", err) != "*errors.errorString" {
		t.Errorf(`Type of err = %T, want %s`, err, "*errors.errorString")
	}
}
