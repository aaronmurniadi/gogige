|  Unit | -  |
| --- | --- |
|  Visibility | Guru  |
|  Values | Off On  |

Enables the unconditional action command mode where action commands are processed even when the primary control channel is closed.

Possible values are:

- Off: Unconditional mode is disabled.
- On: Unconditional mode is enabled.

### 14.2.3 ActionDeviceKey

|  Name | ActionDeviceKey  |
| --- | --- |
|  Category | ActionControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Write-Only  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ≥ 0  |

Provides the device key that allows the device to check the validity of action commands. The device internal assertion of an action signal is only authorized if the ActionDeviceKey and the action device key value in the protocol message are equal.

### 14.2.4 ActionQueueSize

|  Name | ActionQueueSize  |
| --- | --- |
|  Category | ActionControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |