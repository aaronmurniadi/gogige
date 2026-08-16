|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Level | Optional  |
| --- | --- |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | NotConnected NotInitialized NotSupported Busy Ready Error Device-specific  |

Reads the status of the filter.

Possible values are:

- NotConnected: The filter controller is physically not connected.
- NotInitialized: The filter controller is not initialized.
- NotSupported: The filter controller is physically connected but not supported.
- Busy: The filter controller executes a feature access/command.
- Ready: The filter controller is ready to use.
- Error: The filter controller encountered an error.

### 23.4.35 Filter

|  Name | Filter[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥0  |

Filter positions in native number system. This is implementation dependent.