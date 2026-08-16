|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Reads the status of the focus.

Possible values are:

- NotConnected: The focus controller is physically not connected.
- NotInitialized: The focus controller is not initialized.
- NotSupported: The focus controller is physically connected but not supported.
- Busy: The focus controller executes a feature access/command.
- Ready: The focus controller is ready to use.
- Error: The focus controller encountered an error.

#### 23.4.21 FocusStepper

|  Name | FocusStepper[OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

FocusStepper controls the stepper value of the focus. E.g. ObjectSensorDistance or FocalPower. 0 is the closest focus.

#### 23.4.22 FocusAutoMode

|  Name | FocusAutoMode [OpticControllerSelector]  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |