|  ![img-241.jpeg](img-241.jpeg) CAM |   | ![img-242.jpeg](img-242.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|   |  |  | is placed into the Input Buffer Pool. Used similar way as BUFFER_INFO_SIZE_FILLED at buffer level.For technologies or use cases where this is difficult to track, it is valid leave the command not implemented.Data type: SIZET  |
|  SEGMENT_INFO_DATA_SIZE | O | 4 | Size of the data intended to be written to the buffer last time it has been filled. This value is reset to 0 when the buffer is placed into the Input Buffer Pool. Used similar way as BUFFER_INFO_DATA_SIZE at buffer level.For technologies or use cases where this is difficult to track, it is valid leave the command not implemented.Data type: SIZET  |
|  SEGMENT_INFO_CUSTOM_ID | O | 1000 | Starting value for GenTL Producer custom IDs which are implementation specific.If a generic GenTL Consumer is using custom SEGMENT_INFO_CMDs provided through a specific GenTL Producer implementation it must differentiate the handling of different GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |

#### 6.4.5 Port Enumerations

##### 6.4.5.1 PORT_INFO_CMD

enum PORT_INFO_CMD

This enumeration defines commands to retrieve information with the GCGetPortInfo function on a module or remote device handle.

The column labeled “Impl” in the following table lists if the implementation of a given command is mandatory (M), optional (O) or conditional mandatory (CM). Mandatory means that a GenTL Producer must implement the listed command even tough it might return NI or NA under certain circumstances. Optional means that it is up to the implementor if a given