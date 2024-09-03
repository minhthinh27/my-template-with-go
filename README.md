### Boilerplate Structure
<pre>├── <font color="#3465A4"><b>bootstrap</b></font>: -> declare struct to map file config and init config
│   ├── bootstrap.go
│   ├── database.go
│   └── .....go
├── <font color="#3465A4"><b>cmd</b></font>
│   └── cron -> start app with process job
│    └── main.go
│   └── server -> start app with process http
│    └── main.go
├── <font color="#3465A4"><b>configs</b></font> -> file config.yaml
│   ├── config.yaml 
├── <font color="#3465A4"><b>container</b></font> -> declare container to implement service(db,redis..)
│   └── container.go
│   └── database.go
│   └── ...go
├── <font color="#3465A4"><b>helper</b></font> -> common funcion
│   └── nl_cron
│    └── ....go
├── <font color="#3465A4"><b>internal</b></font> -> logic bussiness
│   └── biz -> handle logic
│   └── cron -> handle logic cron if any
│   └── data > interact data
│   └── entity -> declare entity
│   └── model -> declare model map to entity
│   └── server -> register restful,cron,authenticate,author...
│   └── service -> controller
├── <font color="#3465A4"><b>logger</b></font> -> init logger
├── <font color="#3465A4"><b>queue</b></font> -> init logger
├── <font color="#3465A4"><b>request</b></font> -> input
├── <font color="#3465A4"><b>response</b></font> -> output
├── <font color="#3465A4"><b>third_party</b></font> -> init 3rd party
│── .gitignore
│── go.mod
│── README.md
</pre>

### Installation
#### Local Setup Instruction
Follow these steps:
- Check your config in `configs/config.yaml`
- To add all dependencies for a package in your module `go get .` in the current directory
- Locally run `go run cmd/server/main.go` or `go build cmd/server/main.go` and run `./main`
- Check Application health available on [0.0.0.0:8000](http://0.0.0.0:8000)
