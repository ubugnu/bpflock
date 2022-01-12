# bpflock - Lock Linux machines

`bpflock` - eBPF driven security for locking and auditing Linux machines.

Note: bpflock is currently in **experimental stage**, it may break, security semantics may change, some BPF programs will be updated to use [Cilium ebpf library](https://github.com/cilium/ebpf/).

## Sections

* [1. Introduction](https://github.com/linux-lock/bpflock#1-introduction)
* [2. Functionality Overview](https://github.com/linux-lock/bpflock#2-functionality-overview)
  - [2.1 Security features](https://github.com/linux-lock/bpflock#21-security-features)
  - [2.2 Semantics](https://github.com/linux-lock/bpflock#12-semantics)
* [3. Deployment](https://github.com/linux-lock/bpflock#2-deployment)
* [4. Documentation](https://github.com/linux-lock/bpflock#3-documentation)
* [5. Build](https://github.com/linux-lock/bpflock#3-build)

## 1. Introduction

bpflock combines multiple bpf programs to strength Linux security. By restricting access to a various range of Linux features, bpflock is able to reduce the attack surface and block some well known attack techniques.

Only programs like container managers, systemd and other containers/programs that run in the host [pid namespace](https://man7.org/linux/man-pages/man7/namespaces.7.html) may be able to access those features, containers that run on their own namespace will be restricted. If bpflock bpf programs run under a deny permission model then all programs/containers will be denied access even privileged ones.

bpflock protects Linux machines using a system wide approach taking advantage of [Linux Security Modules + BPF](https://www.kernel.org/doc/html/latest/bpf/bpf_lsm.html). The permission model will be augmented soon to include per cgroupv2 filtering.

Architecture and Security design notes:
- bpflock is not a mandatory access control labeling solution, and it does not intent to replace [AppArmor](https://apparmor.net/), [SELinux](https://github.com/SELinuxProject/selinux), and other MAC solutions.
- bpflock offers multiple small bpf programs that can be reused in multiple contexts from Cloud Native deployments to Linux IoT devices.
- bpflock is able to restrict root to access certain Linux features, however it does not protect against evil root users that can disable it.

## 2. Functionality Overview

### 2.1 Security features

bpflock bpf programs offer multiple security protections that can be classified as:

* [Hardware Addition Attacks](https://github.com/linux-lock/bpflock/tree/main/docs/hardware-additions.md)
  - [USB Additions Protection](https://github.com/linux-lock/bpflock/tree/main/docs/hardware-additions.md#1-usb-additions-protection)

* [Memory Protections](https://github.com/linux-lock/bpflock/tree/main/docs/memory-protections.md)
  - [Kernel Image Lockdown](https://github.com/linux-lock/bpflock/tree/main/docs/memory-protections.md#1-kernel-image-lockdown)
  - [Kernel Modules Protection](https://github.com/linux-lock/bpflock/tree/main/docs/memory-protections.md#2-kernel-modules-protections)
  - [BPF Protection](https://github.com/linux-lock/bpflock/tree/main/docs/memory-protections.md#3-bpf-protection)
  - [Execution of Memory ELF binaries](https://github.com/linux-lock/bpflock/tree/main/docs/memory-protections.md#4-execution-of-memory-elf-binaries)

* [Filesystem Protections](https://github.com/linux-lock/bpflock/tree/main/docs/filesystem-protections.md)

  - Read-only root filesystem protection
  - sysfs protection

* [Linux Namespaces Protections](https://github.com/linux-lock/bpflock#34-namespaces-protections)

* Network protections

  - bpflock may include in future a simple network protection that can be used in single machine workload or Linux-IoT, but will not include a Cloud Native protection. [Cilium](https://github.com/cilium/cilium) and other kubernetes CNI related solutions are by far better.

### 2.2 Semantics

bpflock tries to keep the security semantics simple without introducing complex policies. It uses a simple `permission` model that takes the following values:

* `permission`: each bpf program supports three different permission models.
  - `allow|none` : access is allowed. It can be used to log security events.
  - `restrict` : access is restricted. Only processes that are in the initial pid namespace.
  - `deny` : access is denied for all processes.

* `Allowed` or `blocked` operations/commands:

  Depending on the `permission` model a list of allowed or blocked commands can be specified with:
  - `allow` : comma-separated list of allowed operations. Valid under `restrict` and `deny` permissions.
  - `block` : comma-separated list of blocked operations. Valid under `restrict` permission.

For bpf security examples check [bpflock configuration examples](https://github.com/linux-lock/bpflock/tree/main/deploy/configs/)


## 3. Deployment

### 3.1 Prerequisites

bpflock needs the following:

* Linux kernel version >= 5.15 with the following configuration:

  ```code
  CONFIG_BPF_SYSCALL=y
  CONFIG_DEBUG_INFO=y
  CONFIG_DEBUG_INFO_BTF=y
  CONFIG_KPROBES=y
  CONFIG_LSM="...,bpf"
  CONFIG_BPF_LSM=y
  ```

* Obviously a BTF enabled kernel.


### 3.2 Docker deployment

**Note: this is used for current dev/testing, it will be changed soon so configurations are shipped inside image.**

First fetch one of the following bpf security configuration: [bpf
configurations](https://github.com/linux-lock/bpflock/tree/main/deploy/configs/bpf.d) and save it into current
directory.

Then run container using:
```
docker run --name bpflock -it --rm --cgroupns=host --pid=host --privileged -v /sys/kernel/security:/sys/kernel/security -v /sys/fs/bpf:/sys/fs/bpf -v $(pwd)/deny.yaml:/etc/bpflock/bpf.d/deny.yaml linuxlock/bpflock
```

## 4. Documentation

Documentation files can be found [here](https://github.com/linux-lock/bpflock/tree/main/docs/).

## 5. Build

bpflock uses [docker BuildKit](https://docs.docker.com/develop/develop-images/build_enhancements/) to build and [Golang](https://go.dev/doc/install) for running tests.

To build it just run:

```bash
make
```

Bpf programs are built using libbpf. The docker image used is Ubuntu.

## Credits

bpflock uses lot of resources including source code from the [Cilium](https://github.com/cilium/cilium) and
[bcc](https://github.com/iovisor/bcc) projects.

## License

The bpflock user space components are licensed under the [Apache License, Version 2.0](https://github.com/linux-lock/bpflock/blob/main/LICENSE). The BPF code where it is noted is licensed under the [General Public License, Version 2.0](https://github.com/linux-lock/bpflock/blob/main/src/COPYING).
