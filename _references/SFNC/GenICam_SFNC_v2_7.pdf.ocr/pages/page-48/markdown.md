|  DeviceConnectionStatus [DeviceConnectionSelector] | O | IEnumeration | R | - | E | Indicates the status of the specified Connection.  |
| --- | --- | --- | --- | --- | --- | --- |
|  DeviceLinkSelector | R | IInteger | R/(W) | - | B | Selects which Link of the device to control.  |
|  DeviceLinkSpeed [DeviceLinkSelector] | O | IInteger | R | Bps | E | Indicates the speed of transmission negotiated on the specified Link.  |
|  DeviceLinkThroughputLimitMode [DeviceLinkSelector] | R | IEnumeration | R/W | - | E | Controls if the DeviceLinkThroughputLimit is active.  |
|  DeviceLinkThroughputLimit [DeviceLinkSelector] | R | IInteger | R/(W) | Bps | E | Limits the maximum bandwidth of the data that will be streamed out by the device on the selected Link.  |
|  DeviceLinkConnectionCount [DeviceLinkSelector] | O | IInteger | R | - | B | Returns the number of physical connection of the device used by a particular Link.  |
|  DeviceLinkHeartbeatMode [DeviceLinkSelector] | O | IEnumeration | R/W | - | E | Activate or deactivate the Link's heartbeat.  |
|  DeviceLinkHeartbeatTimeout [DeviceLinkSelector] | O | IFloat | R/W | us | G | Controls the current heartbeat timeout of the specific Link.  |
|  DeviceLinkCommandTimeout [DeviceLinkSelector] | O | IFloat | R | us | G | Indicates the command timeout of the specified Link.  |
|  DeviceStreamChannelCount | O | IInteger | R | - | E | Indicates the number of streaming channels supported by the device.  |
|  DeviceStreamChannelSelector | O | IInteger | R/W | - | E | Selects the stream channel to control.  |
|  DeviceStreamChannelType [DeviceStreamChannelSelector] | O | IEnumeration | R | - | G | Reports the type of the stream channel.  |
|  DeviceStreamChannelLink [DeviceStreamChannelSelector] | O | IInteger | R/(W) | - | G | Index of device's Link to use for streaming the specified stream channel.  |
|  DeviceStreamChannelEndianness [DeviceStreamChannelSelector] | O | IEnumeration | R/(W) | - | G | Endianness of multi-byte pixel data for this stream.  |
|  DeviceStreamChannelPacketSize [DeviceStreamChannelSelector] | R | IInteger | R/(W) | B | E | Specifies the stream packet size, in bytes, to send on the selected channel for a Transmitter or specifies the maximum packet size supported by a receiver.  |
|  DeviceEventChannelCount | O | IInteger | R | - | E | Indicates the number of event channels supported by the device.  |