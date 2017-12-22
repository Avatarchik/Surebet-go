package chrome

import (
	"context"
	"github.com/korovkinand/chromedp"
	"github.com/korovkinand/surebetSearch/common"
	"github.com/korovkinand/surebetSearch/config/chrome"
	"log"
)

type action = chromedp.Action

var ctx context.Context
var cancel context.CancelFunc
var pool *chromedp.Pool
var targets []*chromedp.Res

func RunPool(targetNumber int) error {
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
			return common.NewStackErrorf("instance (%d) error: %v", i, err)
		}
	}

	log.Print("Pool allocated")
	return nil
}

func load(errChan chan errorInfo, target int, act action) {
	errChan <- errorInfo{targets[target].Run(ctx, act), target}
}

func RunActions(actions ...action) error {
	targetNumber := len(targets)

	if targetNumber != len(actions) {
		return common.NewStackError("numbers of actions and targets aren't equal")
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

func ClosePool() {
	if ctx != nil {
		cancel()
	}
	if pool != nil {
		//Expects proper pool closing in fork of knq repo
		if err := pool.Shutdown(); err != nil {
			log.Panic(err)
		}
	}
	log.Print("Pool closed")
}
