# Docker

Docker is the preferred approach to running the examples from this repository.

---

:warning: **Please note** while it _is_ possible to use a container to run the examples locally, this method does not enable support for Go 1.18beta1 in VS Code. VS Code _does_ support [remote development via a container](https://code.visualstudio.com/docs/remote/containers), but I have never done so, and unless someone wants to write down how to do it, I will likely not add it here.

---

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

The image may be used one of two ways:

* [**Self-contained**](#self-contained): run the examples from the image's filesystem
* [**Mount repository**](#mount-repository): run the examples from your local filesystem using the image

### Self-contained

This method does not require cloning this repository locally since the Docker image includes the contents of this repository. For example, the following command will run the boxing benchmark using the examples sources from inside the image:

```bash
docker run -it --rm go-generics-the-hard-way \
  go test -bench . -benchmem -count 5 -v ./04-benchmarks/boxing/
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
  go test -bench . -benchmem -count 5 -memprofile memprofile.out -v ./04-benchmarks/boxing/ && \
  go tool pprof -svg memprofile.out'
```

At the end there will be a file in the current, local directory named `profile001.svg` that visualizes the memory profile produced from the benchmark.

---

Next: [Hello world](../02-hello-world/)
