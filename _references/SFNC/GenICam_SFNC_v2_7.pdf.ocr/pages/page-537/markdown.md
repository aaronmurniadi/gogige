|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|   | Faulty Disabled Listening PreMaster Master Passive Uncalibrated Slave  |
| --- | --- |

This feature is deprecated (See PtpStatus). It was used to Provide the status of the IEEE 1588 clock.

### 27.4.41 GevGVCPExtendedStatusCodesSelector

|  Name | GevGVCPExtendedStatusCodesSelector  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | Version1_1 Version2_0  |

Selects the GigE Vision version to control extended status codes for.

### 27.4.42 GevGVCPExtendedStatusCodes

|  Name | GevGVCPExtendedStatusCodes[GevGVCPExtendedStatusCodesSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | True False  |