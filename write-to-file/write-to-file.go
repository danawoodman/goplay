package main

import (
  "io"
  "os"
)

func main() {
  // open input file
  input_file, error := os.Open("input.txt")
  if error != nil { panic(error) }

  // close file on exit and check for its returned error
  defer func() {
    if error := input_file.Close(); error != nil {
      panic(error)
    }
  }()

  // open output file
  output_file, error := os.Create("output.txt")
  if error != nil { panic(error) }
  // close output_file on exit and check for its returned error
  defer func() {
    if error := output_file.Close(); error != nil {
      panic(error)
    }
  }()

  // make a buffer to keep chunks that are read
  buffer := make([]byte, 1024)
  for {
    // read a chunk
    n, error := input_file.Read(buffer)
    if error != nil && error != io.EOF { panic(error) }
    if n == 0 { break }

    // write a chunk
    if _, error := output_file.Write(buffer[:n]); error != nil {
      panic(error)
    }
  }
}
