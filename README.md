# GoPlay

A set of [Golang](http://golang.org) examples, starting at the most basic "Hello World" example.

## Examples

- **Hello World**: A simple program that just outputs "Hello World"
- **Writing to a file**: Takes the contents of one file and writes it to another using pure Go.
- **Simple web server**: A simple web server that returns "hello world" when a certain path is visited, using pure Go.
- **Image Analysis**: Create a color histogram of an image.
- **Object Composition**: In Go, you cannot inherit or sub-class; instead you compose structs.

Coming soon:

- RESTful API
- Getting data from an API
- Daemon
- Accessing a database

## Install Go

If you're on a Mac and you use [Homebrew](http://brew.sh/) you can install go by running:

```bash
brew install go
```

... and you're ready!

If you don't have Homebrew, you can install [go in many other ways](http://golang.org/doc/install).

## Running Examples
g
Most examples can be run by just running `go build` and executing the generated binary file:

```bash
cd helloworld
go build
./helloworld
```

All the examples have a README file explaining what they do and how to run them.

## Go Resources

* [GOOD - Go Object Oriented Design](http://nathany.com/good/)
* [Great Golang intro focused on OO design in Go](http://areyoufuckingcoding.me/2012/07/25/object-desoriented-language/)

## Credits

Some examples are taken, sometime in their entirety, from other sources and we'll do our best to link back to them.

All other examples written by [Dana Woodman](http://danawoodman.com) of [Teleporter](http://teleporter.io).

## License

Examples are published under the MIT license (see attached).
