|  ![img-252.jpeg](img-252.jpeg) CAM |   | ![img-253.jpeg](img-253.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|   |  |  | returns GC_ERR_NOT_AVAILABLE. Data type: UINT64  |
|  EVENT_DATA_CUSTOM_ID | O | 1000 | Starting value for GenTL Producer custom IDs which are implementation specific. If a generic GenTL Consumer is using custom EVENT_DATA_INFO_CMDs provided through a specific GenTL Producer implementation it must differentiate the handling of different GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |

##### 6.4.6.2 EVENT_INFO_CMD

enum EVENT_INFO_CMD

This enumeration defines command to retrieve information with the EventGetInfo function on an event handle.

The column labeled “Impl” in the following table lists if the implementation of a given command is mandatory (M), optional (O) or conditional mandatory (CM). Mandatory means that a GenTL Producer must implement the listed command even tough it might return NI or NA under certain circumstances. Optional means that it is up to the implementor if a given command is implemented or not. Conditional Mandatory means that command is to be implemented if possible.

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|  EVENT_EVENT_TYPE | M | 0 | The event type of the event handle. Data type: INT32 (EVENT_TYPE enum value).  |
|  EVENT_NUM_IN_QUEUE | M | 1 | Number of events in the event data queue. Data type: SIZET  |
|  EVENT_NUM_FIRED | O | 2 | Number of events that were fired since the registration of the event through a call to GCRegisterEvent. A fired event is either still in the internal queue or already delivered to the user or discarded through EventFlush. Data type: UINT64  |
|  EVENT_SIZE_MAX | M | 3 | Maximum size in bytes of the event data provided by the event. In case this  |