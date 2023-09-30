# tinygo-rpi-adc-viewer

## About the project

ADC value viewer on display

## Project setup

## Daily development routine commands

Flash (build and flash) project onto device.

Run within the same directory where main.go file is placed.

Please note the dot [.] at the end, meaning flash "everything"

```bash
tinygo flash -target=pico .
```

Simply build without flashing

```bash
tinygo build -target=pico .
```

Connect to device and read serial

```bash
tinygo monitor
```

## Reference

<https://github.com/va1da5/tinygo-pico-ssd1306/blob/main/README.md>
<https://github.com/tinygo-org/drivers>
<https://github.com/tinygo-org/tinydraw>
<https://github.com/Nondzu/ssd1306_font>
