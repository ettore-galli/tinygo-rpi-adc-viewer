# TINYGO

## Tinygo environment fastpath setup

### 1. Install go

Install go from official go lang website

<https://go.dev>

### 2. Download tinygo from official website and place it locally

Note: we suggest the "direct download" method, which is far clearer than doing it via homebrew.

Download the latest tinygo distribution from the site

<https://tinygo.org/getting-started/install/macos/>

At the time of writing these notes the direct link was the following:

<https://github.com/tinygo-org/tinygo/releases/download/v0.29.0/tinygo0.29.0.darwin-amd64.tar.gz>

Uncompress the tinygo directory and place it under home directory:

```bash
~/tinygo
```

### 3. Configure your shell to add tinygo to your path

Add the following configuration to ```~/.zshrc``` (or any other shell configuration file) in order to see tinygo (or do the following by hand each time)

```shell
#
# Tiny Go Setup
#
TINYGO_PATH=~/tinygo/bin
export PATH=$TINYGO_PATH:$PATH
```

## Online documentation

### Video tutorial

<https://youtu.be/B-6GsoEg0Lw?si=KrkhNcieKCxIBY-w>

### Examples

ADC
<https://github.com/soypat/tinygo-arduino-examples/blob/main/lcdscreen_adc/sense.go>

Display
<https://github.com/va1da5/tinygo-pico-ssd1306/blob/main/README.md>

Display Driver
<https://github.com/tinygo-org/drivers/blob/release/ssd1306/ssd1306.go>

## Project setup

### 1. Install and configure TinyGo extension

[tinygo.vscode-tinygo]

### 2. Select the target platform

Thois step is needed for the VSCode syntax check and code cpmpletion to "see" the target libraries

In VSCode:

- Press CMD + Shift + P to enter command mode

- Type "tinygo" to look for extension commands

- Choose TinyGo target

  - Then select "pico" for Raspberry Pi Pico microcontroller.

### 3. How to install third party library for display

#### 3/a Install third party libraries

```bash
# (As you would do in go lang)
go get tinygo.org/x/drivers  
```

### 3/b Autofix and optimize module installation

Perform after installing all needed external modules

```bash
# (As you would do in go lang)
go mod tidy  
```
