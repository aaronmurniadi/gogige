|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Values | CounterActive CounterOverflow  |
| --- | --- |

Returns the current status of the selected Cxp Error Counter on the connection selected by CxpConnectionSelector.

Possible values are:

- CounterActive: The counter is actively counting errors.
- CounterOverflow: The counter exceeded its maximum error count.

### 27.7.23 CxpPoCxpAuto

|  Name | CxpPoCxpAuto  |
| --- | --- |
|  Category | CoaXPress  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Activate automatic control of the Power over CoaXPress (PoCXP) for the Link.

This feature shall be present only on receiver or transceiver Devices controlling PoCXP.

### 27.7.24 CxpPoCxpTurnOff

|  Name | CxpPoCxpTurnOff  |
| --- | --- |
|  Category | CoaXPress  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |