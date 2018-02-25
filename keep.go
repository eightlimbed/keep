// keep: write and view notes to yourself from the command line
// usage: ./keep <type your note here>
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

func get_filename() string {

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
        fmt.Println("usage: ./keep <type your note here>")
        os.Exit(1)
    }

    // get string from args
    args := os.Args[1:]

    // convert args (type []string) to type string
    keepsake := strings.Join(args, " ")

    // add a date stamp
    time := time.Now()
    //time_string := time.String()
    date_stamp := time.Format("Mon Jan 02")
    time_stamp := time.Format("03:04:05 PM")
    keepsake = date_stamp + " [" + time_stamp + "]: " + keepsake

    return keepsake
}

func main() {

    keepsake := make_entry()
    keepfile := get_filename()

    // open up the  file, or make it if it doesn't exist
    file, err := os.OpenFile(keepfile,
    os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // write the entry (keepsake) to the file
    fmt.Fprintf(file, keepsake + "\n")
}
