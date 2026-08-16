|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Sets the numerical aperture of a lens. It is inversely proportional to the f/#.

### 23.4.19 FocusInitialize

|  Name | FocusInitialize[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Initializes the focus and makes it ready for use. The focus position after initialization is implementation dependent. This feature is only implemented if an additional initialization is required after OpticControllerInitialize.

### 23.4.20 FocusStatus

|  Name | FocusStatus[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | NotConnected NotInitialized NotSupported Busy Ready Error Device-specific  |