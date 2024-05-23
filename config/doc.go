/*
Package config provides functionality to read and manage configuration settings for the application.
It uses the Viper library to handle configuration files and environment variables.

Usage:

To use this package, import it and call the GetConfig function with the appropriate parameters.

Example:

	package main

	import (
	    "fmt"
	    "path/to/your/project/config"
	)

	func main() {
	    conf, err := config.GetConfig("config", ".", "yaml")
	    if err != nil {
	        fmt.Println("Error:", err)
	        return
	    }

	    fmt.Println("Database Address:", conf.DatabaseAddress())
	}
*/
package config
