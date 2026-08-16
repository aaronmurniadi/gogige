|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Level | Optional  |
| --- | --- |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category that contains the features pertaining to the Camera Link transport layer of the device.

This category is optional especially if the device supports only one transport layer.

### 27.6.2 CIConfiguration

|  Name | CIConfiguration  |
| --- | --- |
|  Category | CameraLink  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Base Medium Full DualBase EightyBit Deca (Deprecated)  |

This Camera Link specific feature describes the configuration used by the camera. It helps especially when a camera is capable of operation in a non-standard configuration, and when the features PixelSize, SensorDigitizationTaps, and DeviceTapGeometry do not provide enough information for interpretation of the image data provided by the camera.

Possible values are:

- Base: Standard base configuration described by the Camera Link standard.
- Medium: Standard medium configuration described by the Camera Link standard.
- Full: Standard full configuration described by the Camera Link standard.