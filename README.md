<h1 align="center">InfraSights</h1>

<div align="center">

![STATUS](https://img.shields.io/badge/status-active-brightgreen?style=for-the-badge)
<!--- ![LICENSE](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge) -->

</div>

---

## 🔎 About

The Infrasights CLI Tool is a command-line utility designed to provide developers with essential information about their development environment. It gathers various details such as Git status, Kubernetes configuration, Docker images, and more, helping developers streamline their workflows and troubleshoot issues effectively.

## 🏁 Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before running the InfraSights CLI, make sure you have the following prerequisites installed:

**For usage**
- [Git](https://git-scm.com/downloads)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Docker](https://docs.docker.com/engine/install/)

**For development (additional)**
- [Go 1.21](https://go.dev/doc/install) or higher
- [Just](https://github.com/casey/just) command runner

### Installing

To get started with this project, clone the repository to your local machine and install the required dependencies.

```bash
git clone https://github.com/jgfranco17/infrasights.git
cd infrasights
just tidy
```

## 🚀 Usage

The tool will display information about your development environment, including Git status, Kubernetes configuration, Docker images, and more. Infrasights does not require any configuration. However, you can customize certain aspects by editing the source code or providing command-line arguments. Use the basic `help` to get an overview of commands available.

```shell
infrasights --help
```

## 🔧 Testing

### Running unittest suite

In order to run diagnostics and unittests, first install the testing dependencies. This will allow you to utilize the full capacity of the test modules we have built. To run the full test suite, run the Justfile command as follows:

```bash
just test
```

## ✒️ Authors

- [Chino Franco](https://github.com/jgfranco17)
