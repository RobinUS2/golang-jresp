# golang-jresp
JSON response helper for GoLang (useful for REST API services)

## Usage
```
jr := jresp.NewJsonResp()
jr.Set("name", "John Doe")
jr.OK()
jr.ToString(false) // Will print: {"parsetime":"1.988µs","ping":"pong","status":"OK"}
```

or in case of an error
```
jr := jresp.NewJsonResp()
jr.Error("Oops, something went wrong")
jr.ToString(false) // Will print:  {"parsetime":"1.988µs","error":"Oops, something went wrong","status":"ERROR"}
```

you can also pretty print
```
jr.ToString(true) // Will print the same JSON structure as usual, but then nicely formatted
```
