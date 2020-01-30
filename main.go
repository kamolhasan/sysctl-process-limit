package main

import (
	"fmt"

	xos "github.com/m3db/m3/src/x/os"
	"github.com/the-redback/go-oneliners"
)

func main() {

	limits, err := xos.GetProcessLimits()
	if err != nil {
		fmt.Errorf("unable to determine process limits: %v", err)
	}
	oneliners.PrettyJson(limits)
}
