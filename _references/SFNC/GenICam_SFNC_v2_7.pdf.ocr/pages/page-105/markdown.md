|  Level | Optional  |
| --- | --- |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | Transmitter Receiver  |

Reports the type of the stream channel.

Possible values are:

- Transmitter: Data stream transmitter channel.
- Receiver: Data stream receiver channel.

### 3.45 DeviceStreamChannelLink

|  Name | DeviceStreamChannelLink [DeviceStreamChannelSelector]  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | ≥0  |

Index of device's Link to use for streaming the specified stream channel.

### 3.46 DeviceStreamChannelEndianness

|  Name | DeviceStreamChannelEndianness [DeviceStreamChannelSelector]  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | Big  |