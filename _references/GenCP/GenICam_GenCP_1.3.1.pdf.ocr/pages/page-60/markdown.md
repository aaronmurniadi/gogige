|  GEN<i>CAM |   | ![img-63.jpeg](img-63.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

#### 5.4.16. Timestamp

A read of this register provides a timestamp of a free running, device internal clock in ns. Before reading, the timestamp register must be latched to the device's internal clock by writing to the Timestamp Latch register.

|  Offset | Hex 1F0  |
| --- | --- |
|  Length | 8  |
|  Access Type | R  |
|  Support | CM  |
|  Data Type | UINT64  |
|  Factory Default | 0  |

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 64 | TimestampDevice Time in ns.  |

Table 28 – Register Timestamp

The Timestamp Supported bit in the Device Capability register indicates if this register is present or not.