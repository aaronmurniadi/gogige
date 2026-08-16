|  GEN<i>CAM |   | ![img-36.jpeg](img-36.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

#### 4.4.3. WriteMem Command

Any write access start address and length is byte aligned unless the underlying technology states different rules. The number of bytes to write is deduced through the length field of the CCD header.

|  Width (Bytes) | Offset (Bytes) | Description  |
| --- | --- | --- |
|  Prefix  |   |   |
|  CCD (command_id = WRITEMEM_CMD)  |   |   |
|  8 | 0 | register address64 bit register address.  |
|  x | 8 | dataNumber of bytes to write to the remote device's register map.  |
|  Postfix  |   |   |

Table 10 – WriteMem Command SCD-Fields

#### 4.4.4. WriteMem Acknowledge

The WriteMem acknowledge states the result of a WriteMem command.

|  Width (Bytes) | Offset (Bytes) | Description  |
| --- | --- | --- |
|  Prefix  |   |   |
|  CCD-ACK (command_id = WRITEMEM_ACK)  |   |   |
|  2 | 0 | reservedThis reserved field is only sent if the length_written field is sent with the acknowledge. If it is sent it is to be set to 0.  |
|  2 | 2 | length writtenNumber of bytes successfully written to the remote device's register map. The length_written field must only be sent if the according bit in the Device Capability register is set.  |
|  Postfix  |   |   |

Table 11 – WriteMem Ack SCD-Fields

The length field in CCD section of the WriteMem Ack must be set to 0 or 4 depending on the bit in