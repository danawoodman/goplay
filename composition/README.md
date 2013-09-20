# Object Composition

A simple example of OOP using Go; specifically composing objects together (similar to inheritance).

This creates a `Person` type and then a `Woman` type that embeds a `Person` type within it.

It also shows an example usage of pointers in Go.

## Explanation

In Go, there is no concept of inheritance or sub-classing. In fact, there are no classes at all!

However, that doesn't mean that Go is not capable of being an object oriented language; far from it. 

Go allows you to create methods (just simple functions, `func`) on any Struct. Structs can be composed of other structs and can include their respective methods and properties.

But Go approaches this differently than other languages in that when you compose one struct with another, you're actually including the whole struct into your new struct. So, for example, if we create a `Person` struct (as in `composition.go`):

```go
type Person struct {
  Name string
}

// Add a method `ProperName` onto the `Person` struct.
func (p *Person) ProperName() string {
  return p.Name
}

// Add another method to the `Person` struct.
func (p *Person) Species() string {
  return "Homo sapien"
}

```

We can then compose a new struct with the `Person` struct attached (or embedded) into it:

```go
// This embeds the `Person` struct into the `Woman` struct so you have 
// access to the `Person` structs methods and properties.
type Woman struct {
  *Person
}

// The attaches a method `ProperName` onto the `Woman` struct, but does not 
// override the method in the `Person` struct. Instead, you have access to 
// the original `Person` struct and can call it's properties and methods.
func (w *Woman) ProperName() string {
  return "Mrs. " + w.Person.ProperName()
}
```

Now, if you create a new instance of the `Woman` struct, you can access it's composed struct `Person` as well:

```go
jane := &Woman { Person: &Person { Name: "Jane Smith" } }
jane.ProperName()) // => "Mrs. Jane Smith"
jane.Person.ProperName() // => "Jane Smith"
jane.Species() // => "Homo sapien"
jane.Person.Species() // => "Homo sapien"
```

What this shows is that when composing structs, you still have complete access to the embedded struct, while also getting shortcuts for methods on the embedded struct if you have not redefined the method on your new struct. So, `jane.Species()` and `jane.Person.Species()` call the same method on the `Person` struct.

## Usage

As with most other of the examples, you build the `.go` file and run it:

```bash
go build
./composition
```
