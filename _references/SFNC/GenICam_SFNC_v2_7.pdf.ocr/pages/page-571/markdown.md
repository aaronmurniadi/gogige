|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Interface | ICommand  |
| --- | --- |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Resets the selected Cxp Error Counter on the connection selected by CxpConnectionSelector. The counter starts counting events immediately after the reset.

### 27.7.21 CxpErrorCounterValue

|  Name | CxpErrorCounterValue[CxpConnectionSelector][CxpErrorCounterSelector]  |
| --- | --- |
|  Category | CoaXPress  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Reads the current value of the selected Cxp Error Counter on the connection selected by CxpConnectionSelector.

### 27.7.22 CxpErrorCounterStatus

|  Name | CxpErrorCounterStatus[CxpConnectionSelector][CxpErrorCounterSelector]  |
| --- | --- |
|  Category | CoaXPress  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |