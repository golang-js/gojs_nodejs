package eventloop

import (
	"github.com/golang-js/gojs"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	const SCRIPT = `
	setTimeout(function() {
		console.log("ok");
	}, 1000);
	console.log("Started");
	`

	loop := NewEventLoop()
	prg, err := goja.Compile("main.js", SCRIPT, false)
	if err != nil {
		t.Fatal(err)
	}
	loop.Run(func(vm *goja.Runtime) {
		vm.RunProgram(prg)
	})
}

func TestStart(t *testing.T) {
	const SCRIPT = `
	setTimeout(function() {
		console.log("ok");
	}, 1000);
	console.log("Started");
	`

	prg, err := goja.Compile("main.js", SCRIPT, false)
	if err != nil {
		t.Fatal(err)
	}

	loop := NewEventLoop()
	loop.Start()

	loop.RunOnLoop(func(vm *goja.Runtime) {
		vm.RunProgram(prg)
	})

	time.Sleep(2 * time.Second)
	loop.Stop()
}

func TestInterval(t *testing.T) {
	const SCRIPT = `
	var count = 0;
	var t = setInterval(function() {
		console.log("tick");
		if (++count > 2) {
			clearInterval(t);
		}
	}, 1000);
	console.log("Started");
	`

	loop := NewEventLoop()
	prg, err := goja.Compile("main.js", SCRIPT, false)
	if err != nil {
		t.Fatal(err)
	}
	loop.Run(func(vm *goja.Runtime) {
		vm.RunProgram(prg)
	})
}
