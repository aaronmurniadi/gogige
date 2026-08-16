|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 27.3.9 PtpParentClockID

|  Name | PtpParentClockID  |
| --- | --- |
|  Category | PtpControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Returns the latched parent clock ID of the PTP device. The parent clock ID is the clock ID of the current master clock.

Note: Byte 0 of the IEEE ClockIdentity field is mapped to the MSB.

### 27.3.10 PtpGrandmasterClockID

|  Name | PtpGrandmasterClockID  |
| --- | --- |
|  Category | PtpControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Returns the latched grandmaster clock ID of the PTP device. The grandmaster clock ID is the clock ID of the current grandmaster clock.

Note: Byte 0 of the IEEE ClockIdentity field is mapped to the MSB.

### 27.3.11 PtpMeanPropagationDelay

|  Name | PtpMeanPropagationDelay  |
| --- | --- |
|  Category | PtpControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | ns  |
|  Visibility | Expert  |
|  Values | -  |