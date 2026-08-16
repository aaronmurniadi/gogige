|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Enables the alternate IP destination for stream packets resent due to a packet resend request. When True, the source IP address provided in the packet resend command packet is used. When False, the value set in the GevSCDA[GevStreamChannelSelector] feature is used.

This feature is only valid for GVSP transmitters.

### 27.4.58 GevSCCFGAllInTransmission

|  Name | GevSCCFGAllInTransmission[GevStreamChannelSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | True False  |

Enables the selected GVSP transmitter to use the single packet per data block All-in Transmission mode.

### 27.4.59 GevSCCFGUnconditionalStreaming

|  Name | GevSCCFGUnconditionalStreaming[GevStreamChannelSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | True False  |

Enables the camera to continue to stream, for this stream channel, if its control channel is closed or regardless of the reception of any ICMP messages (such as destination unreachable messages).

### 27.4.60 GevSCCFGExtendedChunkData

|  Name | GevSCCFGExtendedChunkData[GevStreamChannelSelector]  |
| --- | --- |