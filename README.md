# ESPHomeKit

ESPHomeKit is a simple solution to connect ESPHome devices (https://esphome.io) to HomeKit.
The service is written in Go and can run on any platform like Raspberry PI.

I've written this service because I needed a simple, small and fast solution to make my devices available in 
any HomeKit compatible app like Apple's Home.

It's still more like a proof of concept because I added only what I need for myself. At my home ESPHomeKit is now running on an Raspberry PI Zero W and works without latency or problems for a long time.

## Setup and Run

You need some ESPHome powered devices with enabled REST api. Have a look at https://esphome.io on
how to setup ESPHome on any ESP8266 or ESP32 board. Its a great tool and supports
almost all ESP devices even the famous Sonoff.

To compile ESPHomeKit you need an working Go environment. 
Compile (`make build-arm` for Raspberry Pis) and run ESPHomeKit with a specific 8 digit pin (given as `-pin` command line flag) and open Apple's Home app. Tap to add a new device ("EspHomeKit Bridge" should already appear in the dialog) accept the message about untrusted devices and enter the pin. That's it.

Commandline flags:
`-pin` the 8 digit pin for HomeKit
`-config` path to config file with all you accessories
`-logConsole` log in human readable format to console, otherwise json log is created
`-logLevel` one of debug info or error - defaults to debug

## Limitations

By now only color lights and temperature sensors are supported because that's what
I needed.
 

