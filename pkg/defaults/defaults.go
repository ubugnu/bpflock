// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Djalal Harouni
// Copyright 2016-2020 Authors of Cilium

package defaults

import (
	"time"
)

const (

	// RmBpfOnExit is the default value for option.DeleteBpfOnExit
	RmBpfOnExit = true

	// AgentHealthPort is the default value for option.AgentHealthPort
	AgentHealthPort = 19876

	// GopsPortAgent is the default value for option.GopsPort in the agent
	GopsPortAgent = 19890

	// RuntimePath is the default path to the runtime directory
	RuntimePath = "/var/run/bpflock"

	// RuntimePathRights are the default access rights of the RuntimePath directory
	RuntimePathRights = 0775

	// StateDirRights are the default access rights of the state directory
	StateDirRights = 0770

	//StateDir is the default path for the state directory relative to RuntimePath
	StateDir = "state"

	// TemplatesDir is the default path for the compiled template objects relative to StateDir
	TemplatesDir = "templates"

	// TemplatePath is the default path for a symlink to a template relative to StateDir/<EPID>
	TemplatePath = "template.o"

	// ConfigurationPath
	ConfigurationPath = "/etc/bpflock/"

	// ProgramLibPath is the default path for the bpflock libraries and programs
	ProgramLibPath = "/usr/lib/bpflock"

	// BpfDir is the default path for bpf programs relative to ProgramLibDir
	BpfDir = "bpf"

	// VariablePath is the default path to the bpflock variable state directory
	VariablePath = "/var/lib/bpflock"

	// SockPath is the path to the UNIX domain socket exposing the API to clients locally
	SockPath = RuntimePath + "/bpflock.sock"

	// SockPathEnv is the environment variable to overwrite SockPath
	SockPathEnv = "BPFLOCK_SOCK"

	// MonitorSockPath1_2 is the path to the UNIX domain socket used to
	// distribute BPF and agent events to listeners.
	// This is the 1.2 protocol version.
	MonitorSockPath1_2 = RuntimePath + "/monitor1_2.sock"

	// PidFilePath is the path to the pid file for the agent.
	PidFilePath = RuntimePath + "/bpflock.pid"

	// DefaultMapRoot is the default path where BPFFS should be mounted
	DefaultMapRoot = "/sys/fs/bpf"

	// DefaultMapPrefix
	DefaultMapPrefix = "bpflock"

	// ShortExecTimeout is a short timeout for executing commands.
	ShortExecTimeout = 10 * time.Second

	// ExecTimeout is a timeout for executing commands.
	ExecTimeout = 30 * time.Second

	// ClientConnectTimeout is the time the bpflock agent client is
	// (optionally) waiting before returning an error.
	ClientConnectTimeout = 30 * time.Second

	// StatusCollectorInterval is the interval between a probe invocations
	StatusCollectorInterval = 5 * time.Second

	// StatusCollectorWarningThreshold is the duration after which a probe
	// is declared as stale
	StatusCollectorWarningThreshold = 15 * time.Second

	// StatusCollectorFailureThreshold is the duration after which a probe
	// is considered failed
	StatusCollectorFailureThreshold = 1 * time.Minute

	// EnableIPv4 is the default value for IPv4 enablement
	EnableIPv4 = true

	// EnableIPv6 is the default value for IPv6 enablement
	EnableIPv6 = true

	// BpfProfileAllow is the "allow" "none" or "privileged" profile
	BpfProfileAllow      = "allow"
	BpfProfileNone       = "none"
	BpfProfilePrivileged = "privileged"

	// BpfProfileBaseline is for privileged applications
	BpfProfileBasleine = "baseline"

	// BpfProfileRestricted is deny some privileged operations
	BpfProfileRestricted = "restricted"
)
