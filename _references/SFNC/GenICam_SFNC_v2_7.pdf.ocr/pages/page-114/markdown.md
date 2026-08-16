Typical values listed should be used whenever possible. Arbitrary values can also be used by defining new enumeration entries.

Possible values are:

- Baud9600: Serial port speed of 9600 baud.
- Baud19200: Serial port speed of 19200 baud.
- Baud38400: Serial port speed of 38400 baud.
- Baud57600: Serial port speed of 57600 baud.
- Baud115200: Serial port speed of 115200 baud.
- Baud230400: Serial port speed of 230400 baud.
- Baud460800: Serial port speed of 460800 baud.
- Baud921600: Serial port speed of 921600 baud.

### 3.66 Timestamp

|  Name | Timestamp  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Recommended  |
|  Interface | Integer  |
|  Access | Read  |
|  Unit | ns  |
|  Visibility | Expert  |
|  Values | ≥0  |

Reports the current value of the device timestamp counter.

The same timestamp counter is used for tagging images, chunk and event data.

Note that the increment of the Timestamp feature must correspond to the resolution of the devices's timestamp in nanoseconds.

### 3.67 TimestampReset

|  Name | TimestampReset  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | Command  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |