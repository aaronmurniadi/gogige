|  GEN<i>CAM |   | ![img-64.jpeg](img-64.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

#### 5.4.17. Timestamp Latch

A write with the Timestamp Latch bit set to 1 latches the current device time into the timestamp register.

|  Offset | Hex 1F8  |
| --- | --- |
|  Length | 4  |
|  Access Type | W  |
|  Support | CM  |
|  Data Type | Bitfield  |
|  Factory Default | -  |

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 1 | Timestamp LatchLatch the current device time into the timestamp register. The bit is self-clearing which means that you do not need to set it to 0.  |
|  1 | 31 | ReservedSet to 0.  |

Table 29 – Register Timestamp Latch

The Timestamp Supported bit in the Device Capability register indicates if this register is present or not. This register must be supported if the Timestamp register is supported.