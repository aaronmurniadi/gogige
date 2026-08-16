|  Level | Optional  |
| --- | --- |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Off Once Continuous Device-specific  |

Sets the automatic gain control (AGC) mode. The exact algorithm used to implement AGC is device-specific.

Some other device-specific features might be used to allow the selection of the algorithm.

Possible values are:

- Off: Gain is User controlled using Gain.
- Once: Gain is automatically adjusted once by the device. Once it has converged, it automatically returns to the Off state.
- Continuous: Gain is constantly adjusted by the device.

On top of the previous standard values, a device might also provide device-specific values.

### 6.5 GainAutoBalance

|  Name | GainAutoBalance  |
| --- | --- |
|  Category | AnalogControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Off Once Continuous Device-specific  |