// SPDX-FileCopyrightText: 2020 jecoz
//
// SPDX-License-Identifier: MIT

package main

import (
	"flag"
	"os"
	"fmt"
	"net"
	"strconv"

	"github.com/hypebeast/go-osc/osc"
)

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "error * "+format+"\n", args...)
}

func exitf(format string, args ...interface{}) {
	errorf(format, args...)
	os.Exit(1)
}

func makeMsg(fields ...string) *osc.Message {
	var msg *osc.Message
	switch len(fields) {
	case 0:
		msg = osc.NewMessage("/max/stop")
	default:
		msg = osc.NewMessage("/max/play")
		for _, v := range fields {
			msg.Append(v)
		}
	}

	return msg
}

func main() {
	u := flag.String("u", "localhost:5498", "OSC messages will be sent to this address. Network used will be UDP.")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "If no argument is provided, this program will send a /max/stop message with no payload to the target address, hackable with the `u` flag. If arguments are provided instead, this program will send each one of them as a separate payload string field to /max/play. The receiver expects the payload to be a path(s) to a local image file.\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	host, portraw, err := net.SplitHostPort(*u)
	if err != nil {
		exitf("unable to recognise target url: %v", err)
	}
	port, err := strconv.Atoi(portraw)
	if err != nil {
		exitf("unable to parse port: %v", err)
	}

	c := osc.NewClient(host, port)
	msg := makeMsg(os.Args[1:]...)

	fmt.Fprintf(os.Stderr, "** sending message to %v -> {%v}\n", *u, msg)
	if err := c.Send(msg); err != nil {
		exitf("unable to send message: %v", err)
	}
}