module example.com/hello

go 1.21.10

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote/v4 v4.0.1
)

replace example.com/greetings => ../greetings

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)
