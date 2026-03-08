HotReload – Go CLI Hot Reload Engine

A lightweight CLI tool written in Go that automatically rebuilds and restarts a server when source code changes.
This tool eliminates the need for developers to manually stop, rebuild, and restart servers during development.

The tool watches a project directory and triggers a rebuild whenever a file is modified.

Problem

During development, engineers often follow this workflow:

1. Modify code
2. Stop server
3. Rebuild project
4. Start server again

This process is repetitive and slows down development.

HotReload automates this workflow by watching for file changes and restarting the server automatically.

Features

• Watches project directory recursively
• Automatically rebuilds the project when code changes
• Automatically restarts the server
• Streams server logs in real time
• Performs initial build when the tool starts
• Supports nested project folders
• Lightweight CLI tool written in Go

Project Structure
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
Installation

Clone the repository:

git clone https://github.com/<your-username>/hotreload.git
cd hotreload

Install dependencies:

go mod tidy
Running the Demo

Run the hot reload tool:

go run ./cmd/hotreload \
--root ./testserver \
--build "go build -o bin/server.exe ./testserver" \
--exec "bin/server.exe"
How It Works

The tool works in four steps:

Watch Files

The watcher monitors the project directory using fsnotify.

Detect Change

When a file is modified, a change event is triggered.

Rebuild Project

The build command provided by the user is executed.

Restart Server

The running server process is terminated and restarted.

Example Workflow

Start the tool:

go run ./cmd/hotreload --root ./testserver --build "go build -o bin/server.exe ./testserver" --exec "bin/server.exe"

Output:

HotReload starting...
Building project...
Starting server...
Server running on :8080

Now modify:

testserver/main.go

Save the file.

Output:

File changed
Building project...
Stopping previous server...
Starting server...

Refresh the browser and the updated server will be running.

Demo Server

A sample HTTP server is included in the repository for demonstration.

Open:

http://localhost:8080

Edit the message in testserver/main.go and save the file to see hot reload in action.

Technologies Used

Go

fsnotify (file system notifications)

os/exec for process management

Future Improvements

Potential improvements include:

• Debouncing file events to avoid repeated builds
• Ignoring unnecessary directories like .git and node_modules
• Detecting newly created folders automatically
• Better process termination for stubborn processes
• Cross-platform support improvements

Demo Video

A Loom video demonstrating the architecture and functionality of the tool will be provided as part of the assignment submission.

Author

Dhairya Tiwari