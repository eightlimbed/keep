// keep: write and view notes to yourself from the command line
// usage: ./keep <type your note here>
// config: set the path for the file you want to save your notes in.
// author: lee gaines (@eightlimbed) -- feb 25, 2018

package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "time"
)

func make_entry() string {

    // make sure to get at least one arg
    if len(os.Args) < 2 {
        fmt.Println("usage: ./keep <type your note here>")
        os.Exit(1)
    }

    // get string from args
    args := os.Args[1:]

    // convert args (type []string) to type string
    entry := strings.Join(args, " ")

    // add a date stamp
    time := time.Now()
    //time_string := time.String()
    date_stamp := time.Format("Mon Jan 02")
    time_stamp := time.Format("03:04:05 PM")
    entry = date_stamp + " [" + time_stamp + "]: " + entry

    return entry
}

func main() {

    path := os.Getenv("HOME")
    entry := make_entry()
    filename := path + "/notes.txt"

    // open up the  file, or make it if it doesn't exist
    file, err := os.OpenFile(filename,
    os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // write the entry (keepsake) to the file
    fmt.Fprintf(file, entry + "\n")
}
