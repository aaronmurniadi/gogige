|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Access | Read/Write  |
| --- | --- |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ≥0  |

Controls the destination IP address of the selected stream channel to which a GVSP transmitter must send data stream or the destination IP address from which a GVSP receiver may receive data stream.

### 27.4.70 GevSCSP

|  Name | GevSCSP[GevStreamChannelSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ≥0  |

Indicates the source port of the stream channel.

### 27.4.71 GevSCZoneCount

|  Name | GevSCZoneCount[GevStreamChannelSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | 1 to 32  |

Reports the number of zones per block transmitted on the selected stream channel.