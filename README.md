# foldpurge

`foldpurge` is a command-line utility for scanning folders to calculate their size and deleting all items within specified folders. 

Powered by `cobra` + `fatih/color`.

## features

- scan folders to calculate their size
- delete all items under specified folders
- supports custom configuration files for specifying paths
- quiet mode to suppress all output except errors

## installation

to install `foldpurge`, clone the repository and build the binary using `go build`, or get it from the releases.

```sh
git clone https://github.com/yourusername/foldpurge.git
cd foldpurge
go build
```
### commands and flags

| command | description | flags | example |
|---------|-------------|-------|---------|
| `scan`  | scans folders and calculates their size | `--config, -c` : path to config file <br> `--path, -p` : manually specify folders to process <br> `--quiet, -q` : suppresses all output except stderr | `./foldpurge scan -p /path/to/folder1 -p /path/to/folder2` |
| `nuke`  | deletes all items under specified folders | `--config, -c` : path to config file <br> `--path, -p` : manually specify folders to process <br> `--quiet, -q` : suppresses all output except stderr | `./foldpurge nuke -p /path/to/folder1 -p /path/to/folder2` |

### global flags

| flag         | description                         | example |
|--------------|-------------------------------------|---------|
| `--config, -c` | path to config file                | `./foldpurge scan --config /path/to/config.yaml` |
| `--path, -p`   | manually specify folders to process | `./foldpurge scan --path /path/to/folder1 --path /path/to/folder2` |
| `--quiet, -q`  | suppresses all output except stderr | `./foldpurge scan --quiet` |

## examples

to scan folders and calculate their size:

```sh
./foldpurge scan --path /path/to/folder1 --path /path/to/folder2
```

to delete all items under specified folders:

```sh
./foldpurge nuke --path /path/to/folder1 --path /path/to/folder2
```

## license

see the [LICENSE](LICENSE) file for details.

## contributing

contributions are welcome! please open an issue or submit a pull request on GitHub.
