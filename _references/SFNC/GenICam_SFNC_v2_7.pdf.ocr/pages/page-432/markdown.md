|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Values | -  |
| --- | --- |

Initializes the aperture and makes it ready for use. The aperture position after initialization is implementation dependent. This feature is only implemented if an additional initialization is required after OpticControllerInitialize.

### 23.4.15 ApertureStatus

|  Name | ApertureStatus[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | NotConnected NotInitialized NotSupported Busy Ready Error Device-specific  |

Reads the status of the aperture.

Possible values are:

- NotConnected: The aperture controller is physically not connected.
- NotInitialized: The aperture controller is not initialized.
- NotSupported: The aperture controller is physically connected but not supported.
- Busy: The aperture controller executes a feature access/command.
- Ready: The aperture controller is ready to use.
- Error: The aperture controller encountered an error.

### 23.4.16 Aperture

|  Name | Aperture[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |