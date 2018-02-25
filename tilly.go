// tilly: write and view notes to yourself from the command line
// usage: tilly <write your note here>
// config: set the path for the file you want to save your notes in.
// author: lee gaines (@eightlimbed) -- feb 25, 2018

package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strings"
    "time"
)

func get_file_pathname() string {

    var tokens []string
    var path string

    // open config file for reading
    file, err := os.OpenFile("config", os.O_RDONLY, 644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // read file line-by-line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // get file path name by tokenizing the line starting with "PATH="
        if strings.HasPrefix(scanner.Text(), "PATH") {
            tokens = strings.Split(scanner.Text(), "=")
            path = tokens[1]
        }
    }
    return path
}

func make_entry() string {

    // make sure to get at least one arg
    if len(os.Args) < 2 {
        fmt.Println("usage: tilly <your thoughts here>")
        os.Exit(1)
    }

    // get string from args
    args := os.Args[1:]

    // convert args (type []string) to type string
    tilly := strings.Join(args, " ")

    // add a date stamp
    time := time.Now()
    stamp := time.Format("Monday Jan 1: ")
    tilly = stamp + tilly

    return tilly
}

func main() {

    entry := make_entry()
    filename := get_file_pathname()

    // open up the tilly file, or make it if it doesn't exist
    file, err := os.OpenFile(filename,
    os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }

    // write to the file
    fmt.Fprintf(file, entry + "\n")

    // auto close file before main ends
    defer file.Close()
}
