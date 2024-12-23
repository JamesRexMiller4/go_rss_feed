package main

import (
	"fmt"

	"github.com/JamesRexMiller4/go_rss_feed/internal/config"
)

func main() {
	str, err := config.GetConfigPath()
	if err != nil {
		fmt.Errorf("we made an oopsie")
	}
	fmt.Println(str)
}
