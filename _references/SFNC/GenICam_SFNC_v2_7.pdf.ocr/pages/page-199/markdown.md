|  Unit | us  |
| --- | --- |
|  Visibility | Beginner  |
|  Values | ≥0  |

Sets the Exposure time when ExposureMode is Timed and ExposureAuto is Off. This controls the duration where the photosensitive cells are exposed to light.

### 5.7.5 ExposureAuto

|  Name | ExposureAuto  |
| --- | --- |
|  Category | AcquisitionControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Off Once Continuous Device-specific  |

Sets the automatic exposure mode when ExposureMode is Timed. The exact algorithm used to implement this control is device-specific.

Some other device-specific features might be used to allow the selection of the algorithm.

Possible values are:

- Off: Exposure duration is user controlled using ExposureTime.
- Once: Exposure duration is adapted once by the device. Once it has converged, it returns to the Off state.
- Continuous: Exposure duration is constantly adapted by the device to maximize the dynamic range.

On top of the previous standard values, a device might also provide device-specific values.

### 5.8 Multi-slope Exposure Control features