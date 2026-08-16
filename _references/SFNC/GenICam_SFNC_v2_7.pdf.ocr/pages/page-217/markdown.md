|   | Continuous Device-specific  |
| --- | --- |

Controls the mode for automatic black level adjustment. The exact algorithm used to implement this adjustment is device-specific.

Some other device-specific features might be used to allow the selection of the algorithm.

Possible values are:

- Off: Analog black level is user controlled using BlackLevel.
- Once: Analog black level is automatically adjusted once by the device. Once it has converged, it automatically returns to the Off state.
- Continuous: Analog black level is constantly adjusted by the device.

On top of the previous standard values, a device might also provide device-specific values.

### 6.9 BlackLevelAutoBalance

|  Name | BlackLevelAutoBalance  |
| --- | --- |
|  Category | AnalogControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Off Once Continuous Device-specific  |

Controls the mode for automatic black level balancing between the sensor color channels or taps. The black level coefficients of each channel are adjusted so they are matched.

Possible values are:

- Off: Black level tap balancing is user controlled using BlackLevel.
- Once: Black level tap balancing is automatically adjusted once by the device. Once it has converged, it automatically returns to the Off state.
- Continuous: Black level tap balancing is constantly adjusted by the device.