|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

## 13 Software Signal Control

The Software Signal Control chapter describes the model and features related to the control and the generation of software generated signals.

A software signal is a general source and can be used as a trigger signal source for diverse functions in other SFNC modules.

### 13.1 SoftwareSignalControl

|  Name | SoftwareSignalControl  |
| --- | --- |
|  Category | Root  |
|  Level | Optional  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category that contains the Software Signal Control features.

### 13.2 SoftwareSignalSelector

|  Name | SoftwareSignalSelector  |
| --- | --- |
|  Category | SoftwareSignalControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | SoftwareSignal0 (If 0 based) SoftwareSignal1, SoftwareSignal2, ...  |

Selects which Software Signal features to control.

Possible values are: