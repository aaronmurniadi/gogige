|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Access | (Read)/Write  |
| --- | --- |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Initializes the shutter and makes it ready for use. The shutter position after initialization is implementation dependent. This feature is only implemented if an additional initialization is required after OpticControllerInitialize.

### 23.4.31 ShutterStatus

|  Name | ShutterStatus[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | NotConnected NotInitialized NotSupported Busy Ready Error Device-specific  |

Reads the status of the shutter.

Possible values are:

- NotConnected: The shutter controller is physically not connected.
- NotInitialized: The shutter controller is not initialized.
- NotSupported: The shutter controller is physically connected but not supported.
- Busy: The shutter controller executes a feature access/command.
- Ready: The shutter controller is ready to use.
- Error: The shutter controller encountered an error.