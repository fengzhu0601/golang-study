package main

import (
	"os"
	"os/signal"
	"syscall"
	"log"
	"fengzhu/golang-study/toml_config_hot/config"
	"fmt"
)

func main(){
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR1, syscall.SIGINT)

	fmt.Println(config.Config().Title)

	exitChan := make(chan int)
	go func(){
		for{
			sig := <-s
			switch sig {
			case syscall.SIGUSR1:
				config.ReloadConfig()
				log.Println("sigusr1 Reloaded config")
				fmt.Println(config.Config().Title)
			case syscall.SIGINT:
				fmt.Println(config.Config().Title)
				exitChan <- 0
			default:
				fmt.Println("unknown signal.")
			}
		}
	}()

	code := <- exitChan
	os.Exit(code)
}
