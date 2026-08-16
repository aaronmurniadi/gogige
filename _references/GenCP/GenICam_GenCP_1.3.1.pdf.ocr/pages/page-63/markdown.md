|  ![img-66.jpeg](img-66.jpeg) |   | ![img-67.jpeg](img-67.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

#### 5.4.19. Access Privilege

This register reflects the current access privilege.

|  Offset | Hex 204  |
| --- | --- |
|  Length | 4  |
|  Access Type | RW  |
|  Support | CM  |
|  Data Type | Bitfield  |
|  Factory Default | 0  |

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 3 | Access PrivilegeCurrent Access Privilege as described in 3.20 = Available1 = Open (Exclusive)2-7 = reserved  |
|  3 | 29 | ReservedSet to 0.  |

Table 31 – Register Access Privilege

This register is available if the Access Privilege Supported bit in the Device Capability register is set. In case the Access Privilege register is available and the Heartbeat Enable bit is set in the Device Configuration register, the Access Privilege is reset to 0 after the Heartbeat expired.