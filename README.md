# larPING

```text
 /$$                       /$$$$$$$  /$$$$$$ /$$   /$$  /$$$$$$ 
 | $$                      | $$__  $$|_  $$_/| $$$ | $$ /$$__  $$
 | $$  /$$$$$$   /$$$$$$   | $$  \ $$  | $$  | $$$$| $$| $$  \__/
 | $$ |____  $$ /$$__  $$  | $$$$$$$/  | $$  | $$ $$ $$| $$ /$$$$
 | $$  /$$$$$$$| $$  \__/  | $$____/   | $$  | $$  $$$$| $$|_  $$
 | $$ /$$__  $$| $$        | $$        | $$  | $$\  $$$| $$  \ $$
 | $$|  $$$$$$$| $$        | $$       /$$$$$$| $$ \  $$|  $$$$$$/
 |__/ \_______/|__/        |__/      |______/|__/  \__/ \______/ 
```

[![asciicast](https://asciinema.org/a/1264209.svg)](https://asciinema.org/a/1264209)

**larPING** - l33t-speak themed ping CLI tool written in Go, that wraps standard ping in a retro hollywood hacker aesthetic, complete with randomized ASCII art, fake payload injections, and colored typewriter terminal effects.

## Features

* **Functional ICMP Ping:** Uses `prometheus-community/pro-bing` to send real ICMP packets and measure RTT.
* **L33t-speak Terminal Output:** Replaces boring standard ping logs with randomized "hacking" phrases.
* **ASCII Art Injection:** Periodically drops random ASCII logos into the terminal stream.
* **Typewriter Effects & Colors:** Animated startup text and ANSI color-coded event logging.
* **Standard Statistics:** Catching `Ctrl+C` stops the ping and prints standard ICMP statistics (packet loss, count) before exiting.

+ whois WIP

## Installation

Go 1.27.0+

```bash
git clone [https://github.com/clswhre/larPING.git](https://github.com/clswhre/larPING.git)
cd larPING
go mod download
go build -o larPING main.go
```

## Usage

Provide a single, valid IP as an argument.

```bash
./larPING <ip_address>
```

**For example:**
```bash
./larPING 1.1.1.1
```

## Dependencies

* [pro-bing](https://github.com/prometheus-community/pro-bing) (`v0.9.1`)

## License

This project is released under the **WTFPL** (Do What the Fuck You Want to Public License), Version 2. Copyright (C) 2026 clswhre.