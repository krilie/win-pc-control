module github.com/lico603/win-pc-control

replace (
	golang.org/x/crypto v0.0.0 => github.com/golang/crypto v0.0.0
	golang.org/x/net v0.0.0 => github.com/golang/net v0.0.0
	golang.org/x/sys v0.0.0 => github.com/golang/sys v0.0.0
	golang.org/x/text v0.0.0 => github.com/golang/text v0.0.0
)

require (
	bou.ke/monkey v1.0.1 // indirect
	github.com/braintree/manners v0.0.0-20160418043613-82a8879fc5fd
	github.com/gin-contrib/gzip v0.0.0-20190101123152-0eb78e93402e
	github.com/gin-gonic/gin v1.3.0
	github.com/go-vgo/robotgo v0.0.0-20190125165610-5cb95037c5d1
	github.com/otiai10/mint v1.2.1 // indirect
)
