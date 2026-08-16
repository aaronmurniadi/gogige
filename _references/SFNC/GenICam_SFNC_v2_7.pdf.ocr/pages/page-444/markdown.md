|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- NotInitialized: The stabilization controller is not initialized.
- NotSupported: The stabilization controller is physically connected but not supported.
- Busy: The stabilization controller executes a feature access/command.
- Ready: The stabilization controller is ready to use.
- Error: The stabilization controller encountered an error.

### 23.4.38 Stabilization

|  Name | Stabilization[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥0  |

Control of image stabilization function build into the optic controller.

Possible values are:

- 0: Image stabilization is disabled.
- >0: Image stabilization is enabled. Value is implementation dependent.

### 23.4.39 MagnificationInitialize

|  Name | MagnificationInitialize[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |