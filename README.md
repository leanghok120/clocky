# Clocky

Clocky is a **simple** command-line timer application written in Go. It displays the current time in an ASCII art format that updates every second.

## Installation

### Go

```sh
go install github.com/leanghok120/clocky@latest
```

From source:

```sh
git clone https://github.com/leanghok120/clocky.git
cd clocky
go install
```

## Usage

### 24 hour format

By default clocky uses 24 hour format

```sh
clocky
```

### 12 hour format

You can also use 12 hour format

```sh
clocky 12
```

## Contributing

Contributions are welcome! If you have any ideas or improvements, feel free to open an issue or submit a pull request.

## Acknowledgements

- Thanks to [figlet4go](https://github.com/mbndr/figlet4go) for providing ASCII art rendering functionality.
