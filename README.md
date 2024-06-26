## Usage

```
foldpurge scan -p folder_1 folder_2 folder_n
foldpurge nuke -p folder_1 folder_2 folder_n
```

## Full usage:
```
Usage:
  foldpurge [command]

Available Commands:
  help        Help about any command
  nuke        Delete all items under specified folders
  scan        Scans folders and calculate their size

Flags:
  -c, --config string    Path to config file
  -h, --help             help for foldpurge
  -p, --path strings     Manually specify folders to process
  -q, --quiet            Suppresses all output except stderr

Use "foldpurge [command] --help" for more information about a command.
```
