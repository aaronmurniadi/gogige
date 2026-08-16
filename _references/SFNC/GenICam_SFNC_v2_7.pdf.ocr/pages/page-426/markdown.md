|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Figure 23-4: Optic controller and optical axis state machine

The command “Initialize” can also be handled internally which means that it is not required to start initialization with the command “Initialize” by the user but this can also be done automatically at system startup. The detection if an optic controller or optical axis is present or not is quality of implementation.

The state “Busy” is a transition state, so the state machine can switch to this state with every transition shown in the state machine above before reaching the new end state. Depending on the implementation of a feature access or command execution (blocking or non-blocking) this state can be visible to the user or not.

### 23.4 Optic Control features

This section gives the detailed description of all the Optic Control features.

#### 23.4.1 OpticControl

|  Name | OpticControl  |
| --- | --- |
|  Category | Root  |
|  Level | Optional  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category for optical control features.

#### 23.4.2 OpticControllerSelector

|  Name | OpticControllerSelector  |
| --- | --- |
|  Category | OpticControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | OpticController0 (If 0 based), OpticController1,  |