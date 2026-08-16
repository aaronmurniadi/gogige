|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

#### 5.2.6 Acquire Image Data

The following action is performed in a loop:

- Wait for the “New Buffer” event to be signaled (see 4.2 Signaling page 35ff)
- Process image data
- Requeue buffer in the Input Buffer Pool

With the event data from the signaled event the newly filled buffer can be obtained and then processed. As stated before no assumptions on the order of the buffers are made except if the buffer handling mode defines it.

Requeuing the buffers can be done in any order using the DSQueueBuffer function. As long as the buffer is not in the Input Buffer Pool or in the Output Buffer Queue the acquisition engine will not write into the buffer. This mechanism locks the buffer effectively.

#### 5.2.7 Stop Acquisition

When finished acquiring image data the acquisition on the remote device is to be stopped if necessary. This can be done by setting the “AcquisitionStop” standard feature on the remote device. If it is present the command should be executed. Afterwards the DSStopAcquisition function is called to stop the acquisition on the host. By doing that the status of the buffers does not change. That implies that a buffer that is in the Input Buffer Pool remains there. The same is true for buffers in the Output Buffer Queue. This has the advantage that buffers which were filled during the acquisition stop process still can be retrieved and processed. If ACQ STOP FLAGS KILL is specified in the call to DSStopAcquisition a partially filled buffer is by default moved to the output buffer queue for processing. DSGetBufferInfo with BUFFER INFO IS INCOMPLETE would indicate that the buffer is not complete.

If a device implements the SFNC Transfer Control features, the GenTL Consumer may need to stop the transfer on the remote device, depending on the operating mode.

#### 5.2.8 Flush Buffer Pools and Queues

In order to clear the state of the buffers to allow revoking them, the buffers have to be flushed either with the DSFlushQueue function or with the EventFlush function. With the DSFlushQueue function buffers from the Input Buffer Pool can either be flushed to the Output Buffer Queue or discarded. Buffers from the Output Buffer Queue also must either be processed or flushed. Flushing the Output Buffer Queue is done by calling EventFlush function. Using the EventFlush function on the "New Buffer" event discards the buffers from the Output Buffer Queue.

#### 5.2.9 Revoke Buffers

To enable the acquisition engine to free all resources needed for acquiring image data, revoke the announced buffers. Buffers that are referenced in either the Input Buffer Pool or the Output Buffer Queue cannot be revoked. After revoking a buffer with the DSRevokeBuffer function it is not known to the acquisition engine and thus can neither be queued nor receive any image data.