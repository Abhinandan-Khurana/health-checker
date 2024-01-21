package main

import (
	"fmt"
	"net"
	"time"

	"github.com/fatih/color"
)

func Check(destination string, port string) string {
	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string

	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	if err != nil {
		status = fmt.Sprintf("%s %v is unreachable, \n Error: %v", red("[DOWN]"), red(destination), err)
	} else {
		status = fmt.Sprintf("%s %v is reachable, \n %v %v \n %v %v", green("[UP]"), green(destination), blue("FROM:"), yellow(conn.LocalAddr()), blue("TO:"), yellow(conn.RemoteAddr()))
	}
	return status
}
