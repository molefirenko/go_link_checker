# Link checked

Small Go project to check links status code.
Main goal is to check Go threading functionality

To process links just run application
```
go run main.go
```

This appliction use POST endpoint `localhost:8080/process`

Example of the boby
```
{
    "Links": [
        "http://www.golang.org/",
        "http://www.google.com/",
        "http://www.somestupidname.com/"
    ]
}
```
