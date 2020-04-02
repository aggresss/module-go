package demo

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aggresss/module-go/v2/delta"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	fmt.Printf("%s\n", delta.Version())

	select {
	case s := <-ch:
		fmt.Println(s.String())
	}
}
