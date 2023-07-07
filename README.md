Libre Hardware Monitor To Telegraf
----------------------------------

Takes output from [Libre Hardware Monitor's][1] [built in HTTP server][2] and
converts it into [InfluxDB input line format][3].

[1]: https://github.com/LibreHardwareMonitor/LibreHardwareMonitor/
[2]: http://localhost:8085/data.json
[3]: https://docs.influxdata.com/influxdb/v1.7/write_protocols/line_protocol_tutorial/

I hook into Telegraf thusly:

    [[inputs.exec]]
      commands = [
         '''"C:\Program Files\Telegraf\libre-hardware-monitor-to-telegraf.exe"'''
      ]
      timeout = "5s"
      data_format = "influx"

Building
--------

    go build

## Cross Compiling

    make

Copyright and License
---------------------

License: [ISC](https://en.wikipedia.org/wiki/ISC_license)

Copyright © 2019-2020 David Caldwell \<david@porkrind.org\>

Permission to use, copy, modify, and/or distribute this software for any
purpose with or without fee is hereby granted, provided that the above
copyright notice and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY
SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION
OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN
CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
