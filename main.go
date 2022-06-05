package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/ssup2/golang-pprof-example/pkg/block"
	"github.com/ssup2/golang-pprof-example/pkg/cpu"
	"github.com/ssup2/golang-pprof-example/pkg/memory"
	"github.com/ssup2/golang-pprof-example/pkg/mutex"
	"github.com/ssup2/golang-pprof-example/pkg/threadcreate"
)

func main() {
	// Run HTTP server to expose profile endpoint
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Set profile rate
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)

	// Run goroutines for load weighting
	go cpu.IncreaseInt()
	go cpu.IncreaseIntGoroutine()
	go memory.AllocMemory()
	go block.PrintHello()
	go block.PrintWorld()
	go threadcreate.CreateGoroutine1000()
	go mutex.Mutex01()
	go mutex.Mutex02()
	go mutex.Mutex03()

	// Block until receive a terminal signal
	log.Println("Waiting a terminal signal to shutdown gracefully")
	termSignal := make(chan os.Signal, 1)
	signal.Notify(termSignal, syscall.SIGTERM, syscall.SIGINT)
	<-termSignal
}
