|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Interface | IInteger  |
| --- | --- |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | ≥0  |

This feature is deprecated (See DeviceLinkCommandTimeout). It was used to indicate the longest GVCP command execution time before a device returns a PENDING_ACK.

#### 27.4.46 GevPrimaryApplicationSwitchoverKey

|  Name | GevPrimaryApplicationSwitchoverKey  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Write-Only  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ≥ 0  |

Controls the key to use to authenticate primary application switchover requests.

#### 27.4.47 GevGVSPExtendedIDMode

|  Name | GevGVSPExtendedIDMode  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | OffOn  |

Enables the extended IDs mode.