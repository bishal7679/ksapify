![Untitled-2022-11-18-1214](https://github.com/bishal7679/ksapify/assets/70086051/7475533e-fbdb-4514-b056-e1f3bb5e663d)



**A Multi-Featured Lightweight Kubernetes command-line tool which can interact with k8s API-SERVER**

[![Testing Build process](https://github.com/kubesimplify/ksctl/actions/workflows/testBuilder.yaml/badge.svg)](https://github.com/kubesimplify/ksctl/actions/workflows/testBuilder.yaml) [![goreleaser](https://github.com/kubesimplify/ksctl/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/kubesimplify/ksctl/actions/workflows/goreleaser.yml)


# Project Scope 🧪
There is always need to install one kubernetes CLI to interact with k8s api-server in order to get all the objects/resources. It is built on Golang and utilizes the `client-go` library to interact with Kubernetes Cluster. The aim is to build a plug & play k8s CLI that can be used to interact with any Kubernetes cluster. It can also be considered as a lightweight solution 😄

# **Project Features** :exploding_head:
 🔹Get details about any resource in the cluster <br>
 🔹Create new resources in the cluster <br>
 🔹Delete resources in the cluster <br>
 🔹Run CLI commands with various Flags <br>
 🔹Switching any namespace and also back to previous namespce <br>
 🔹Get live events from the cluster <br>
 🔹More features coming soon... :construction:

# Demo Screenshot
![Screenshot from 2023-05-15 13-19-02](https://github.com/bishal7679/ksapify/assets/70086051/482af7de-8720-4f30-8926-bc5899c9b559)


# Supported Platforms

Platform | Status
--|--
Linux | :heavy_check_mark:
macOS | :heavy_check_mark:
Windows | TODO

# Installation

## Linux and MacOS

```bash
bash <(curl -s https://raw.githubusercontent.com/bishal7679/ksapify/main/install.sh)
```

# Uninstall?

## Linux & MacOs

```bash
bash <(curl -s https://raw.githubusercontent.com/bishal7679/ksapify/main/uninstall.sh)
```
# Build from src
## Linux
### Install


```zsh
make install_linux
```

## macOS
### Install

```zsh
# macOS on M1
make install_macos
```

### Uninstall
```zsh
make uninstall
```

# Contribution Guidelines
Please refer to our [contribution guide](CONTRIBUTING.md) if you wish to contribute to the project :smile:

# ❤️ Show your support
Give a ⭐ if this project helped you, Happy learning!
