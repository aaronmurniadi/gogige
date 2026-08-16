|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Returns the latched state of the clock servo. When the servo is in a locked state, the value returned is 'Locked'. When the servo is in a non-locked state, a device-specific value can be returned to give specific information. If no device-specific value is available to describe the current state of the clock servo, the value should be 'Unknown'.

Possible values are:

- Unknown: Servo state is unknown.
- Locked: Servo is locked

### 27.3.7 PtpOffsetFromMaster

|  Name | PtpOffsetFromMaster  |
| --- | --- |
|  Category | PtpControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | ns  |
|  Visibility | Expert  |
|  Values |   |

Returns the latched offset from the PTP master clock in nanoseconds.

### 27.3.8 PtpClockID

|  Name | PtpClockID  |
| --- | --- |
|  Category | PtpControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Returns the latched clock ID of the PTP device.

Note: Byte 0 of the IEEE ClockIdentity field is mapped to the MSB.