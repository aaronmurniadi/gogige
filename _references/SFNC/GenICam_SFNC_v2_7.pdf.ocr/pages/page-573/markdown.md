|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Disable Power over CoaXPress (PoCXP) for the Link.

This feature shall be present only on receiver or transceiver Devices controlling PoCXP.

### 27.7.25 CxpPoCxpTripReset

|  Name | CxpPoCxpTripReset  |
| --- | --- |
|  Category | CoaXPress  |
|  Level | Optional  |
|  Interface | ICommand  |
|  Access | Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Reset the Power over CoaXPress (PoCXP) Link after an over-current trip on the Device connection(s).

This feature shall be present only on receiver or transceiver Devices controlling PoCXP.

### 27.7.26 CxpPoCxpStatus

|  Name | CxpPoCxpStatus  |
| --- | --- |
|  Category | CoaXPress  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Auto Off Tripped  |

Returns the Power over CoaXPress (PoCXP) status of the Device.

Possible values are:

- Auto: Normal automatic PoCXP operation.