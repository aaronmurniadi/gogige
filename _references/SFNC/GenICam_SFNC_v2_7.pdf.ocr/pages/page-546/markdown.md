|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Values | 0 to 3  |
| --- | --- |

Index of the logical link to use.

Specific streams might be hard-coded to specific logical links. Therefore this field might be read-only on certain devices.

### 27.4.63 GevSCPHostPort

|  Name | GevSCPHostPort[GevStreamChannelSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ≥0  |

Controls the port of the selected channel to which a GVSP transmitter must send data stream or the port from which a GVSP receiver may receive data stream. Setting this value to 0 closes the stream channel.

### 27.4.64 GevSCPSFireTestPacket

|  Name | GevSCPSFireTestPacket[GevStreamChannelSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | True False  |

Sends a test packet. When this feature is set, the device will fire one test packet.

The "don't fragment" bit of IP header must be set for this test packet.