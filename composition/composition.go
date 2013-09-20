package main
import "fmt"

type Person struct {
  Name string
}

func (p *Person) ProperName() string {
  return p.Name
}

type Woman struct {
  *Person
}

func (w *Woman) ProperName() string {
  return "Mrs. " + w.Person.ProperName()
}

func main() {
  jane := &Woman { Person: &Person { Name: "Jane Smith" } }
  fmt.Println(jane.ProperName())
  fmt.Println(jane.Person.ProperName())
}
