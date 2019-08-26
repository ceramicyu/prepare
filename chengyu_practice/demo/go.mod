module demo

go 1.12

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190312061237-fead79001313

replace golang.org/x/net => github.com/golang/net v0.0.0-20190311183353-d8887717615a

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190308221718-c2843e01d9a2

replace golang.org/x/text => github.com/golang/text v0.3.0

require github.com/gin-gonic/gin v1.4.0 // indirect
