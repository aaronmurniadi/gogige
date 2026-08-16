|  GEN<i>CAM |   | ![img-77.jpeg](img-77.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

can not be sure that the Acknowledge-Request bit is set in the command. Therefore, the received command has to be discarded. The sender will run into a timeout and the normal resend procedure has to be applied.

For other errors like unsupported command_ids the failure procedure as described in the GenCP document is to be applied.

### 1.8. Technology Specific Bootstrap Register Map

|  Width (Bytes) | Offset (Bytes) | Support | Access | Description  |
| --- | --- | --- | --- | --- |
|  4 | 0 | M | R | Supported Baudrates  |
|  4 | 4 | M | (R)W | Current Baudrate  |

Table 36 – Serial BRM

#### 1.8.1. Supported Baudrate

Bitfield indicating the supported baud rates.

|  Offset | Hex 000  |
| --- | --- |
|  Length | 4  |
|  Access Type | R  |
|  Support | M  |
|  Data Type | Bitfield  |
|  Factory Default | Device specific  |