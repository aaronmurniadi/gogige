|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Interface | IInteger  |
| --- | --- |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ≥0  |

This feature indicates the source port for the message channel.

27.4.56 GevStreamChannelSelector

|  Name | GevStreamChannelSelector  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Selects the stream channel to control.

27.4.57 GevSCCFGPacketResendDestination

|  Name | GevSCCFGPacketResendDestination[GevStreamChannelSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | TrueFalse  |