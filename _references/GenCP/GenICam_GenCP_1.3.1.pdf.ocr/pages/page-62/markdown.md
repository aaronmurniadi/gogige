|  GEN<i>CAM |   | ![img-65.jpeg](img-65.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

#### 5.4.18. Timestamp Increment

This register indicates the ns/tick of the device internal clock. This allows the application to deduce the accuracy of the timestamp provided by the bootstrap register. For example a value of 1000 indicates the device clock runs at 1MHz.

|  Offset | Hex 1FC  |
| --- | --- |
|  Length | 8  |
|  Access Type | R  |
|  Support | CM  |
|  Data Type | UINT64  |
|  Factory Default | Device specific  |

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 64 | Timestamp Increment Timestamp increment in ns/tick.  |

Table 30 – Register Timestamp Increment

The Timestamp bit in the Device Capability register indicates if this register is present or not. This register must be supported if the Timestamp register is supported.