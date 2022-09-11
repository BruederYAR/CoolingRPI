# How to use

## Use by standard docker

    Build project in docker
    - docker build -t cooling-rpi .
    Start
    - docker run -d --device /dev/mem --device /dev/gpiomem --volume /sys/class/thermal/thermal_zone0:/usr/local/bin/thermal --restart always cooling-rpi

## Use by docker-compose

    Just start the command
    - docker-compose up -d

# Config file
**Сonfig should be located next to the executable file** 

|                 |               |       
| -------------   |:-------------:| 
| Pin             | 21            | 
| Delay           | 1             | 
| MaxTemperature  | 55            | 
| TemperaturePath | "/usr/local/bin/thermal/temp"      | 