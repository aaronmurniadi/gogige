|  ![img-188.jpeg](img-188.jpeg)CAM |   | ![img-189.jpeg](img-189.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|  DEVICE_INFO_VERSION | O | 8 | Device version in string format. If the information is not known, the query should result in GC_ERR_NOT_AVAILABLE. Data type: STRING  |
|  DEVICE_INFO_TIMESTAMP_FREQUENCY | O | 9 | Tick frequency of the device's timestamp counter in ticks per second. The counter is used for example to assign timestamps to the individual acquired buffers (BUFFER_INFO_TIMESTAMP). Data type: UINT64  |
|  DEVICE_INFO_CUSTOM_ID | O | 1000 | Starting value for GenTL Producer custom IDs which are implementation specific. If a generic GenTL Consumer is using custom DEVICE_INFO_CMDs provided through a specific GenTL Producer implementation it must differentiate the handling of different GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |

#### 6.4.4 Data Stream Enumerations

##### 6.4.4.1 ACQ_QUEUE_TYPE

enum ACQ_QUEUE_TYPE

This enumeration commands define from which to which queue/pool buffers are flushed with the DSFlushQueue function.

|  Enumerator | Value | Description  |
| --- | --- | --- |