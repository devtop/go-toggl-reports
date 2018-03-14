go-toggl-reports
========

go-toggl-reports is Go library for accessing Toggl Reports API.

**Documentation:**
no Documentation yet :(

**Build Status:**
no build yet :(

## Basic Usage

~~~go
c := togglreports.NewClient("YOUR_API_TOKEN")

// fetches a summary report (for the last 7 days)
s, err := c.Summary.Get(workspaceID, nil)
checkError(err)

if err != nil {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
}

fmt.Println("Total: ", s.TotalGrand)
fmt.Println("Billable: ", s.TotalBillable)
~~~

Please see [examples](./examples) for a complete example.

## Todos / Roadmap

* Complete [summary report](togglreports/summary.go)
* Tests for [togglreports](togglreports/togglreports.go)
* Tests for [summary report](togglreports/summary.go)
* Build environment / CI
* Backlink from [Toggl API docs](https://github.com/toggl/toggl_api_docs/)
* Detailed report
* Weekly report
* Project dashboard

## Credits

* [go-toggl](https://github.com/gedex/go-toggl) inspired this library
* [Toggl API docs](https://github.com/toggl/toggl_api_docs/)

## License

This library is distributed under the BSD-style license found in the LICENSE.md file.
