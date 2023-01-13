# Network Error Logging

[![Go Reference](https://pkg.go.dev/badge/github.com/zakarynichols/nel.svg)](https://pkg.go.dev/github.com/zakarynichols/nel)

The NEL (Network Error Logging) specification was designed to provide a way for web developers to send reports of network errors to a specified endpoint, which can be used for debugging and improving the performance of web applications.

## Installation

To install the `nel` package, you can use the `go get` command:

```sh
go get -u github.com/zakarynichols/nel
```

## Importing the package

```go
import "github.com/zakarynichols/nel"
```

## Creating a NEL struct

```go
n := &nel.NEL{
    ReportTo: "default",
    MaxAge: 2592000,
    // There are other fields to customize further.
}
```

## Validate struct

Validate all struct fields respective of the NEL specification.

```go
err := n.Validate()
if err != nil {
    // Handle error..
}
```

## Setting headers

```go
headers := make(http.Header)
headers.Add("NEL", n.String())
```

## Removing headers

When it comes to removing an existing NEL policy, the specification defines that a new NEL header with a max_age of 0 should be sent in the response. This is because the max_age field is used to specify the time, in seconds, that the user agent should retain the policy and send reports. By setting max_age to 0, the user agent is instructed to remove any existing NEL policy and stop sending reports.

This approach is useful in cases where the web developer wants to remove the NEL policy for a specific domain, and not for all domains. It also allows web developers to change or update the NEL policy for a specific domain without having to wait for the previous policy to expire.

Sending an empty NEL header or removing the NEL header entirely may not be compatible with all user agents and may not effectively remove the NEL policy for a specific domain.

```go
n := NEL{
    ReportTo: "default",
    MaxAge: 3600,
}

n.Remove() // Remove sets the max_age to zero.
w.Header.Set("NEL", n) // Set the NEL header yourself

// A function is exported from the nel package for convenience.
RemoveNEL(w, &n)

```
