|  GEN<i>CAM |   | ![img-51.jpeg](img-51.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

|  Width (Bytes) | Offset (Bytes) | Support | Access | Description  |
| --- | --- | --- | --- | --- |
|  4 | 0x001E8 | CM | RW | Heartbeat TimeoutHeartbeat Timeout in ms  |
|  4 | 0x001EC | CM | RW | Message Channel IDchannel_id used for the message channel  |
|  8 | 0x001F0 | CM | R | TimestampLast latched device time in ns  |
|  4 | 0x001F8 | CM | W | Timestamp Latch  |
|  8 | 0x001FC | CM | R | Timestamp Increment  |
|  4 | 0x00204 | CM | RW | Access Privilege  |
|  4 | 0x00208 |  |  | Reserved (deprecated Protocol Endianness, do not reuse)  |
|  4 | 0x0020C | CM | R | Implementation EndiannessEndianness of device implementation registers  |
|  64 | 0x00210 | CM | R | Device Software Interface VersionVersion of the public software interface of the device.  |
|  64944 | 0x00250 | M | no | Reserved Register Space  |

Table 19 – Technology agnostic BRM

- Width Size of the register in bytes.
- Offset Address of the register (Offset in Bytes) in the device's BRM
- Support M=Mandatory/R=Recommended/ CM=Conditional Mandatory (depending on the capability bits)
- Access R=READONLY, W=WRITEONLY, RW=READWRITE
- Description Name and very short hint on the meaning