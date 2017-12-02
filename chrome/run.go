package chrome

import (
	"context"
	"github.com/knq/chromedp"
	"github.com/knq/chromedp/runner"
	"fmt"
	"log"
	"sync"
	"errors"
)

type RunReply struct {
	Ctxt    context.Context
	Pool    *chromedp.Pool
	Cancel  context.CancelFunc
	Targets []*chromedp.Res
}

func RunPool(targetNumber int, address string) (RunReply, error) {
	var options = []runner.CommandLineOption{
		runner.ExecPath("/usr/bin/google-chrome"),
		runner.Flag("headless", true),
		runner.Flag("disable-gpu", true),
		runner.Flag("remote-debugging-address", address),
		runner.Flag("no-first-run", true),
		runner.Flag("no-default-browser-check", true),
	}

	ctxt, cancel := context.WithCancel(context.Background())

	startPort := 9222
	// create pool
	pool, err := chromedp.NewPool(chromedp.PortRange(startPort, startPort+targetNumber))
	if err != nil {
		return RunReply{}, err
	}

	targets := make([]*chromedp.Res, targetNumber)
	for i := 0; i < targetNumber; i++ {
		// allocate
		targets[i], err = pool.Allocate(ctxt, options...)
		if err != nil {
			return RunReply{}, fmt.Errorf("instance# %d error: %v", i, err)
		}
	}

	return RunReply{ctxt, pool, cancel, targets}, nil
}

func load(ctxt context.Context, wg *sync.WaitGroup, errChan chan error, target *chromedp.Res, action chromedp.Action) {
	defer wg.Done()

	if err := target.Run(ctxt, action); err != nil {
		errChan <- err
	}
}

func ExecActions(ctxt context.Context, targets []*chromedp.Res, actions []chromedp.Action) []error {
	targetNumber := len(targets)

	if targetNumber != len(actions) {
		return []error{errors.New("number of actions is not equal to number of targets")}
	}

	errChan := make(chan error, targetNumber)
	var wg sync.WaitGroup

	for i := 0; i < targetNumber; i++ {
		wg.Add(1)
		go load(ctxt, &wg, errChan, targets[i], actions[i])
	}
	// wait for to finish
	wg.Wait()

	var errs []error
loop:
	for {
		select {
		case err := <-errChan:
			errs = append(errs, err)
		default:
			break loop
		}
	}
	return errs
}

func ClosePool(cancel context.CancelFunc, targets []*chromedp.Res, pool *chromedp.Pool) {
	defer cancel()
	//release resources
	for _, target := range targets {
		target.Release()
	}
	// shutdown pool
	err := pool.Shutdown()
	if err != nil {
		log.Panic(err)
	}
}
