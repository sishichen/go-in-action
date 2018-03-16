https://www.sohamkamani.com/blog/2018/02/18/golang-data-race-and-how-to-fix-it/

to check where data race cause
$ go run -race main.go

fix data racing
1. use waitGroup
2. use chan
3. use return chan/ mutiple return

ref: range loops with Channels
http://programming.guide/go/for-loop-range-array-slice-map-channel.html

