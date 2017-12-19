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

func runPool(targetNumber int) error {
	var err error
	pool, err = chromedp.NewPool(chromedp.PortRange(chrome.StartPort, chrome.StartPort+targetNumber))
	if err != nil {
		return err
	}

	ctx, cancel = context.WithCancel(context.Background())

	targets = make([]*chromedp.Res, targetNumber)
	for i := 0; i < targetNumber; i++ {
		targets[i], err = pool.Allocate(ctx, chrome.Options...)
		if err != nil {
			ClosePool()
			log.Println("Pool allocating error")
			return fmt.Errorf("instance (%d) error: %v", i, err)
		}
	}

	log.Print("Pool allocated")
	return nil
}

func load(errChan chan errorInfo, target int, action chromedp.Action) {
	errChan <- errorInfo{targets[target].Run(ctx, action), target}
}

func RunActions(actions []chromedp.Action) error {
	targetNumber := len(targets)

	if targetNumber != len(actions) {
		return errors.New("numbers of actions and targets aren't equal")
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
		log.Println("Running actions error")
		return &GoroutinesError{errsInfo}
	}
	return nil
}

func InitPool(actions []chromedp.Action) error {
	if err := runPool(len(actions)); err != nil {
		return err
	}
	return RunActions(actions)
}

func ClosePool() {
	if cancel != nil {
		defer cancel()
	}
	if pool != nil {
		//Expects proper pool closing in fork of knq repo
		if err := pool.Shutdown(); err != nil {
			log.Panic(err)
		}
	}
	log.Print("Pool closed")
}
