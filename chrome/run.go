package chrome

import (
	"context"
	"github.com/knq/chromedp"
	"github.com/knq/chromedp/runner"
	"fmt"
	"log"
	"sync"
	"errors"
	"time"
	"surebetSearch/common"
)

var ctx context.Context
var pool *chromedp.Pool
var cancel context.CancelFunc
var targets []*chromedp.Res

var options = []runner.CommandLineOption{
	runner.ExecPath("/usr/bin/google-chrome"),
	runner.Flag("headless", true),
	runner.Flag("disable-gpu", true),
	runner.Flag("no-first-run", true),
	runner.Flag("no-default-browser-check", true),
}

func RunPool(targetNumber int, address string) error {
	var err error
	options = append(options, runner.Flag("remote-debugging-address", address))
	startPort := 9222

	ctx, cancel = context.WithCancel(context.Background())

	// create pool
	pool, err = chromedp.NewPool(chromedp.PortRange(startPort, startPort+targetNumber))
	if err != nil {
		return err
	}

	targets = make([]*chromedp.Res, targetNumber)
	for i := 0; i < targetNumber; i++ {
		// allocate
		targets[i], err = pool.Allocate(ctx, options...)
		if err != nil {
			return fmt.Errorf("instance# %d error: %v", i, err)
		}
	}
	return nil
}

func load(timeout time.Duration, wg *sync.WaitGroup, errChan chan error, target *chromedp.Res, action chromedp.Action) {
	defer wg.Done()

	ctxTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel() // releases resources if slow operation completes before timeout elapses
	if err := target.Run(ctxTimeout, action); err != nil {
		errChan <- err
	}
}

func ExecActions(timeout time.Duration, actions []chromedp.Action) []error {
	targetNumber := len(targets)

	if targetNumber != len(actions) {
		return []error{errors.New("number of actions is not equal to number of targets")}
	}

	errChan := make(chan error, targetNumber)
	var wg sync.WaitGroup

	for i := 0; i < targetNumber; i++ {
		wg.Add(1)
		go load(timeout, &wg, errChan, targets[i], actions[i])
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

func ReloadTarget(number int, initLoad chromedp.Action, url string) error {
	var html string
	if err := targets[number].Run(ctx, chromedp.Tasks{
		GetHtml(&html),
		SaveScn(url),
	}); err != nil {
		return err
	}

	if err := common.SaveHtml(url, html); err != nil {
		return err
	}

	targets[number].Release() //Do not handle error

	var err error
	targets[number], err = pool.Allocate(ctx, options...)
	if err != nil {
		return fmt.Errorf("reallocate target# %d error: %v", number, err)
	}

	if err := targets[number].Run(ctx, initLoad); err != nil {
		return err
	}
	return nil
}

func ClosePool() {
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
