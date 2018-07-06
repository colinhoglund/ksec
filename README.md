# ksec

A command line tool that simplifies the management of Kubernetes Secrets.

## Installation

Compiled binaries can be found in the [GitHub releases](https://github.com/colinhoglund/ksec/releases).

Install compiled binary as a Helm plugin (requires [Helm](https://docs.helm.sh/using_helm/#installing-helm)).

    helm plugin install https://github.com/colinhoglund/ksec

Install from source (requires [golang](https://golang.org/doc/install#install)).

    go get github.com/colinhoglund/ksec

## Usage
```
A tool for managing Kubernetes Secret data

Usage:
  ksec [command]

Available Commands:
  completion  Generate command completion scripts
  create      Create a Secret
  delete      Delete a Secret
  get         Get values from a Secret
  help        Help about any command
  list        List all secrets in a namespace
  pull        Pull values from a Secret into a .env file
  push        Push values from a .env file into a Secret
  set         Set values in a Secret
  unset       Unset values in a Secret

Flags:
      --config string      config file (Default: $HOME/.ksec.yaml)
  -h, --help               help for ksec
  -n, --namespace string   Operate in a specific NAMESPACE (Default: current kubeconfig namespace)
      --version            version for ksec

Use "ksec [command] --help" for more information about a command.
```
