// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/hello-world
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"flag"
	"fmt"
	"os"
)

// Updated via Makefile builds. Setting placeholder value here so that
// something resembling a version string will be provided for non-Makefile
// builds.
var version = "x.y.z"

// Application metadata for easy reference.
const (
	myAppName string = "hw"
	myAppURL  string = "https://github.com/atc0005/" + myAppName
)

// Config represents configuration details for this application.
type Config struct {
	URL     string
	Verbose bool
	Version bool
}

// NewConfig processes flag values and returns an application configuration.
func NewConfig() *Config {
	var cfg Config
	setupFlags(&cfg)

	return &cfg
}

// Version emits application name, version and repo location.
func Version() string {
	return fmt.Sprintf("%s %s (%s)", myAppName, version, myAppURL)
}

// usage is a custom override for the default Help text provided by the flag
// package. Here we prepend some additional metadata to the existing output.
func usage() {
	_, _ = fmt.Fprintln(flag.CommandLine.Output(), "\n"+Version()+"\n")
	_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

// setupFlags processes given flag values.
func setupFlags(c *Config) {
	flag.BoolVar(&c.Version, "version", false, "Display version information and immediately exit")
	flag.Usage = usage
	flag.Parse()
}
