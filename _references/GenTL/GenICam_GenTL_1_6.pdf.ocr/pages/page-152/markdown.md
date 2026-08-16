|  GENICAM |   | ![img-216.jpeg](img-216.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

The column labeled “Impl” in the following table lists if the implementation of a given command is mandatory (M), optional (O) or conditional mandatory (CM). Mandatory means that a GenTL Producer must implement the listed command. Optional means that it is up to the implementor if a given command is implemented or not. Conditional Mandatory means that command is to be implemented if possible.

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|  STREAM_INFO_ID | M | 0 | Unique ID of the data stream.Data type: STRING  |
|  STREAM_INFO_NUM_DELIVERED | O | 1 | Number of delivered buffers since last acquisition start.Data type: UINT64  |
|  STREAM_INFO_NUM_UNDERRUN | O | 2 | Number of lost frames due to queue underrun.This number is initialized with zero at the time the stream is opened and incremented every time the data could not be acquired because there was no buffer in the input pool.Data type: UINT64  |
|  STREAM_INFO_NUM_ANNOUNCED | O | 3 | Number of announced buffers.Data type: SIZET  |
|  STREAM_INFO_NUM_QUEUED | O | 4 | Number of buffers in the input pool plus the buffer(s) currently being filled.This does not include the buffers in the output queue. The intention of this informational value is to prevent/early detect an underrun of the acquisition buffers.Data type: SIZET  |
|  STREAM_INFO_NUM_AWAIT_DELIVERY | O | 5 | Number of buffers in the output buffer queue.Data type: SIZET  |
|  STREAM_INFO_NUM_STARTED | O | 6 | Number of frames started in the acquisition engine.This number is incremented every time a new buffer is started to be filled (data written to) regardless if the buffer is later delivered to the user or discarded for any reason. This number is initialized with 0 at at the time of the stream is opened. It is not reset until the stream is closed.Data type: UINT64  |
|  STREAM_INFO_PAYLOAD_SIZE | CM | 7 | Size of the expected data in bytes.Data type: SIZET  |