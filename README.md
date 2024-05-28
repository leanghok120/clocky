# clocky

clocky is a **simple** CLI clock written in Go. clocky is a project inspired by [tty-clock](https://github.com/xorg62/tty-clock).

## Installation

### Go

```sh
go install github.com/leanghok120/clocky@latest
```

From source:

```sh
git clone https://github.com/leanghok120/clocky.git
cd clocky/cmd/clocky/
go build
go install
```

## Usage

### Help

```sh
clocky -h
```

### 24 hour format

By default clocky uses 24 hour format

```sh
clocky
```

### 12 hour format

You can also use 12 hour format

```sh
clocky -f 12
```

## Contributing

Contributions are welcome! If you have any ideas or improvements, feel free to open an issue or submit a pull request.

## Acknowledgements

- Thanks to [figlet4go](https://github.com/mbndr/figlet4go) for providing ASCII art rendering functionality.
