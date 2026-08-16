|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Latches the current values from the device's PTP clock data set.

### 27.3.5 PtpStatus

|  Name | PtpStatus  |
| --- | --- |
|  Category | PtpControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Initializing Faulty Disabled Listening PreMaster Master Passive Uncalibrated Slave  |

Returns the latched state of the PTP clock.

The state is indicated by values 1 to 9, corresponding to the states INITIALIZING, FAULTY, DISABLED, LISTENING, PRE_MASTER, MASTER, PASSIVE, UNCALIBRATED, and SLAVE. Refer to the IEEE 1588-2008 specification for additional information.

### 27.3.6 PtpServoStatus

|  Name | PtpServoStatus  |
| --- | --- |
|  Category | PtpControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Unknown Locked Device-specific  |