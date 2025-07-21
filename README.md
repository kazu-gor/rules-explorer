# Rules Explorer

A fast, interactive terminal UI application for exploring and searching configuration files in your project. Built specifically for navigating Cursor rules, Claude configurations, and related documentation files.

## Features

- 🔍 **Real-time Search**: Search across file paths and content simultaneously
- 📁 **Smart File Discovery**: Automatically finds relevant configuration files
- 👀 **Live Preview**: View file contents in a dedicated preview pane
- ⌨️ **Keyboard Navigation**: Efficient terminal-based interface
- 🚀 **Lightweight**: Fast startup and responsive performance
- 🎯 **Focused Scope**: Targets specific file types for better organization

## Supported File Types

Rules Explorer automatically discovers and indexes:

- **Cursor Rules**: `.cursor/rules/*.mdc` files
- **Claude Configuration**: `CLAUDE.md` files (anywhere in the project)
- **Claude Settings**: `.claude/*` files (direct children only)

## Installation

### Prerequisites

- Go 1.21 or later

### Build from Source

```bash
# Clone the repository
git clone https://github.com/your-username/rules-explorer.git
cd rules-explorer

# Build the application
go build -o rules-explorer ./cmd/rules-explorer

# Or install directly
go install ./cmd/rules-explorer
```

## Usage

### Basic Usage

```bash
$ sudo mv rules-explorer /usr/local/bin/rules-explorer
$ chmod +x /usr/local/bin/rules-explorer
$ rules-explorer
```

```bash
# Run from your project directory
./rules-explorer

# Or run directly with go
go run ./cmd/rules-explorer
```

### Keyboard Shortcuts

| Key | Action |
|-----|--------|
| `Tab` | Switch focus between search and file list |
| `↑` / `↓` | Navigate file list |
| `Ctrl+P` / `Ctrl+N` | Alternative navigation (vim-style) |
| `Enter` | Open selected file in preview |
| `Ctrl+C` / `Escape` | Exit application |

### Workflow

1. **Launch** the application in your project root
2. **Search** by typing in the search field (searches both filenames and content)
3. **Navigate** the filtered results using arrow keys or vim-style shortcuts  
4. **Preview** file contents in the right pane as you navigate
5. **Switch focus** with Tab to move between search and file list

## Project Structure

```
rules-explorer/
├── cmd/rules-explorer/main.go    # Application entry point
├── internal/
│   ├── ui/                       # UI components and application logic
│   │   ├── app.go               # Main application struct and setup
│   │   ├── keybinds.go          # Keyboard handling
│   │   └── theme/               # UI theming system
│   │       └── theme.go         # Color schemes and styling
│   └── file/                     # File operations
│       └── explorer.go          # File discovery and filtering
├── go.mod
├── go.sum
├── CLAUDE.md                     # Development instructions
└── README.md
```

## Architecture

The application follows clean architecture principles:

### File Layer (`internal/file`)
- **`Item`**: Struct containing file path and content
- **`Explorer`**: Handles file discovery, loading, and filtering
- Efficiently walks directory trees and loads only matching files

### UI Layer (`internal/ui`)  
- **`App`**: Main application state management
- **Keybinding handlers**: Centralized keyboard input processing
- **Component setup**: Search input, file list, and preview pane coordination

### Key Design Decisions
- **Memory-efficient**: Files are loaded once at startup
- **Responsive search**: Real-time filtering without re-reading files
- **Separation of concerns**: Clear boundaries between file operations and UI

## Development

### Setup

```bash
# Install dependencies
go mod download

# Run tests
go test ./...

# Run linter (requires staticcheck)
staticcheck ./...
```

### Building

```bash
# Development build
go build ./cmd/rules-explorer

# Production build with optimizations
go build -ldflags="-s -w" -o rules-explorer ./cmd/rules-explorer
```

### Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes following Go conventions
4. Add tests for new functionality
5. Run linter and tests (`staticcheck ./...` and `go test ./...`)
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

## Requirements

- **Go**: 1.21 or later
- **Terminal**: Any terminal with basic color support
- **OS**: macOS, Linux, or Windows

## Dependencies

- [`github.com/rivo/tview`](https://github.com/rivo/tview) - Terminal UI library
- [`github.com/gdamore/tcell`](https://github.com/gdamore/tcell) - Low-level terminal interface

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by the [`pst`](https://github.com/skanehira/pst) project architecture
- Built with the excellent [`tview`](https://github.com/rivo/tview) library
- Designed for developers using Cursor and Claude AI tools
