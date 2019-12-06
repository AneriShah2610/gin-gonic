# Basic examples

### PreRequired steps:
1. Install Golang from [Golang](https://golang.org/dl/) (`Go 1.11+` required)
2. Set `GOPATH` and `GOROOT`. For more details [Doc1](https://github.com/golang/go/wiki/SettingGOPATH) or [Doc2](https://blog.learngoprogramming.com/what-are-goroot-and-gopath-1231b084723f)
3. [Install Glide](https://github.com/Masterminds/glide)

### Getting started
1. Clone repo under `$GOPATH/src/test`. If `test` directory does not exist than first create it. Then run `git clone `
2. Run `glide install`
3. Run `go run main.go` 
4. Open any testing tool like postman or any other to hit query on host: `http://localhost:8080`

### Routes
1. Get `/` --> Get home page details
2. Post `/` --> Post home page details
3. Get `/query?name=abc&age=1` --> Get data based on query string
4. Get `/path/:name/:age` --> Get data based on path parameters
5. Post `/body` --> Post data with body 