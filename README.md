# Okane

<div align="center">
<br>

<img width=100% src="https://github.com/NikhilSharma03/Okane/blob/main/assets/demo.gif"></p>

</div>

<div align="center">
<br>

[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=95)](https://github.com/NikhilSharma03/Okane)
[![Build by Nikhil](https://img.shields.io/badge/Built%20by-Nikhil-Green)](https://github.com/NikhilSharma03/Okane)

</div>

<br>

> Okane is a CLI 💻 application built using Cobra Go🚀 to help manage your expenses 💵

<br>

## Technology Stack

<div align="center">

<img alt="Go" src="https://img.shields.io/badge/go%20-%231572B6.svg?&style=for-the-badge&logo=go&logoColor=white"/> <img alt="gRPC"
src="https://img.shields.io/badge/grpc%20-%231572B6.svg?&style=for-the-badge"/> <img alt="Redis"
src="https://img.shields.io/badge/redis%20-%231572B6.svg?&color=red&style=for-the-badge&logo=redis&logoColor=white"/>

</div>

<br>

## Install CLI

```
go install github.com/NikhilSharma03/Okane/okanecli@latest
```

Now you can access CLI app using `okanecli` in your terminal

<br>

## Setup and Installation

First install `Protocol compiler`

```
brew install protobuf
```

Now, install `gRPC Go` plugin

```
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest \
    google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Then, add GOPATH in `.bashrc` file

```
export PATH=$PATH:$HOME/go/bin
export PATH=$PATH:/usr/local/go/bin
```

To generate proto

```
make generate-proto
```

Now, Add `Environment Variables` by creating a new `.env` file in `root` folder and add the contents following `.env.example`

Once you have added correct credentials, run the server using

```
make run-server
```

If using `docker compose`, then first build the image

```
make compose-dev-build
```

Now start the dev server

```
make compose-dev-up
```

To access logs

```
make compose-dev-logs
```

To shut down the server

```
make compose-dev-down
```

Now the server is running at `localhost:8000`

Now, Lets build the `okane cli` app

```
make build-cli
```

Now you can use the app

```
./okane_cli
```

# License

<div align="center">
<br>

<img width=35% src="https://media0.giphy.com/media/3ornjXbo3cjqh2BIyY/200.gif"></p>

<br>
</div>
