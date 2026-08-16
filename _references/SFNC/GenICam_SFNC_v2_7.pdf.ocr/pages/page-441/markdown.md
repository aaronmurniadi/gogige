|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 23.4.32 Shutter

|  Name | Shutter[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥0  |

Controls whether the shutter is opened or closed.

Possible values are:

- 0: The shutter is closed.
- >0: The shutter is opened. Value is implementation dependent. The maximum value is the most opened.

### 23.4.33 FilterInitialize

|  Name | FilterInitialize[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Initializes the filter and makes it ready for use. The filter position after initialization is implementation dependent. This feature is only implemented if an additional initialization is required after OpticControllerInitialize.

### 23.4.34 FilterStatus

|  Name | FilterStatus[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |