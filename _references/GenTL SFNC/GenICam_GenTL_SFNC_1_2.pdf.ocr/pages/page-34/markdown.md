|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|   |  |  |  |  |  |  | used for the stream channel.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  DeviceStreamChannelPacketSizeMax | O | GEV | IInteger | R/(W) | B | G | Controls desired maximum of the packet size feature to be used for the stream channel.  |
|  DeviceStreamChannelPacketSizeInc | O | GEV | IInteger | R/(W) | B | G | Controls desired increment of the packet size feature to be used for the stream channel.  |
|  DeviceStreamChannelNegotiatePacketSize | O | GEV | ICommand | (R)/W | - | E | Starts negotiation for the optimal packet size considering the remote device, host and their connection path.  |

### 2.4.3 Buffer Handling Control

Contains the features related to GenICam control and access of a specific Data Stream module.

Table 2-17: GenICam Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  BufferHandlingControl | R | All | ICategory | R | - | B | Contains all features of the Data Stream module that control the used buffers.  |
|  StreamAnnouncedBufferCount | M | All | IInteger | R | - | E | Number of announced (known) buffers on this stream.  |
|  StreamBufferHandlingMode | M | All | IEnumeration | R/(W) | - | B | Available buffer handling modes of this Data Stream.  |
|  StreamAnnounceBufferMinimum | M | All | IInteger | R | - | E | Minimal number of buffers to announce to enable selected buffer handling mode.  |
|  StreamDeliveredFrameCount | R | All | IInteger | R | - | E | Number of delivered frames since last acquisition start.  |
|  StreamLostFrameCount | R | All | IInteger | R | - | E | Number of lost frames due to queue underrun.  |
|  StreamInputBufferCount | O | All | IInteger | R | - | E | Number of buffers in the input buffer pool plus the buffers(s) currently being filled.  |
|  StreamOutputBufferCount | R | All | IInteger | R | - | E | Number of buffers in the output buffer queue.  |