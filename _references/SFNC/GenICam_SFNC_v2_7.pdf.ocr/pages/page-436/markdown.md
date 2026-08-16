|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Values | DistanceBased ImageBased Device-specific  |
| --- | --- |

Sets automatic focus mode.

Possible values are:

- DistanceBased: The OpticController has a built-in distance sensor and computes the focus based on the detected distance (open loop).
- ImageBased: The OpticController evaluates the image for closed loop control (closed loop).

### 23.4.23 FocusAuto

|  Name | FocusAuto [OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Off Once Continuous Device-specific  |

Sets automatic focus. If enabled, the optic device shall be set into a mode of operation, where the lens automatically finds the best possible focus.

Possible values are:

- Off: The focus is set manually.
- Once: The focus is adapted once by the device. Once it has converged, it returns to the Off state.
- Continuous: The focus is constantly adapted by the device to maximize the dynamic range.