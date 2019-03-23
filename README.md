# Garbage Generator
## Why?

I'm trying to get better at Go, so I'm porting some of the scripts I frequently use for pentesting from bash and Python to Go.


## Installation
`git clone https://github.com/iJackWilson/GarbageGenerator.git`
`go build generator.go`

## Usage
`./generator -length=1234 -filename=out.txt`

The default length is 100 characters

The default filename is set to output.txt

## Future plans
* Add an equivalent of [pattern_create.rb](https://github.com/rapid7/metasploit-framework/blob/master/tools/exploit/pattern_create.rb)
