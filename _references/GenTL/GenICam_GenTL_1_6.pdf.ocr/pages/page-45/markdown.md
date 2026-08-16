|  ![img-59.jpeg](img-59.jpeg)CAN |   | ![img-60.jpeg](img-60.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

In case the returned information of DSGetInfo with STREAM_INFO_DEFINES_PAYLOADSIZE is false the Consumer needs to inquire the PayloadSize through the node map of the remote device. The remote device port can be retrieved via the DevGetPort function from the according Device module. The GenTL Consumer has to select the streaming channel in the remote device and read the "PayloadSize" standard feature.

In any case the GenTL Producer together with the underlying technology must provide a way to retrieve the payload size. When the device does not provide the PayloadSize feature (for example in case of a GenTL Producer which is implementing an interface standard which is not specifying PayloadSize as a mandatory feature), the GenTL Producer itself must report the required payload size using stream info commands STREAM_INFO_DEFINES_PAYLOADSIZE and STREAM_INFO_PAYLOAD_SIZE. Failure to query the required payload size would typically disallow the GenTL Consumer to set up the acquisition properly. It might try to calculate the payload size based on the device configuration, but such calculation would never be reliable.

If STREAM_INFO_DEFINES_PAYLOADSIZE returns true the Data Stream module must provide the buffer describing parameters. This allows the GenTL Producer to modify the buffer parameters to preprocess an image. In case the GenTL Producer is doing that it must implement all buffer describing parameters. For a detailed description please refer to chapter 5.5.

With that information one or multiple buffers can be allocated as the GenTL Consumer needs. The allocation can also be done by the GenTL Producer driver with the combined DSAllocAndAnnounceBuffer function. If the buffers are larger than requested it does not matter and the real size can be obtained through the DSGetBufferInfo function.

If the buffers are smaller than requested the error event is fired on the Buffer module (if the error event is implemented on the Buffer module) and on the transmitting Data Stream module with a GC ERR BUFFER TOO SMALL error code. It is up to the GenTL Producer if a "too small" buffer is filled with parts of the transferred payload or if the buffer is not filled at all. In both cases the buffer should be delivered to the GenTL Consumer if the underlying technology allows it and the according BUFFER INFO CMDs BUFFER INFO IS INCOMPLETE, BUFFER INFO SIZE FILLED and BUFFER INFO DATA LARGER THAN BUFFER should report the fill state. Also in case one or more of the announced buffers are smaller than the payload size the GenTL Producer can refuse to start the acquisition through DSStartAcquisition returning an error code GC ERR BUFFER TOO SMALL.

The payload size for each buffer, no matter if defined by the GenTL Producer or by the remote device, may change during acquisition as long as the acquired payload size delivered is smaller than the actual payloadsize reported at acquisition start. The payload size of a given buffer can be queried through the BUFFER_INFO CMDs.

#### 5.2.2 Announce Buffers

All buffers to be used in the acquisition engine must be made known prior to their use. Buffers can be added (announced) and removed (revoked) at any time. While usually all