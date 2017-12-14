package chrome

import (
	"context"
	"errors"
	"fmt"
	"github.com/korovkinand/chromedp"
	"github.com/korovkinand/surebetSearch/config/chrome"
	"log"
)

var ctx context.Context
var pool *chromedp.Pool
var cancel context.CancelFunc
var targets []*chromedp.Res

func RunPool(targetNumber int) error {
	var err error
	// create pool
	pool, err = chromedp.NewPool(chromedp.PortRange(chrome.StartPort, chrome.StartPort+targetNumber))
	if err != nil {
		return err
	}

	ctx, cancel = context.WithCancel(context.Background())

	targets = make([]*chromedp.Res, targetNumber)
	for i := 0; i < targetNumber; i++ {
		// allocate
		targets[i], err = pool.Allocate(ctx, chrome.Options...)
		if err != nil {
			cancel()
			return fmt.Errorf("instance (%d) error: %v", i, err)
		}
	}
	return nil
}

func load(errChan chan errorInfo, target int, action chromedp.Action) {
	errChan <- errorInfo{targets[target].Run(ctx, action), target}
}

func RunActions(actions []chromedp.Action) error {
	targetNumber := len(targets)

	if targetNumber != len(actions) {
		return errors.New("number of actions is not equal to number of targets")
	}

	errChan := make(chan errorInfo, targetNumber)
	for target := 0; target < targetNumber; target++ {
		go load(errChan, target, actions[target])
	}

	var errsInfo []errorInfo
	for target := 0; target < targetNumber; target++ {
		if errInfo := <-errChan; errInfo.err != nil {
			errsInfo = append(errsInfo, errInfo)
		}
	}
	if len(errsInfo) != 0 {
		return &GoroutinesError{errsInfo}
	}
	return nil
}

func ClosePool() {
	defer cancel()
	//Expects proper pool closing in fork of knq repo
	err := pool.Shutdown()
	if err != nil {
		log.Panic(err)
	}
}
