|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Level | Optional  |
| --- | --- |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Initializes the focal length and makes it ready for use. The focal length position after initialization is implementation dependent. This feature is only implemented if an additional initialization is required after OpticControllerInitialize.

### 23.4.27 FocalLengthStatus

|  Name | FocalLengthStatus[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | NotConnected NotInitialized NotSupported Busy Ready Error Device-specific  |

Reads the status of the focal length.

Possible values are:

- NotConnected: The focal length controller is physically not connected.
- NotInitialized: The focal length controller is not initialized.
- NotSupported: The focal length controller is physically connected but not supported.
- Busy: The focal length controller executes a feature access/command.
- Ready: The focal length controller is ready to use.