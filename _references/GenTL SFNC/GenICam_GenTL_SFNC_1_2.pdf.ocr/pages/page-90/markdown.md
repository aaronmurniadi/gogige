|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.3.2 Device Control

The Device Control section contains all features related to control specific properties of the Device module.

#### 3.3.2.1 DeviceControl

|  Name | DeviceControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains all Device Control features of the Device module.

#### 3.3.2.2 DeviceEndianessMechanism

|  Name | DeviceEndianessMechanism  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Legacy Standard  |

Identifies the endianess handling mode.

- Legacy: Handling the device endianess according to GenICam Schema 1.0
- Standard: Handling the device endianess according to GenICam Schema 1.1 and later