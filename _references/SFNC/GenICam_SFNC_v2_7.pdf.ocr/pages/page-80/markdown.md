Precision Time Protocol

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  PtpControl | O | ICategory | R | - | E | Category that contains the features related to the Precision Time Protocol (PTP) of the device.  |
|  PtpEnable | O | IBoolean | R/W | - | E | Enables the Precision Time Protocol (PTP).  |
|  PtpClockAccuracy | O | IEnumeration | R | - | E | Indicates the expected accuracy of the device PTP clock when it is the grandmaster, or in the event it becomes the grandmaster.  |
|  PtpDataSetLatch | O | ICommand | W | - | E | Latches the current values from the device's PTP clock data set.  |
|  PtpStatus | O | IEnumeration | R | - | E | Returns the latched state of the PTP clock.  |
|  PtpServoStatus | O | IEnumeration | R | - | E | Returns the latched state of the clock servo.  |
|  PtpOffsetFromMaster | O | IInteger | R | ns | E | Returns the latched offset from the PTP master clock in nanoseconds.  |
|  PtpClockID | O | IInteger | R | - | E | Returns the latched clock ID of the PTP device.  |
|  PtpParentClockID | O | IInteger | R | - | E | Returns the latched parent clock ID of the PTP device.  |
|  PtpGrandmasterClockID | O | IInteger | R | - | E | Returns the latched grandmaster clock ID of the PTP device.  |
|  PtpMeanPropagationDelay | O | IInteger | R | ns | E | Returns the latched mean propagation delay from the current PTP master clock in nanoseconds.  |

GigE Vision

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  GigEVision | O | ICategory | R | - | B | Category that contains the features pertaining to the GigE Vision transport layer of the device.  |
|  GevVersionMajor | R | IInteger | R | - | I | This feature is deprecated (See DeviceTLVersionMajor).  |
|  GevVersionMinor | R | IInteger | R | - | I | This feature is deprecated (See DeviceTLVersionMinor).  |
|  GevDeviceModeIsBigEndian | O | IBoolean | R | - | I | This feature is deprecated (See DeviceRegistersEndianness).  |
|  GevDeviceClass | O | IEnumeration | R | - | I | This feature is deprecated (See DeviceType).  |