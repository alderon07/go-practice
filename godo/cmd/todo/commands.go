package main

import (
	"fmt"
	"log"
	"time"
)

const (
    GetItDoneNow = 3
    IGotTime = 2
    WeAreChilling = 1
)

type Todo struct {
    Id int                  `json:"id"`
    Title string            `json:"title"`
    DueAt time.Time         `json:"due_at,omitempty"`
    CompleteAt time.Time    `json:"complete_at,omitempty"`
    CreatedAt time.Time     `json:"created_at"`
    
    Repeat string           `json:"repeat,omitempty"`
    Tags []string           `json:"tags,omitempty"`
    Priority int8           `json:"priority,omitempty"`
    AfterTodo int           `json:"after_todo,omitempty"`
}

// Helper for consistent error handling
func must(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func RunAdd(title string, due string, repeat, tags, priority, after_todo string) {
    if title == "" {
        log.Fatal("Error: -title is required")
    }
    tasks, err := store.Load[core.Task]()
    must(err)

    var due *time.Time
    if whenStr != "" {
        if t, err := core.ParseWhen(whenStr, time.Now()); err == nil {
            due = &t
        } else {
            log.Fatalf("Invalid -when: %v", err)
        }
    } else if dueStr != "" {
        if t, err := time.Parse("2006-01-02", dueStr); err == nil {
            due = &t
        } else {
            log.Fatalf("Invalid -due: %v", err)
        }
    }

    tasks = core.Add(tasks, title, due)
    i := len(tasks) - 1
    tasks[i].Priority = priority
    if repeat != "" {
        tasks[i].Repeat = repeat
    }
    tasks[i].Tags = core.ParseTags(tags)
    tasks[i].DependsOn = core.ParseIDs(after)

    must(store.Save(tasks))
    fmt.Println("Added:", title)
}

func RunList(showAll bool, grep, tags, sortKey, before, after string) {
    tasks, err := store.Load[core.Task]()
    must(err)
    visible := core.SortedWith(tasks, showAll, grep, sortKey)
    visible = core.FilterByTags(visible, tags)
    visible = core.FilterByDate(visible, before, after)

    if len(visible) == 0 {
        fmt.Println("(no tasks)")
        return
    }
    for i, t := range visible {
        status := " "
        if t.DoneAt != nil {
            status = "x"
        }
        fmt.Printf("%2d. [%s] %s", i+1, status, t.Title)
        if t.Priority > 1 {
            fmt.Printf("  (p%d)", t.Priority)
        }
        if t.Due != nil {
            fmt.Printf("  (due %s)", t.Due.Format("2006-01-02"))
        }
        if len(t.Tags) > 0 {
            fmt.Printf("  #%s", t.Tags)
        }
        fmt.Println()
    }
}

// func RunDone(indexStr string) {
// 	idx, err := core.Atoi1(indexStr)
// 	must(err)

// 	tasks, err := store.Load[core.Task]()
// 	must(err)
// 	visible := core.SortedWith(tasks, true, "", "due")

// 	tasks, err = core.MarkDone(tasks, visible, idx)
// 	must(err)
// 	must(store.Save(tasks))
// 	fmt.Println("Marked done:", visible[idx-1].Title)
// }

func RunRemove(indexStr string) {
    idx, err := core.Atoi1(indexStr)
    must(err)

    tasks, err := store.Load[core.Task]()
    must(err)
    visible := core.SortedWith(tasks, true, "", "due")

    tasks, err = core.Remove(tasks, visible, idx)
    must(err)
    must(store.Save(tasks))
    fmt.Println("Removed task", idx)
}

// func RunAlerts(watch bool, interval, ahead time.Duration) {
// 	for {
// 		tasks, err := store.Load[core.Task]()
// 		must(err)
// 		results := alerts.Scan(tasks, time.Now(), ahead)
// 		if len(results) == 0 {
// 			fmt.Println("(no alerts)")
// 		} else {
// 			fmt.Println("Upcoming / Overdue:")
// 			for _, a := range results {
// 				fmt.Printf("- %s (%s)\n", a.Task.Title, a.Status)
// 			}
// 		}
// 		if !watch {
// 			return
// 		}
// 		time.Sleep(interval)
// 	}
// }

func RunStats() {
    tasks, err := store.Load[core.Task]()
    must(err)
    fmt.Print(core.StatsReport(tasks, time.Now()))
}
