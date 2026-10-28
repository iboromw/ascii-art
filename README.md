#  ASCII-Art

An efficient CLI tool written in **Go** that converts text input into stylized ASCII art banner text using standard 8-line tall font templates (`standard`, `shadow`, `thinkertoy`).

Built strictly with the **Go Standard Library** (zero external dependencies).

---

##  Table of Contents

- [Features](#-features)
- [Architecture & Data Pipeline](#-architecture--data-pipeline)
- [Banner Styles](#-banner-styles)
- [Installation & Setup](#-installation--setup)
- [Usage & Examples](#-usage--examples)
- [Project Structure](#-project-structure)
- [Team & Responsibilities](#-team--responsibilities)
- [Testing & Quality Assurance](#-testing--quality-assurance)
- [License](#-license)

---

##  Features

-  **ASCII Art Rendering**: Transforms standard text strings into 8-line stacked ASCII character matrices.
-  **Newline Support**: Handles literal `\n` sequences and multi-consecutive newlines (e.g., `Hello\n\nWorld`).
-  **Multiple Banners**: Supports standard font styles including `standard.txt`, `shadow.txt`, and `thinkertoy.txt`.
-  **Zero External Dependencies**: Built strictly using Go's built-in standard packages (`os`, `bufio`, `strings`, etc.).
-  **Extensive Test Coverage**: Includes comprehensive unit tests and CLI integration test suites.

---

## Architecture & Data Pipeline

```
┌─────────────────────────┐     ┌─────────────────────────┐     ┌─────────────────────────┐
│     Bilal ZOUHAIDI      │     │    Nicolas Lindberg     │     │  Iboro Monday William   │
│      Input Parser       │ ──> │   Banner & Renderer     │ ──> │     QA & Integration    │
│  os.Args ➔ []string    │     │ []string ➔ ASCII Lines │     │   Validation & Output   │
└─────────────────────────┘     └─────────────────────────┘     └─────────────────────────┘
```

1. **Input Parsing (`input.go`)**: Reads CLI arguments (`os.Args`), handles empty strings, space padding, and converts literal `\n` sequences into logical line slices (`[]string`).
2. **Banner Loading & Rendering (`banner.go`, `renderer.go`)**: Reads standard banner files, maps characters to exact 8-line matrix representations, and horizontally composites characters side-by-side.
3. **Integration & Validation (`main.go`, tests)**: Validates input, coordinates pipeline execution, and outputs clean formatted ASCII text to standard output.

---

## Banner Styles

- **Standard (`standard.txt`)**: Classic ASCII block typography.
- **Shadow (`shadow.txt`)**: Drop-shadow styled 3D ASCII lettering.
- **Thinkertoy (`thinkertoy.txt`)**: Minimalist single-line stroke ASCII font.

---

## Installation & Setup

### Prerequisites

- [Go 1.18+](https://go.dev/dl/) installed on your machine.

### Clone the Repository

```bash
git clone https://github.com/bilalZOUHAIDI/ascii-art.git
cd ascii-art
```

---

##  Usage & Examples

Run the application using `go run`:

### Basic Usage

```bash
go run . "Hello"
```

### Multi-Line Input (`\n`)

```bash
go run . "Hello\nWorld"
```

```bash
go run . "Hello\n\nThere"
```

### Testing Output Formatting

To inspect exact trailing spaces and line breaks:

```bash
go run . "Hello\nThere" | cat -e
```

---

##  Project Structure

```
ascii-art/
├── main.go               # Entry point & CLI argument coordination
├── input.go              # Command-line input parsing & newline handling
├── banner.go             # Banner file loading & character matrix mapping
├── renderer.go           # 8-line stacked character rendering engine
├── input_test.go         # Unit tests for input parsing
├── renderer_test.go      # Unit tests for banner & rendering logic
├── integration_test.go   # End-to-end integration test suite
├── standard.txt          # Standard ASCII banner font matrix
├── shadow.txt            # Shadow ASCII banner font matrix
├── thinkertoy.txt        # Thinkertoy ASCII banner font matrix
└── README.md             # Project documentation
```

---

##  Team & Responsibilities

| Member | Handle | Role & Main Responsibilities |
| :--- | :--- | :--- |
| **Bilal ZOUHAIDI** | `@bzouhaidi` | Input handling, argument parsing, newline logic (`\n`), program entry point |
| **Nicolas Lindberg** | `@nlindber` | Banner file loading, character-to-ASCII matrix mapping, horizontal renderer |
| **Iboro Monday William** | `@iwilliam` | Integration testing, edge cases validation, formatting (`gofmt`/`govet`), QA |

---

##  Testing & Quality Assurance

Run unit and integration tests with:

```bash
go test -v ./...
```

Enforce formatting & code quality guidelines:

```bash
gofmt -w .
go vet ./...
```

---
