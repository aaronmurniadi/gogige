|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 27.4 GigE Vision Control

This section describes the GigE Vision specific transport layer control features.

#### 27.4.1 GigEVision

|  Name | GigEVision  |
| --- | --- |
|  Category | TransportLayerControl  |
|  Level | Optional  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category that contains the features pertaining to the GigE Vision transport layer of the device.

#### 27.4.2 GevVersionMajor (Deprecated)

|  Name | GevVersionMajor  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | >0  |

This feature is deprecated (See DeviceTLVersionMajor). It was representing the major version of the specification.

For instance, GigE Vision version 1.0 would have the major version set to 1.

#### 27.4.3 GevVersionMinor (Deprecated)

|  Name | GevVersionMinor  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Recommended  |