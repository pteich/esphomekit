# ESPHomeKit

ESPHomeKit is a simple solution to connect ESPHome devices (https://esphome.io) to HomeKit.
The service is written in Go and can run on any platform like Raspberry PI.

I've written this service because I needed a simple, small and fast solution to make my devices available in 
any HomeKit compatible app like Apple's Home.

It's still more like a proof of concept because I added only what I need for myself.
In my setup ESPHomeKit is running on an Raspberry PI Zero W and works without latency for some time now.

## Setup

You need some ESPHome powered devices with enabled REST api. Have a look at https://esphome.io on
how to setup ESPHome on any ESP8266 or ESP32 board. Its a great tool and supports
almost all ESP devices even the famous Sonoff.

To compile ESPHomeKit you need an working Go environment. 
Compile and run ESPHomeKit with a specific 8 digit pin (given as `-pin` command line flag)
and open Apple Home app. Add a new (untrusted) device (it should already appear in the dialog) and
enter the pin. That's it.

 
 

