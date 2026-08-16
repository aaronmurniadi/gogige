|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|   | TransferEnd BufferTooSmall BuffersDiscarded  |
| --- | --- |

Selects which Event to signal to the host application.

Possible values are:

- **NewBufferData**: A new buffer is available.
- **TransferEnd**: The transfer of a data for new buffer finished; this is not directly related with delivering the buffer, the data might be appended to end of Output Buffer Queue, dropped, etc., depending on the buffer handling mode and acquisition engine status.
- **BufferTooSmall**: The buffer was too small to receive the expected amount of data.
- **BuffersDiscarded**: Buffers discarded by GenTL or device. This event could optionally carry two numeric child data fields EventBuffersDiscardedDeviceCount and EventBuffersDiscardedProducerCount.

**EventBuffersDiscardedDeviceCount**: Number of buffers discarded by the device since last fired instance of this event (the producer would get to know about this for example by observing a gap in the block_id sequence)

**EventBuffersDiscardedProducerCount**: Number of buffers discarded by the producer since last fired instance of this event (this would happen e.g. if there are no free buffers available or if given buffer handling mode requires discarding old buffers etc.)

### 3.4.5.3 EventNotification

|  Name | EventNotification[EventSelector]  |
| --- | --- |
|  Category | EventControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |