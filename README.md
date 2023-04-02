# glitch-clock
A glitchy clock for the terminal

## Why do I need another clock?
Good point. You probably don't. I bet if you look around you right now there are 3-4 clocks in
your field of view. I don't need another normal clock either.

However, I spend a lot of time looking at dashboards and logs that use UTC time for timestamps. I could subtract 6 or
7 hours depending on the time of year to correlate to my local time, but that's too much mathing. I wanted a clock
that runs in a terminal window and just shows me the UTC time and date. That what glitch-clock does. It can also show
local if you need yet another clock.

## Usage
There are just a few command line options:
```bash
Usage of glitch-clock:
  -local-time
        Use local time instead of UTC
  -sep string
        Date separator ('-' or '/', '-' by default) (default "-")
  -show-date
        Show date (default true)
```
To exit **glitch-clock**, just hit `Ctrl-C` or `Esc`.