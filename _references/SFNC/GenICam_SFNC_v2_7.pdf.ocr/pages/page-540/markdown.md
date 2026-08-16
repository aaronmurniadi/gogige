|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

If the standard IDs mode is not supported, this feature should return On and be read only.

This feature is not applicable for GigE Vision 1.x devices.

### 27.4.48 GevCCP

|  Name | GevCCP  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | OpenAccess ExclusiveAccess ControlAccess ControlAccessSwitchoverActive  |

Controls the device access privilege of an application.

Only one application is allowed to control the device. This application is able to write into device's registers. Other applications can read device's register only if the controlling application does not have the exclusive privilege.

### 27.4.49 GevPrimaryApplicationSocket

|  Name | GevPrimaryApplicationSocket  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ≥0  |

Returns the UDP source port of the primary application.