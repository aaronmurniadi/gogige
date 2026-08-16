|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Level | Recommended  |
| --- | --- |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | B  |
|  Visibility | Expert  |
|  Values | >0  |

This GigE Vision specific feature corresponds to DeviceStreamChannelPacketSize and should be kept in sync with it. It specifies the stream packet size, in bytes, to send on the selected channel for a GVSP transmitter or specifies the maximum packet size supported by a GVSP receiver.

This does not include data leader and data trailer and the last data packet which might be of smaller size (since packet size is not necessarily a multiple of block size for stream channel).

If a device cannot support the requested packet size, then it must not fire a test packet when requested to do so.

### 27.4.68 GevSCPD

|  Name | GevSCPD[GevStreamChannelSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit |   |
|  Visibility | Expert  |
|  Values | ≥0  |

Controls the delay (in GEV timestamp counter unit) to insert between each packet for this stream channel. This can be used as a crude flow-control mechanism if the application or the network infrastructure cannot keep up with the packets coming from the device.

### 27.4.69 GevSCDA

|  Name | GevSCDA[GevStreamChannelSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |