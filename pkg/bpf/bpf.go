// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Djalal Harouni
// Copyright 2018 Authors of Cilium

package bpf

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/linux-lock/bpflock/pkg/defaults"

	"github.com/sirupsen/logrus"
)

var (
	bpfProgramsPath = filepath.Join(defaults.ProgramLibraryPath, "bpf")
	bpftool         = filepath.Join(defaults.ProgramLibraryPath, "bpftool")
)

// #rm -fr /sys/fs/bpf/bpflock/$pinnedProg
func bpftoolUnload(pinnedProg string) {
	bpffs := filepath.Join(MapPrefixPath(), pinnedProg)

	log.Infof("removing bpf-program=%s", pinnedProg)
	os.RemoveAll(bpffs)
}

// #bpftool prog show name progName
func bpftoolGetProgID(progName string) (string, error) {
	args := []string{"prog", "show", "name", progName}
	log.WithFields(logrus.Fields{
		"bpftool": bpftool,
		"args":    args,
	}).Debug("GetProgID:")
	output, err := exec.Command(bpftool, args...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to show %s: %s: %s", progName, err, output)
	}

	// Scrap the prog_id out of the bpftool output after libbpf is dual licensed
	// we will use programatic API.
	s := strings.Fields(string(output))
	if s[0] == "" {
		return "", fmt.Errorf("Failed to find prog %s: %s", progName, err)
	}
	progID := strings.Split(s[0], ":")
	return progID[0], nil
}

// BpfLsmEnable will execute all programs according to configuration
// and corresponding bpf programs will be pinned automatically
func BpfLsmEnable() error {
	return nil
}

// BpfLsmDisable will detach any bpf programs and unloads them.
// All the programs and maps associated with it will be deleted
// from the bpf filesystem.
func BpfLsmDisable() error {
	p := MapPrefixPath()
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return fmt.Errorf("failed to read directory '%s': %s", p, err)
	}

	for _, f := range files {
		if strings.HasPrefix(f.Name(), "..") {
			continue
		}
		if f.IsDir() {
			bpftoolUnload(f.Name())
		}
	}

	return nil
}
