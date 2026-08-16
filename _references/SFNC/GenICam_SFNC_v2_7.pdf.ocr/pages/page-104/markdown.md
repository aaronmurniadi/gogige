|  Values | >0  |
| --- | --- |

Indicates the command timeout of the specified Link. This corresponds to the maximum response time of the device for a command sent on that link.

Note that some Transport Layer Protocols might support that the device responds (within the DeviceLinkCommandTimeout period) that the completion of a particularly long command will be delayed by a specific amount of time. This notion is generally known as a "Pending Acknowledge" command.

### 3.42 DeviceStreamChannelCount

|  Name | DeviceStreamChannelCount  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | Integer  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Indicates the number of streaming channels supported by the device.

### 3.43 DeviceStreamChannelSelector

|  Name | DeviceStreamChannelSelector  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | Integer  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Selects the stream channel to control.

### 3.44 DeviceStreamChannelType

|  Name | DeviceStreamChannelType [DeviceStreamChannelSelector]  |
| --- | --- |
|  Category | DeviceControl  |