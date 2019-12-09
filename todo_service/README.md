# Todo list service
~  basic functionality for todo list

### PreRequired steps:
1. Install Golang from [Golang](https://golang.org/dl/) (`Go 1.11+` required)
2. Set `GOPATH` and `GOROOT`. For more details [Doc1](https://github.com/golang/go/wiki/SettingGOPATH) or [Doc2](https://blog.learngoprogramming.com/what-are-goroot-and-gopath-1231b084723f)
3. [Install Glide](https://github.com/Masterminds/glide)
4. [Install Dbmate](https://github.com/amacneil/dbmate)
5. [Install CockroachDb](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-windows.html)

### Getting started
1. Clone repo under `$GOPATH/src/test`. If `test` directory does not exist than first create it. Then run `git clone https://github.com/AneriShah2610/gin-gonic.git`
2. Run `glide install`
3. Create `.config` folder. Create `.env` file in this folder as per `.env.example`.
4. Run `go get github.com/amacneil/dbmate` to install dbmate
5. Add `DATABASE_URL` environment variable and set your database connection url in `DATABASE_URL` variable
6. Run `dbmate migrate` (Refer doc here: [Dbmate](https://github.com/amacneil/dbmate))
7. Run `go run main.go`
8. Open `http://localhost:8083` and hit query