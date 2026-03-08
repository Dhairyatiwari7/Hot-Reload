# HotReload – Go CLI Hot Reload Engine

A lightweight CLI tool written in Go that automatically rebuilds and restarts a server when source code changes.

HotReload eliminates the need for developers to manually stop, rebuild, and restart servers during development. The tool continuously watches a project directory and triggers a rebuild whenever a file is modified.

---

## Problem

During development, engineers typically follow this repetitive workflow:

1. Modify source code
2. Stop the running server
3. Rebuild the project
4. Restart the server

This manual process slows down development and interrupts productivity.

HotReload automates this workflow by watching file changes and restarting the server automatically.

---

## Features

- Watches project directory recursively
- Automatically rebuilds the project on code changes
- Automatically restarts the server
- Streams server logs in real time
- Performs an initial build when the tool starts
- Supports nested project folders
- Lightweight CLI tool written in Go

---

## Project Structure
```
hotreload
│
├── cmd
│   └── hotreload
│       └── main.go          # CLI entry point
│
├── internal
│   ├── builder
│   │   └── builder.go       # Runs build command
│   │
│   ├── runner
│   │   └── runner.go        # Manages server process
│   │
│   └── watcher
│       └── watcher.go       # File watching logic
│
├── testserver
│   └── main.go              # Demo HTTP server
│
├── bin                      # Built server binary
│
├── Makefile                 # Build and run commands
├── run.bat                  # Windows run script
├── go.mod
└── README.md
```

---

## Installation

### 1. Clone the Repository
```bash
git clone https://github.com/<your-username>/hotreload.git
cd hotreload
```

### 2. Install Dependencies
```bash
go mod tidy
```

---

## Running the Demo

Run the hot reload engine with the following command:
```bash
go run ./cmd/hotreload \
  --root ./testserver \
  --build "go build -o bin/server.exe ./testserver" \
  --exec "bin/server.exe"
```

---

## How It Works

HotReload operates in four main stages:

**1. Watch Files**
The watcher monitors the project directory using `fsnotify`.

**2. Detect Changes**
Whenever a file is modified, a file change event is triggered.

**3. Rebuild Project**
The build command provided by the user is executed.

**4. Restart Server**
The running server process is terminated and restarted automatically.

---

## Example Workflow

Start the tool:
```bash
go run ./cmd/hotreload \
  --root ./testserver \
  --build "go build -o bin/server.exe ./testserver" \
  --exec "bin/server.exe"
```

Console output:
```
HotReload starting...
Building project...
Starting server...
Server running on :8080
```

Now modify `testserver/main.go`. After saving the file, HotReload detects the change and rebuilds automatically:
```
File changed
Building project...
Stopping previous server...
Starting server...
```

Refreshing the browser will show the updated server.

---

## Technologies Used

- [Go](https://golang.org/)
- [fsnotify](https://github.com/fsnotify/fsnotify) – File system notifications
- `os/exec` – Process management

---

## Future Improvements

- Debouncing file events to prevent repeated builds
- Ignoring directories like `.git` and `node_modules`
- Automatic detection of newly created folders
- Improved process termination for stubborn processes
- Enhanced cross-platform compatibility

---

## Author

**Dhairya Tiwari**