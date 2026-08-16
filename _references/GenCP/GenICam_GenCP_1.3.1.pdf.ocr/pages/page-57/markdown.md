|  GEN<i>CAM |   | ![img-61.jpeg](img-61.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 64 | SBRM AddressTechnology Specific Bootstrap Register Map Address  |

Table 24 – Register Technology Specific Bootstrap Register Map

#### 5.4.13. Device Configuration

Device Configuration bits describing implementation specific details.

|  Offset | Hex 1E0  |
| --- | --- |
|  Length | 8  |
|  Access Type | RW  |
|  Support | M  |
|  Data Type | Bitfield  |
|  Factory Default | Device specific  |

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 1 | Heartbeat EnableSet to enable the Heartbeat Timer. The Access Privilege Supported bit in the Device Capability register indicates if this bit is available or not. If it is not available it must be set to 0.  |
|  1 | 1 | MultiEvent EnableSet to allow multiple events in a single event command packet. This bit is only available if the MultiEvent Supported bit is set in the Device Capability register. Otherwise it must be set to 0.  |
|  2 | 62 | ReservedSet to 0.  |

Table 25 – Register Device Configuration

#### 5.4.14. Heartbeat Timeout

The register is available if the Access Privilege Supported bit in the Device Capability register is set. If the Heartbeat expires the communication parameters of a device are reset, for example the baud rate of a serial device. It is technology dependent which parameters are affected. After a Heartbeat timeout, a host should be able to communicate with a device using default