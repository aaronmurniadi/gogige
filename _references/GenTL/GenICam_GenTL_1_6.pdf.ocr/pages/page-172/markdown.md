|  ![img-254.jpeg](img-254.jpeg) CAM |   | ![img-255.jpeg](img-255.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|   |  |  | is not known a priori by the GenTL Producer the EventGetInfo function returns GC_ERR_NOT_AVAILABLE. This maximum size must not change during runtime. Data type: SIZET  |
|  EVENT_INFO_DATA_SIZE_MAX | M | 4 | Maximum size in bytes of the information output buffer of EventGetDataInfo function for EVENT_DATA_VALUE. In case this is not known a priori by the GenTL Producer the EventGetDataInfo function returns the GC_ERR_NOT_AVAILABLE error. This maximum size must not change during runtime. Data type: SIZET  |
|  EVENT_INFO_CUSTOM_ID | O | 1000 | Starting value for GenTL Producer custom IDs which are implementation specific. If a generic GenTL Consumer is using custom EVENT_INFO_CMDs provided through a specific GenTL Producer implementation it must differentiate the handling of different GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |

##### 6.4.6.3 EVENT_TYPE

enum EVENT_TYPE

Known event types that can be registered on certain modules with the GCRegisterEvent function. See chapter 4.2 Signaling page 35 for more information.

Specific values of the event data can be queried with the EventGetDataInfo function. It is stated in the table which enumerators specify values that can be retrieved by a specific event type.

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  EVENT_ERROR | 0 | Notification on module errors. Values that can be retrieved are:  |