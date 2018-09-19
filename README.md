# SmartHR SDK for Go

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/ktsujichan/smarthr-sdk-go/issues)

SmartHR API client library written in Golang.

## Install
```
go get -u github.com/ktsujichan/smarthr-sdk-go
```

## How to Use
```golang
package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/ktsujichan/smarthr-sdk-go"
	"gopkg.in/yaml.v2"
)

func main() {
	buf, err := ioutil.ReadFile("smarthr.yaml")
	if err != nil {
		panic(err)
	}
	var config smarthr.Config
	yaml.Unmarshal(buf, &config)
	c, _ := smarthr.NewClient(&config)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	crews, _ := c.ListCrews(ctx, nil)
	fmt.Println(crews)
}
```
