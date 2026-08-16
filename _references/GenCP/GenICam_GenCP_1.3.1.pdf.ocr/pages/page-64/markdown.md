|  GEN<i>CAM |   | ![img-68.jpeg](img-68.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

#### 5.4.20. Protocol Endianness

This register has been deprecated. Its content should be ignored (neither read nor written)

|  Offset | Hex 208  |
| --- | --- |
|  Length | 4  |
|  Access Type |   |
|  Support |   |
|  Data Type |   |
|  Factory Default | Deprecated  |

#### 5.4.21. Implementation Endianness

This register reflects the endianness of the device implementation. By reading the register the host can detect the endianness of the device specific registers.

|  Offset | Hex 20C  |
| --- | --- |
|  Length | 4  |
|  Access Type | R  |
|  Support | CM  |
|  Data Type | UINT32  |
|  Factory Default | Device specific  |

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 32 | Implementation EndiannessEndianness of the device implementation.0 = big-endian0xFFFFFFF = little-endian  |

Table 32 – Register - Implementation Endianness

This register is available if the Endianness Register Supported bit in the Device Capability register is set.