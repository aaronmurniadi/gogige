|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Initializes the magnification and makes it ready for use. The magnification position after initialization is implementation dependent. This feature is only implemented if an additional initialization is required after OpticControllerInitialize.

### 23.4.40 MagnificationStatus

|  Name | MagnificationStatus[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | NotConnected NotInitialized NotSupported Busy Ready Error Device-specific  |

Reads the status of the magnification.

Possible values are:

- NotConnected: The magnification controller is physically not connected.
- NotInitialized: The magnification controller is not initialized.
- NotSupported: The magnification controller is physically connected but not supported.
- Busy: The magnification controller executes a feature access/command.
- Ready: The magnification controller is ready to use.
- Error: The magnification controller encountered an error.

### 23.4.41 Magnification

|  Name | Magnification[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IFloat  |