# goswaybar

Super lightweight swaybar info.
Somewhat more easy on resources than spawning a shell or executing a python script every second.

![img](https://i.ibb.co/2FmcKgT/image.png)

## Requirements

**goswabar** ~~relies on `acpi` package to retrieve battery statistics and~~ `ip` for network interfaces activity.

## Build

```sh
# Debian / Ubuntu
go build -o swaybar main.go

sudo mv swaybar /usr/bin/swaybar

# edit your swaybar config to add goswaybar
```

## Install

~/.config/sway/config
```
bar {
    position top
    pango_markup enabled

    status_command /usr/bin/swaybar

    colors {
        statusline #ffffff
        background #323232
        inactive_workspace #32323200 #32323200 #5c5c5c
    }
}
```
