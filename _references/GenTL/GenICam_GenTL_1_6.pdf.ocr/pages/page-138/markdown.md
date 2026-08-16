|  ![img-192.jpeg](img-192.jpeg)CAN |   | ![img-193.jpeg](img-193.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 6.4.4.2 ACQ_START_FLAGS

enum ACQ_START_FLAGS

This enumeration defines special start flags for the acquisition engine. The function used is DSStartAcquisition.

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  ACQ_START_FLAGS_DEFAULT | 0 | Default behavior.  |
|  ACQ_START_FLAGS_CUSTOM_ID | 1000 | Starting value for GenTL Producer custom IDs.  |

##### 6.4.4.3 ACQ_STOP_FLAGS

enum ACQ_STOP_FLAGS

This enumeration defines special stop flags for the acquisition engine. The function used is DSStopAcquisition.

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  ACQ_STOP_FLAGS_DEFAULT | 0 | Stops the acquisition engine when the currently running tasks like filling a buffer are completed (default behavior).  |
|  ACQ_STOP_FLAGS_KILL | 1 | Stop the acquisition engine immediately. In case this results in a partially filled buffer the Producer will return the buffer through the regular mechanism to the user, indicating through the info function of that buffer that this buffer is not complete.  |
|  ACQ_STOP_FLAGS_CUSTOM_ID | 1000 | Starting value for GenTL Producer custom IDs which are implementation specific.If a generic GenTL Consumer is using custom ACQ_STOP_FLAGS provided through a specific GenTL Producer implementation it must differentiate the handling of different GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |

##### 6.4.4.4 BUFFER_INFO_CMD

enum BUFFER_INFO_CMD

This enumeration defines commands to retrieve information with the DSGetBufferInfo function on a buffer handle. In case a BUFFER_INFO_CMD is not available or not