# godo

A minimal, lightweight command-line TODO list manager written in Go.

## Features

- ✅ Simple and intuitive CLI interface
- 📝 Create tasks with titles, due dates, and priorities
- 🏷️ Tag system for task organization
- 🔍 Filter and search tasks
- 📊 Task analytics and statistics
- 🔗 Task dependencies
- 🔄 Recurring tasks support
- 💾 JSON-based storage

## Installation

### Build from source

```bash
make build
```

This will create a binary at `bin/todo`.

### Run directly

```bash
make run
```

## Usage

### Basic Commands

```bash
godo <command> [options]
```

### Available Commands

#### Add a new task

```bash
godo add -title "Task description" [options]
```

**Options:**

- `-title` (required): Task title/description
- `-due YYYY-MM-DD`: Set due date
- `-p <1-3>`: Set priority level (1=low, 2=medium, 3=high)
- `-tags "tag1,tag2"`: Add comma-separated tags
- `-repeat "daily|weekly|monthly"`: Set repeat rule
- `-after "1,2,3"`: Comma-separated dependency task IDs

**Example:**

```bash
godo add -title "Review pull requests" -due 2025-10-25 -p 3 -tags "work,urgent"
```

#### List tasks

```bash
godo list [options]
```

**Options:**

- `-all`: Show completed tasks as well
- `-grep "keyword"`: Filter by substring (case-insensitive)
- `-tags "tag1,tag2"`: Filter by tags (comma=OR, plus=AND)
- `-sort <key>`: Sort by: `due`, `priority`, `created`, or `status` (default: `due`)
- `-before YYYY-MM-DD`: Show tasks before date
- `-after YYYY-MM-DD`: Show tasks after date

**Examples:**

```bash
# List all pending tasks
godo list

# Show all tasks including completed
godo list -all

# Filter by tag
godo list -tags "work"

# Search tasks
godo list -grep "meeting"

# Sort by priority
godo list -sort priority
```

#### Remove a task

```bash
godo remove <index>
```

**Example:**

```bash
godo remove 3
```

#### View statistics

```bash
godo stats
```

Shows task analytics including completion rates and other metrics.

#### Help

```bash
godo help
```

#### Version

```bash
godo version
```

## Project Structure

```
godo/
├── cmd/
│   └── todo/
│       ├── main.go         # CLI entry point
│       └── commands.go     # Command implementations
├── internal/
│   ├── alerts/            # Alert/notification logic
│   ├── core/              # Core task management
│   │   ├── task.go        # Task operations
│   │   ├── filter.go      # Filtering utilities
│   │   ├── sort.go        # Sorting utilities
│   │   └── stats.go       # Statistics generation
│   └── store/             # Data persistence
│       ├── jsonstore.go   # JSON storage implementation
│       └── paths.go       # File path management
├── go.mod
├── Makefile
└── README.md
```

## Development

### Prerequisites

- Go 1.25 or higher

### Available Make Targets

```bash
# Display help with all available commands
make help

# Format, vet, test, and build (recommended before commits)
make all

# Build the binary
make build

# Build with specific version
make build VERSION=1.0.0

# Run the application
make run

# Run with arguments
make run ARGS="list -all"

# Run tests
make test

# Run tests with coverage report
make test-coverage

# Format code
make fmt

# Run go vet
make vet

# Run linter (requires golangci-lint)
make lint

# Clean build artifacts
make clean

# Install binary to GOPATH/bin
make install

# Download and tidy dependencies
make deps

# Development mode with auto-reload (requires entr)
make dev
```

The compiled binary will be available at `bin/godo`.

## Data Storage

Tasks are stored in JSON format. The storage location is managed by the internal store package.

## Priority Levels

The application supports three priority levels:

1. **Priority 1** (WeAreChilling): Low priority
2. **Priority 2** (IGotTime): Medium priority
3. **Priority 3** (GetItDoneNow): High priority

## Task Status

Tasks can be in one of two states:

- **Pending**: Active tasks that need to be completed
- **Done**: Completed tasks (shown with `[x]` marker when using `-all` flag)

## Future Features

The codebase includes foundation for additional features (currently commented out):

- ✨ Mark tasks as done/complete
- 🔔 Alert system for due/overdue tasks
- 👀 Watch mode for continuous monitoring

## License

This is a practice project. Feel free to use and modify as needed.

## Contributing

This is a personal practice project, but suggestions and improvements are welcome!
