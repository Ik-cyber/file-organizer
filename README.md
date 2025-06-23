# 📂 File Organizer

A simple CLI tool written in Go to organize files into folders based on their file extensions.

---

## 📦 Project Structure

```text
file-organizer/
├── cmd/
│   └── file-organizer/
│       └── main.go        # Only entry point
├── internal/
│   ├── organizer/
│   │   ├── organizer.go   # File moving logic
│   │   └── map.go         # File extension map
│   ├── utils/
│   │   └── fileutils.go   # Helper functions (optional, like path handling)
├── testfolder/            # Your test directory
├── go.mod
└── README.md
```

---

## 🚀 Features

- Automatically organizes files into folders based on their extensions.
- Accepts a **target directory as an argument.**
- Defaults to the **current working directory** if no argument is provided.
- Cross-platform support (Linux, Windows, Mac).

---

## 🛠️ Installation

### ✅ On Linux / MacOS

#### 1. Clone the project

```bash
git clone https://github.com/Ik-cyber/file-organizer.git
cd file-organizer
```

#### 2. Build the project

```bash
go build -o bin/file-organizer ./cmd/file-organizer
```

#### 3. (optional) Add to Path for global use

```bash
export PATH=$PATH:$(pwd)/bin
```

### ✅ On Windows

#### 1. Clone the project

```bash
git clone https://github.com/Ik-cyber/file-organizer.git
cd file-organizer
```

#### 2. Build the project

```bash
go build -o bin/file-organizer.exe ./cmd/file-organizer
```

# ⚙️ Usage

## Run with a Target Directory:

```bash
file-organizer ./testfolder
```

# ✔️ This will organize files inside the testfolder.

## Run without Arguments (Organizes Current Directory):

```bash
./bin/file-organizer ./testfolder # Linux/Mac
.\bin\file-organizer.exe testfolder # Windows
```

# 📝 Notes

- You can pass absolute or relative paths.
- Supports nested directories.
- If you want to test without moving real files, create a sandbox folder like `testfolder`.

# 💡 Future Improvements

- Add verbose mode flag `-v` for logging detailed actions.
- Add dry-run mode to preview actions without moving files.
- Add support for ignoring specific file types.

# 🙏 Credits

Project built by [Ik-cyber](https://github.com/Ik-cyber).
