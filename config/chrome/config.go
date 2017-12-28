package chrome

import (
	"github.com/korovkinand/chromedp/runner"
)

const address = "0.0.0.0"
const StartPort = 9222

var RunnerOpts = []runner.CommandLineOption{
	runner.ExecPath("/usr/bin/google-chrome"),
	runner.Flag("headless", true),
	runner.Flag("disable-gpu", true),
	runner.Flag("no-first-run", true),
	runner.Flag("no-default-browser-check", true),
	runner.Flag("remote-debugging-address", address),
}

var WithLog = false
