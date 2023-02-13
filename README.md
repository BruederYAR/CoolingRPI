# Cooling Raspberry Pi
This application will help you to automatically turn on active cooling when the processor temperature exceeds the specified threshold 

All you need is to connect your cooler to GPIO and specify the pin in the configuration file. For more customization, see below

## How to build

### Use by standard docker
    Build project in docker
    - docker build -t cooling-rpi .
    Start
    - docker run -d --device /dev/mem --device /dev/gpiomem --volume /sys/class/thermal/thermal_zone0:/usr/local/bin/thermal --restart always cooling-rpi

### Use by docker-compose
    Just start the command
    - docker-compose up -d

## Config file
**Сonfig should be located next to the executable file** 
| Name            | Value                           | Description                                  |               
| -------------   |:-------------------------------:|:--------------------------------------------:| 
| Pin             | 21                              | Pin to which the cooler is connected         |
| Delay           | 1                               | Delay in seconds to read CPU temperature     |
| MaxTemperature  | 55                              | Threshold at which cooling is triggered      |
| TemperaturePath | "/usr/local/bin/thermal/temp"   | Path to the CPU temperature information file |