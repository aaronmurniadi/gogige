|  ![img-190.jpeg](img-190.jpeg) CAM |   | ![img-191.jpeg](img-191.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  ACQ_QUEUE_INPUT_TO_OUTPUT | 0 | Flushes the buffers from the input pool to the output buffer queue and if necessary adds entries in the “New Buffer” event data queue. The buffers currently being filled are not affected by this operation.This only applies to the mandatory default buffer handling mode. Whether the buffer contains real data or is result of a flush operation can be inquired through the buffer info command BUFFER_INFO_NEW_DATA. This allows the GenTL Consumer to maintain all buffers without a second reference in the GenTL Consumer because all buffers are delivered through the new buffer event.  |
|  ACQ_QUEUE_OUTPUT_DISCARD | 1 | Discards all buffers in the output buffer queue and if necessary remove the entries from the event data queue.  |
|  ACQ_QUEUE_ALL_TO_INPUT | 2 | Puts all buffers in the input pool. This is including those in the output buffer queue and the ones which are currently being filled and discard entries in the event data queue.  |
|  ACQ_QUEUE_UNQUEUED_TO_INPUT | 3 | Puts all buffers that are neither in the input pool nor being currently filled nor in the output buffer queue in the input pool.  |
|  ACQ_QUEUE_ALL_DISCARD | 4 | Discards all buffers in the input pool and the buffers in the output queue including buffers currently being filled so that no buffer is bound to any internal mechanism and all buffers may be revoked or requeued.  |
|  ACQ_QUEUE_CUSTOM_ID | 1000 | Starting value for GenTL Producer custom IDs which are implementation specific.If a generic GenTL Consumer is using custom ACQ_QUEUE_TYPES provided through a specific GenTL Producer implementation it must differentiate the handling of different GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |