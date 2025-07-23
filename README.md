

# WATcloud CLI


## Features


## Setup & Installation

**Requirements:** Go 1.22+ (and optionally Docker for daemon status)

Clone the repository:
```sh
git clone https://github.com/WATonomous/watcloud-cli.git
cd watcloud-cli
```

Build:
```sh
go build -o watcloud ./cmd/watcloud
```

Run:
```sh
./watcloud status
./watcloud quota list
./watcloud daemon status
```

## Project Structure

- `cmd/` - CLI entrypoints
- `internal/` - Command implementations

## Commands

### watcloud quota

| Subcommand | Description |
|------------|--------------------------------------------------|
| list       | Lists all quota usage (disk, memory, CPU).       |
| disk       | Shows your disk usage percentage and free space. |
| cpu        | Displays CPU usage percentage.                   |
| memory     | Shows memory usage statistics.                   |

### watcloud daemon

| Subcommand | Description                                      |
|------------|--------------------------------------------------|
| status     | Lists all non-interactive background user processes (daemons). |

---

For help and usage examples, run:
```
./watcloud -h
./watcloud quota -h
./watcloud <command> -h
```
