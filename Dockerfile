FROM golang:1.18beta2


## --------------------------------------
## Authorship
## --------------------------------------

LABEL org.opencontainers.image.authors="sakutz@gmail.com"


## --------------------------------------
## Apt and standard packages
## --------------------------------------

# Update the local apt cache.
RUN apt-get update -y

# Install vim.
RUN apt-get install -y vim


## --------------------------------------
## Golang
## --------------------------------------

# Install graphviz so "go tool pprof" can export images.
RUN apt-get install -y graphviz

# Install the Go debugger.
RUN go install github.com/go-delve/delve/cmd/dlv@latest


## --------------------------------------
## .NET
## --------------------------------------

# Install .NET for Linux.
RUN curl -LO https://dot.net/v1/dotnet-install.sh && \
    bash dotnet-install.sh && \
    rm -f dotnet-install.sh

# Configure the .NET environment variables.
ENV DOTNET_ROOT=/root/.dotnet
ENV PATH="${PATH}:${DOTNET_ROOT}:${DOTNET_ROOT}/tools"

# Install the .NET debugger.
RUN apt-get install -y lldb libicu-dev
RUN dotnet tool install -g dotnet-sos && \
    dotnet-sos install


## --------------------------------------
## Java
## --------------------------------------

# Install the OpenJDK.
RUN apt-get install -y openjdk-11-jdk

# Install expect to interact with the Java debugger (jdb).
RUN apt-get install -y expect

# Configure the Java environment variables.
ARG TARGETARCH
ENV JAVA_HOME="/usr/lib/jvm/java-11-openjdk-${TARGETARCH}"


## --------------------------------------
## Main
## --------------------------------------

# Create the working directory.
RUN mkdir -p /go-generics-the-hard-way
WORKDIR /go-generics-the-hard-way/

# Copy the current repo into the working directory.
COPY . /go-generics-the-hard-way/
