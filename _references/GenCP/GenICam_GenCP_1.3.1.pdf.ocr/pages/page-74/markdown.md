|  GEN<i>CAM |   | ![img-79.jpeg](img-79.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

On a serial link, a baud rate of 9600 must be supported and set at start up so that an initial communication can be established.

#### 1.8.2. Current Baudrate

Register indicating the currently used baud rate. The register is RW with the exception that only one baud rate is supported. In this case the register may also be read only.

|  Offset | Hex 004  |
| --- | --- |
|  Length | 4  |
|  Access Type | RW  |
|  Support | M  |
|  Data Type | Bitfield  |
|  Factory Default | 1  |

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 32 | Current BaudrateBAUDRATE_9600 = 0x00000001BAUDRATE_19200 = 0x00000002BAUDRATE_38400 = 0x00000004BAUDRATE_57600 = 0x00000008BAUDRATE_115200 = 0x00000010BAUDRATE_230400 = 0x00000020BAUDRATE_460800 = 0x00000040BAUDRATE_921600 = 0x00000080A single bit may be set according to the current baudrate setting. 0 is an invalid value.  |

Table 38 – Register – Serial – Current Baudrate

In case the Heartbeat timeout of a serial device expires, the device must fall back to factory default communication parameters (baud rate) in order to allow further communication with the host.