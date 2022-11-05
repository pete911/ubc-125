# Command Line tool for programming Uniden UBC125 Scanner

Tool is cross-platform, but unfortunately macOS does not detect/open serial port for this scanner. For now, I connect
the scanner to my [raspberryPi](https://www.raspberrypi.com) and run commands from there.

If you are looking for GUI on Windows I would recommend [scan125](https://www.nick-bailey.co.uk/scan125/).

## commands

```
Available Commands:
  system      system settings
    contrast    lcd contrast level
    squelch     squelch level
    volume      volume level
    weather     weather priority

Global Flags:
  -d, --dry-run   dry run, do not connect
  -v, --verbose   verbose, print logs
```

## example

```
ubc-125 system volume
? select volume  [Use arrows to move, type to filter]
  0
  1
> 2
  3
  4
```
