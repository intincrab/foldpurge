package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

var (
	configFile string
	paths      []string
	quiet      bool
	rootCmd    = &cobra.Command{Use: "foldpurge"}
)

var infoColor = color.New(color.FgCyan).SprintFunc()
var successColor = color.New(color.FgGreen).SprintFunc()
var errorColor = color.New(color.FgRed).SprintFunc()
var warnColor = color.New(color.FgHiYellow).SprintFunc() 

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Path to config file")
	rootCmd.PersistentFlags().StringSliceVarP(&paths, "path", "p", []string{}, "Manually specify folders to process")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Suppresses all output except stderr")

	var scanCmd = &cobra.Command{
		Use:   "scan",
		Short: "Scans folders and calculate their size",
		Run:   scanFolders,
	}

	var nukeCmd = &cobra.Command{
		Use:   "nuke",
		Short: "Delete all items under specified folders",
		Run:   nukeFolders,
	}

	rootCmd.AddCommand(scanCmd, nukeCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(errorColor(fmt.Sprintf("Error: %v", err)))
		os.Exit(1)
	}
}

func getValidPaths() ([]string, error) {
	var allPaths []string

	if configFile != "" {
		// Read paths from config file
		// This is a placeholder. need to implement the actual config file reading logic
		allPaths = append(allPaths, "/path/from/config")
	}

	allPaths = append(allPaths, paths...)

	var validPaths []string
	for _, path := range allPaths {
		if _, err := os.Stat(path); err == nil {
			validPaths = append(validPaths, path)
		} else {
			fmt.Println(warnColor(fmt.Sprintf("Warning: Path not found: %s", path)))
		}
	}

	if len(validPaths) == 0 {
		return nil, fmt.Errorf(errorColor("no valid paths found"))
	}

	if !quiet {
		fmt.Println(infoColor(fmt.Sprintf("%d valid entries detected", len(validPaths))))
	}

	return validPaths, nil
}

func scanFolders(cmd *cobra.Command, args []string) {
	validPaths, err := getValidPaths()
	if err != nil {
		fmt.Println(errorColor(fmt.Sprintf("Error: %v", err)))
		return
	}

	for _, path := range validPaths {
		size, err := getDirSize(path)
		if err != nil {
			fmt.Println(errorColor(fmt.Sprintf("Error scanning %s: %v", path, err)))
			continue
		}
		if !quiet {
			fmt.Printf("%s: %s\n", path, infoColor(fmt.Sprintf("%d bytes", size)))
		}
	}
}

func nukeFolders(cmd *cobra.Command, args []string) {
    validPaths, err := getValidPaths()
    if err != nil {
        fmt.Println(err)
        return
    }

	for _, path := range validPaths {
        err := os.RemoveAll(path)
        if err != nil {
            fmt.Println(errorColor(fmt.Sprintf("Error deleting %s: %v", path, err)))
            continue
        }
        
        folderName := filepath.Base(path)
        
        if !quiet {
            fmt.Println(successColor(fmt.Sprintf("✔ %s  has been deleted successfully", folderName)))
        }
    }
}

func getDirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}