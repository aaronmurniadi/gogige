|  GEN<i>CAM |   | ![img-14.jpeg](img-14.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Selects which Link of the device to control.

Generally, a device has only one Link that can be composed of one or many connections. But if there are many, this selector can be used to target a particular Link of the device with certain features.

3.35 DeviceLinkSpeed

|  Name | DeviceLinkSpeed[DeviceLinkSelector]  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Bps  |
|  Visibility | Expert  |
|  Values | ≥0  |

Indicates the speed of transmission negotiated on the specified Link.

Note that this represents the total speed of all the connections of the Link.

3.36 DeviceLinkThroughputLimitMode

|  Name | DeviceLinkThroughputLimitMode[DeviceLinkSelector]  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | OnOff  |

Controls if the DeviceLinkThroughputLimit is active. When disabled, lower level TL specific features are expected to control the throughput. When enabled, DeviceLinkThroughputLimit controls the overall throughput.

Possible values are:

- On: Enables the DeviceLinkThroughputLimit feature.
- Off: Disables the DeviceLinkThroughputLimit feature.