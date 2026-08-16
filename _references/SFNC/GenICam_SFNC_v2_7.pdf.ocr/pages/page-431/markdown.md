|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 23.4.12 OpticControllerFirmwareVersion

|  Name | OpticControllerFirmwareVersion[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated String  |

Version of the firmware in the optic controller.

### 23.4.13 OpticControllerTemperature

|  Name | OpticControllerTemperature[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read  |
|  Unit | C  |
|  Visibility | Beginner  |
|  Values | Device-specific  |

Temperature of the optic controller in degrees Celsius (C).

### 23.4.14 ApertureInitialize

|  Name | ApertureInitialize[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Beginner  |