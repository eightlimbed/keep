// keep: save notes to yourself from the command line
// usage: keep <type your note here>
// options: -a : displays the contents of the notes file
// author: lee gaines (@eightlimbed) -- feb 25, 2018

package main

import (
    "fmt"
    "os"
    "os/exec"
    "log"
    "time"
    "syscall"
    "strings"
)

func view_entries(filename string) {

    // e.g: keep -a -- displays contents of the notes file

    // look for `less` on the path
    exe, look_err := exec.LookPath("less")
    if look_err != nil {
        log.Fatal(look_err)
    }

    // give it an argument
    args := []string{"less", filename}

    // syscall.Exec needs the environment to execute the process
    env := os.Environ()

    fmt.Println(args)
    exec_err := syscall.Exec(exe, args, env)
    if exec_err != nil {
        log.Fatal(exec_err)
    }
}

func make_entry(args string) string {

    // convert args (type []string) to type string
    entry := args

    // add a time stamp
    time := time.Now()
    date_stamp := time.Format("Mon Jan 02")
    time_stamp := time.Format("03:04:05 PM")
    entry = date_stamp + " [" + time_stamp + "]: " + entry

    return entry
}

func main() {

    // make sure to get at least one arg
    if len(os.Args) < 2 {
        fmt.Println("usage: ./keep <type your note here>")
        os.Exit(1)
    }

    // get string from args
    args := os.Args[1:]
    path := os.Getenv("HOME")
    filename := path + "/notes.txt"

    // check for options: -a
    if args[0] == "-a" {
        view_entries(filename)
        return
    }

    // default: write the entry to a file
    entry := strings.Join(args, " ")
    entry_string := make_entry(entry)

    file, err := os.OpenFile(filename,
                             os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Fprintf(file, entry_string + "\n")
    defer file.Close()
}
