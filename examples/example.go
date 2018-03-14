// Copyright 2018 The go-toggl-reports AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/devtop/go-toggl-reports/togglreports"
	"os"
	"time"
)

const (
	// See your token at https://toggl.com/app/profile
	apiToken = "YOUR_API_TOKEN"

	// To get a workspace ID go to
	// https://toggl.com/app/workspaces -> Select settings of the workspace
	// -> see workspace ID in the URL
	workspaceID = YourWorkspaceID
)

func main() {
	c := togglreports.NewClient(apiToken)
	s, err := c.Summary.Get(workspaceID, nil)
	checkError(err)

	d := time.Duration(s.TotalGrand) * time.Millisecond
  fmt.Println("Total: ", d.String())

	d = time.Duration(s.TotalBillable) * time.Millisecond
  fmt.Println("Billable: ", d.String())

	// List of project entries
	for _, p := range s.Projects {
		d = time.Duration(p.Total) * time.Millisecond
		fmt.Println("- ", p.ID, p.Title.Name, "(", p.Title.Client, ") -> ", d.String())
	}

  fmt.Println()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
