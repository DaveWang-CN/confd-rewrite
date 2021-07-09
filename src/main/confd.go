package main

import (
	"confdReWrite/src/backends"
	log "confdReWrite/src/log"
	"confdReWrite/src/resource/template"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {

	flag.Parse()

	// -version: print version info
	if config.PrintVersion {
		fmt.Printf("confd %s (Git SHA:%s, Go-Version:%s)\n", Version, GitSHA, runtime.Version())
		os.Exit(0)
	}

	log.Info("--------init confd config begin----------")

	//init basic config,backend config,template config
	if err := initConfig() ; err != nil {
		log.Fatal(err.Error())
	}

	log.Info("--------init confd config end----------")
	storeClient, err := backends.New(config.BackendsConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	config.StoreClient = storeClient
	if config.OneTime {
		if err := template.Process(config.ResourceConfig); err != nil {
			log.Fatal(err.Error())
		}
		os.Exit(0)
	}

	stopChan := make(chan bool)
	doneChan := make(chan bool)
	errChan := make(chan error, 10)

	var processor template.Processor
	switch {
	case config.Watch:
		processor = template.WatchProcessor(config.ResourceConfig, stopChan, doneChan, errChan)
	default:
		processor = template.IntervalProcessor(config.ResourceConfig, stopChan, doneChan, errChan, config.Interval)
	}

	go processor.Process()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case err := <-errChan:
			log.Error(err.Error())
		case s := <-signalChan:
			log.Info(fmt.Sprintf("Captured %v. Exiting...", s))
			close(doneChan)
		case <-doneChan:
			os.Exit(0)
		}
	}
	os.Exit(0)
}
