# Docker

Docker is the preferred approach to running the examples from this repository.

* [**Build the image**](#build-the-image): build the Docker image locally
* [**Use the image**](#use-the-image): use the Docker image to run the examples in this repository

## Build the image

Build the Docker image by executing the following commands:

1. Clone this repository locally:

    ```bash
    git clone https://github.com/akutz/go-generics-the-hard-way
    ```

1. Change directories to the root of the cloned repository:

    ```
    cd go-generics-the-hard-way
    ```

1. Build the image:

    ```bash
    docker build -t go-generics-the-hard-way .
    ```

And voilá, you have the image!

---

:star: **Please note** the image is also available remotely and can obtained by executing the following commands:

```bash
docker pull akutz/go-generics-the-hard-way && \
docker tag akutz/go-generics-the-hard-way go-generics-the-hard-way
```

---

## Use the image

The image may be used one of three ways:

* [**Self-contained**](#self-contained): run the examples from the image's filesystem
* [**Mount repository**](#mount-repository): run the examples from your local filesystem using the image
* [**VSCode Dev Container**](#dev-container): run the examples in VS Code using the `Remote - Containers` extension

### Self-contained

This method does not require cloning this repository locally since the Docker image includes the contents of this repository. For example, the following command will run the boxing benchmark using the examples sources from inside the image:

```bash
docker run -it --rm go-generics-the-hard-way \
  go test -tags benchmarks -run Boxing -bench Boxing -benchmem -count 5 -v ./06-benchmarks/
```

### Mount repository

Alternatively, it is also possible to run the examples cloned on your local filesystem using the Docker image. This has two benefits:

* The examples will be up-to-date
* It is easier to capture file output from the examples such as profiles

For example, the following command will:

* Run the boxing benchmark using the examples from the cloned repository
* Produce a memory profile
* Produce an SVG image from the memory profile

```bash
docker run -it --rm -v "$(pwd):/go-generics-the-hard-way" \
  go-generics-the-hard-way bash -c ' \
  go test -tags benchmarks -run Boxing -bench Boxing -benchmem -count 5 -memprofile memprofile.out -v ./06-benchmarks/ && \
  go tool pprof -svg memprofile.out'
```

At the end there will be a file in the current, local directory named `profile001.svg` that visualizes the memory profile produced from the benchmark.

### Dev Container

Alternatively, it is also possible to run the examples through opening the cloned reposity in a VS Code devcontainer:

Note: This assumes you are using VS Code as your IDE

1. To get started, install the extension `Remote - Containers` by Microsoft (ms-vscode-remote.remote-containers). 
1. Open up your Command Palette and type or select `Remote-Containers: Reopen in Container` command.
  - The extension should be able to pick up the settings described in the .devcontainer/devcontainer.json and build the Dockerfile
1. Open up your terminal and you should be inside a container and at the project root. 
  - A good test is to do 
  ```bash 
  go version
  ```
  If you see `go version go1.18beta2 linux/arm64` you should be good to go.
  
---

Next: [Hello world](../02-hello-world/)
