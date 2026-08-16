|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Access | Read  |
| --- | --- |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥0  |

MAC address of the logical link.

This feature must return a 64-bit value representing the full MAC address of the device (i.e. the high and low parts).

### 27.4.15 GevPAUSEFrameReception

|  Name | GevPAUSEFrameReception[GevInterfaceSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Controls whether incoming PAUSE Frames are handled on the given logical link.

### 27.4.16 GevPAUSEFrameTransmission

|  Name | GevPAUSEFrameTransmission[GevInterfaceSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |