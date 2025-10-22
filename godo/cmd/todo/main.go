package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "time"
)

// Version info (optional: injected at build time via -ldflags "-X main.Version=1.0.0")
var Version = "dev"

func usage() {
    fmt.Printf(`godo — minimal todo CLI
Version: %s

Usage:
  godo <command> [options]

Commands:
  add       Add a new task
  ls        List tasks
  done      Mark task as complete
  rm        Remove a task
  alerts    Show due/overdue tasks
  stats     Show task analytics
  help      Show this help

Run "todo <command> -h" for detailed help.
`, Version)
}

func main() {
    log.SetFlags(0)
    if len(os.Args) < 2 {
        usage()
        os.Exit(2)
    }

    cmd := os.Args[1]
    args := os.Args[2:]

    switch cmd {
    case "add":

        addFlags := flag.NewFlagSet("add", flag.ExitOnError)
        // params for add cmd
        title := addFlags.String("title", "", "Task title (required)")
        dueStr := addFlags.String("due", "", "Due date YYYY-MM-DD")
        // when := addFlags.String("when", "", `Natural date ("tomorrow", "in 3d", "next mon")`)
        repeat := addFlags.String("repeat", "", "Repeat rule: daily|weekly|monthly")
        priority := addFlags.Int("p", 1, "Priority (1-3)")
        tags := addFlags.String("tags", "", "Comma-separated tags")
        after := addFlags.String("after", "", "Comma-separated dependency task IDs")
        _ = addFlags.Parse(args)
        commands.RunAdd(title, dueStr, repeat, priority, tags, after)

    case "list":
        lsFlags := flag.NewFlagSet("ls", flag.ExitOnError)
        showAll := lsFlags.Bool("all", false, "Show completed tasks too")
        grep := lsFlags.String("grep", "", "Filter by substring (case-insensitive)")
        tags := lsFlags.String("tags", "", "Filter by tags (comma=OR, plus=AND)")
        sortKey := lsFlags.String("sort", "due", "Sort by: due|priority|created|status")
        before := lsFlags.String("before", "", "Filter tasks before YYYY-MM-DD")
        after := lsFlags.String("after", "", "Filter tasks after YYYY-MM-DD")
        _ = lsFlags.Parse(args)
        commands.RunList(*showAll, *grep, *tags, *sortKey, *before, *after)

    // case "done":
    // 	doneFlags := flag.NewFlagSet("done", flag.ExitOnError)
    // 	_ = doneFlags.Parse(args)
    // 	if doneFlags.NArg() < 1 {
    // 		log.Fatal("Usage: todo done <index>")
    // 	}
    // 	commands.RunDone(doneFlags.Arg(0))

    case "remove":
        rmFlags := flag.NewFlagSet("rm", flag.ExitOnError)
        _ = rmFlags.Parse(args)
        if rmFlags.NArg() < 1 {
            log.Fatal("Usage: todo rm <index>")
        }
        commands.RunRemove(rmFlags.Arg(0))

    // case "alerts":
    //     alertFlags := flag.NewFlagSet("alerts", flag.ExitOnError)
    //     watch := alertFlags.Bool("watch", false, "Continuously monitor for upcoming tasks")
    //     interval := alertFlags.Duration("interval", 60*time.Second, "Polling interval")
    //     ahead := alertFlags.Duration("ahead", 24*time.Hour, "Lookahead window for alerts")
    //     _ = alertFlags.Parse(args)
    //     commands.RunAlerts(*watch, *interval, *ahead)

    case "stats":
        commands.RunStats()

    case "help", "-h", "--help":
        usage()

    case "version", "-v", "--version":
        fmt.Println(Version)

    default:
        fmt.Printf("Unknown command: %s\n\n", cmd)
        usage()
        os.Exit(2)
    }
}
