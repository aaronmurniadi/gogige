|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|   | OpticController2, ...  |
| --- | --- |

Selects which optic controller to configure.

Possible values are:

- OpticController0: Selects Optic Controller 0.
- OpticController1: Selects Optic Controller 1.
- OpticController2: Selects Optic Controller 2.
• ...

#### 23.4.3 OpticControllerInitialize

|  Name | OpticControllerInitialize[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Initializes the optic controller and makes it ready for use.

#### 23.4.4 OpticControllerDisconnect

|  Name | OpticControllerDisconnect[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Closes the optic controller connection and prepares it for removal.