|  GEN<i>CAM |   | ![img-76.jpeg](img-76.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

### 1.5. Serial Prefix

For a serial connection, we do not have to handle addressing between device and host, because it is a point to point connection, but we do need to mimic multiple communication channels. In addition a packet preamble allows to identify a GenCP packet and differentiate it from other (ASCII based) protocols.

For the default communication channel the channel_id is always 0.

|  Width (Bytes) | Offset (Bytes) | Description  |
| --- | --- | --- |
|  2 | 0 | 0x0100 (preamble)Leading binary 0x1 (SOH) 0x00 (NULL) send on the link to identify a GenCP package to allow the application layers above to distinguish between different protocols.  |
|  2 | 2 | CCD-CRC-16CRC-16 build from the channel_id and CCD  |
|  2 | 4 | SCD-CRC-16CRC-16 build from channel_id, CCD and SCD  |
|  2 | 6 | channel_idA 16bit number identifying a communication channel. Channel 0 is reserved the for the default communication channel.  |

Table 35 – Serial Prefix

This prefix layout is identical for command and acknowledge packets. The checksums are the 16-bit one's complements of the one's complement sums. The computation algorithm is the same as for the UDP checksum referenced in RFC 768.

### 1.6. Serial Postfix

We do not need a Postfix section for serial links.

### 1.7. Packet failure

In case the device or the host receives a command packet with an invalid CCD-CRC, the receiver