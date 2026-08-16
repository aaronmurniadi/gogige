|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

- GigEVision: GigE Vision
- USB3Vision: USB3 Vision
- Custom: Custom transport layer

Corresponds to the INTERFACE_INFO_TLTYPE command of IFGetInfo function.

### 3.2.1.5 InterfaceTLVersionMajor

|  Name | InterfaceTLVersionMajor  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | >0  |

Major version number of the transport layer specification the GenTL Producer interface complies with. The TL version of the Interface can be compared with the TL version of the device to assure compatibility.

### 3.2.1.6 InterfaceTLVersionMinor

|  Name | InterfaceTLVersionMinor  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Minor version number of the transport layer specification the GenTL Producer interface complies with. The TL version of the Interface can be compared with the TL version of the device to assure compatibility.