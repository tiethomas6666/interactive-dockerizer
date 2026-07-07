# 🐳 Interactive Dockerizer

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A fast, interactive Command Line Interface (CLI) tool built with Go that instantly generates production-ready `Dockerfile` and `docker-compose.yml` configurations tailored to your project's specific stack.

## 💡 Why Use It?

Setting up Docker environments from scratch involves writing boilerplate code and remembering complex configurations. **Interactive Dockerizer** streamlines this process. By answering a few simple interactive prompts in your terminal, it dynamically scaffolds optimized container environments for your applications, databases, and caches.

* **Zero Dependencies:** Built with Go and `go:embed`, the templates are compiled directly into a single, lightning-fast binary. No Python, Node.js, or external files required to run the generator.
* **Best Practices:** Generates multi-stage builds for Go, optimized base images for Node/Python, and secure default configurations for databases.

## 🛠️ Supported Stacks

* **Languages & Frameworks:** Go, Node.js, Python, Custom Base Images
* **Databases:** MySQL, PostgreSQL, MongoDB, None
* **Caching Layers:** Redis, None

## 🚀 Installation

### Option A: Build from source (Recommended for developers)
Ensure you have Go 1.22+ installed.
```bash
git clone [https://github.com/your-username/interactive-dockerizer.git](https://github.com/your-username/interactive-dockerizer.git)
cd interactive-dockerizer
go build -o my-dockerizer main.go

# Optional: Move to your bin directory for global access
sudo mv my-dockerizer /usr/local/bin/
