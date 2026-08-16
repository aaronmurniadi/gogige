|  GEN<i>CAM |   | ![img-62.jpeg](img-62.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

communication parameters. The Heartbeat is triggered/reset through any register access initiated by the host.

|  Offset | Hex 1E8  |
| --- | --- |
|  Length | 4  |
|  Access Type | RW  |
|  Support | CM  |
|  Data Type | UINT32  |
|  Factory Default | 3000  |

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 32 | Heartbeat TimeoutHeartbeat timeout in milliseconds.  |

Table 26 – Register Heartbeat Timeout

#### 5.4.15. Message Channel ID

The register contains the channel_id to be used for the message channel. This register has to be written by the host to inform the device which channel to use for the message channel. At start up the register contains 0 indicating that it is not initialized by the host. A channel_id of 0 for the Message Channel is not valid since 0 is used for the command channel.

|  Offset | Hex 1EC  |
| --- | --- |
|  Length | 4  |
|  Access Type | RW  |
|  Support | CM  |
|  Data Type | UINT32  |
|  Factory Default | 0  |

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 32 | Channel IDMessage Channel ID.  |

Table 27 – Register Message Channel ID