|  GEN<i>CAM |   | ![img-15.jpeg](img-15.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

3.37 DeviceLinkThroughputLimit

|  Name | DeviceLinkThroughputLimit[DeviceLinkSelector]  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | Bps  |
|  Visibility | Expert  |
|  Values | ≥0  |

Limits the maximum bandwidth of the data that will be streamed out by the device on the selected Link. If necessary, delays will be uniformly inserted between transport layer packets in order to control the peak bandwidth.

If the device uses many connections to transmit the data, the feature represents the sum of all the traffic and the bandwidth should be distributed uniformly on the various connections.

Any Transport Layer specific bandwidth controls should be kept in sync with this control as much as possible.

3.38 DeviceLinkConnectionCount

|  Name | DeviceLinkConnectionCount[DeviceLinkSelector]  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥0  |

Returns the number of physical connection of the device used by a particular Link.

3.39 DeviceLinkHeartbeatMode

|  Name | DeviceLinkHeartbeatMode[DeviceLinkSelector]  |
| --- | --- |