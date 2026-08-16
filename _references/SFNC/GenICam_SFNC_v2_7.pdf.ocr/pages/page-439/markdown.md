|  GEN<ì>CAM |   | ![img-143.jpeg](img-143.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

• Error: The focal length controller encountered an error.

23.4.28 FocalLength

|  Name | FocalLength[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | mm  |
|  Visibility | Beginner  |
|  Values | >0  |

Focal length in millimeters (mm).

23.4.29 FocalLengthStepper

|  Name | FocalLengthStepper[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

FocalLengthStepper controls the stepper value of the focal length. 0 is the shortest focal length.

23.4.30 ShutterInitialize

|  Name | ShutterInitialize[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | ICommand  |