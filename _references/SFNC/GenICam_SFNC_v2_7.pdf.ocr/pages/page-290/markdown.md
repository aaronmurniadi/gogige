|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- SoftwareSignal0, SoftwareSignal1, SoftwareSignal2, ...: Selects the software generated signal to control.

### 13.3 SoftwareSignalPulse

|  Name | SoftwareSignalPulse[SoftwareSignalSelector]  |
| --- | --- |
|  Category | SoftwareSignalControl  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Generates a pulse signal that can be used as a software trigger. This command can be used to trigger other modules that accept a SoftwareSignal as trigger source.