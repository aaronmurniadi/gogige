|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 23.4.5 OpticControllerAbort

|  Name | OpticControllerAbort[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Aborts the current command or feature access. This will end the process without completing it or having reached the desired end position or value. If no command or feature access is in progress, the command is ignored.

### 23.4.6 OpticControllerStatus

|  Name | OpticControllerStatus[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | NotConnected NotInitialized NotSupported Busy Ready Error Device-specific  |

Reads the status of the optic controller.

Possible values are:

- NotConnected: The optic controller is physically not connected.