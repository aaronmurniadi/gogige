|  GEN<i>CAM |   | ![img-59.jpeg](img-59.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

command. This is not including the time needed to receive the command or send the acknowledge packet but only the time needed to execute the command. In case a device needs longer to process a command, it must send a pending ack.

The maximum time needed to transfer the message is depending on the link speed and the maximum size of the message.

This number may have direct impact on the behavior of software layers above. It is to be kept as short as possible.

The maximum response time must not exceed 300 ms in order to guarantee a good device's behavior.

Reading this register must not exceed 50 ms processing time.

|  Offset | Hex 1CC  |
| --- | --- |
|  Length | 4  |
|  Access Type | R  |
|  Support | M  |
|  Data Type | UINT32  |
|  Factory Default | Implementation Specific  |

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 32 | Maximum Device Response TimeMaximum time until a device sends a response upon a received command, not including the time needed to send the response over the link in ms.  |

Table 22 – Register Maximum Device Response Time