|  ![img-250.jpeg](img-250.jpeg)CAM |   | ![img-251.jpeg](img-251.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  URL_SCHEME_HTTP | 1 | The XML-file can be retrieved from a webserver using the http protocol. The URL where it can be downloaded can be queried using the GCGetPortURLInfo function with the URL_INFO_URL command.  |
| --- | --- | --- |
|  URL_SCHEME_FILE | 2 | The XML-file can be read from the local hard disk. The filename can be queried through GCGetPortURLInfo function using the URL_INFO_FILENAME command.  |
|  URL_SCHEME_CUSTOM_ID | 1000 | Starting value for custom IDs which are implementation specific.  |

#### 6.4.6 Signaling Enumerations

##### 6.4.6.1 EVENT_DATA_INFO_CMD

enum EVENT_DATA_INFO_CMD

This enumeration defines commands to retrieve information with the EventGetDataInfo function on delivered event data.

The availability and the data type of the enumerators depend on the event type (see below).

The column labeled “Impl” in the following table lists if the implementation of a given command is mandatory (M), optional (O) or conditional mandatory (CM). Mandatory means that a GenTL Producer must implement the listed command even tough it might return NI or NA under certain circumstances. Optional means that it is up to the implementor if a given command is implemented or not. Conditional Mandatory means that command is to be implemented if possible.

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|  EVENT_DATA_ID | M | 0 | Attribute in the event data to identify the object or feature the event refers to. This can be, e.g., the error code for an error event or the feature name for GenApi related events.  |
|  EVENT_DATA_VALUE | M | 1 | Defines additional data to an ID. This can be, e.g., the error message for an error event.  |
|  EVENT_DATA_NUMID | M | 2 | Attribute in the event data to identify the object or feature the event refers to. It is the numeric representation of EVENT_DATA_ID if applicable. In case it is not possible to map EVENT_DATA_ID to a number the EventGetDataInfo function  |