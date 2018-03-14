// Copyright 2018 The go-toggl-reports AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package togglreports

import (
	"fmt"
	"time"
)

// SummaryReportService handles communication with the summary report related
// methods of the Toggl Reports API.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/reports/summary.md
type SummaryService struct {
	client *Client
}

// SummaryReport represents a summary report.
type Summary struct {
	TotalGrand      int        `json:"total_grand,omitempty"`
	TotalBillable   int        `json:"total_billable,omitempty"`
	TotalCurrencies []Currency `json:"total_currencies,omitempty"`
	Projects        []Project  `json:"data,omitempty"`
}

// Project represents a project entry in the detailed data of the summary report
type Project struct {
	ID    int    `json:"id,omitempty"`
	Total int    `json:"time,omitempty"`
	Title *Title `json:"title,omitempty"`
}

// Title represents the project name and client name
type Title struct {
	Name   string `json:"project,omitempty"`
	Client string `json:"client,omitempty"`
}

type Currency struct {
	Name   string `json:"currency,omitempty"`
	Amount int    `json:"amount,omitempty"`
}

type Selectparameters struct {
	start *time.Time
	end   *time.Time
}


// List time entries. With start and end parameters you can specify
// the date range of the time entries returned.

// Get SummaryReport.
// workspaceID must be specified
// With start and end parameters you can specify  the date range of
// the time entries returned.
// If start and end are not specified, report starts 7 days ago
// and ends today (Toggl standard).
// Start and end will be casted into ISO 8601 date strings (daywise accuracy).
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/reports/summary.md#request
func (s *SummaryService) Get(wid int, selection *Selectparameters) (*Summary, error) {
	u := "summary"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

  params := req.URL.Query()
  params.Add("workspace_id", fmt.Sprintf("%v", wid))

	if selection != nil {
		if selection.start != nil {
			params.Add("since", selection.start.Format(time.RFC3339))
		}
		if selection.end != nil {
			params.Add("until", selection.end.Format(time.RFC3339))
		}
	}
//
//  url.Add("until", "2018-01-31")
  req.URL.RawQuery = params.Encode()

	data := new(Summary)
	_, err = s.client.Do(req, data)

	return data, err
}
