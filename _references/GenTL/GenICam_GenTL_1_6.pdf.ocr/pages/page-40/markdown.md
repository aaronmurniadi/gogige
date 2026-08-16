|  ![img-51.jpeg](img-51.jpeg) |   | ![img-52.jpeg](img-52.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

- Register a DeviceEvent on the corresponding GenTL module.
- Inquire the max needed buffer size.
- Allocate the buffer to receive the event data.
- Wait for the event and data. The structure of the data in the provided buffer is not defined and GenTL Producer dependent. The only exception to that would be the New Buffer event which provides a defined internal struct.
- Extract the data in the buffer using EventGetDataInfo. This step is not necessary in cases when the GenTL Consumer knows the contents of the buffer delivered through EventGetData, such as a New Buffer event.
• Data processing/usage.
- Unregister event.
- Deallocate buffer.

As described the content of the buffer retrieved through EventGetData is GenTL Producer implementation specific and may be parsed using the EventGetDataInfo function. The only exception to that is the New Buffer event which will return the EVENT NEW BUFFER DATA structure.

For the Device Event events (EVENT_REMOTE_DEVICE) the GenTL Producer must provide two types of information about every single event, so that it can be "connected" to the remote device's nodemap:

- Event ID: queried through EventGetDataInfo(EVENT_DATA_ID). The ID is passed as a string representation of hexadecimal form, for example "CF51" without the leading '0x'. The ID can be also queried directly in numeric form using EventGetDataInfo (EVENT_DATA_NUMID).
- Event data: buffer containing the (optional) data accompanying the event. It must correspond with the data addressable from the remote device nodemap, the beginning of the buffer must correspond with address 0 of the nodemap's event port. For example, for GigE Vision devices this is by convention the entire EVENTDATA packet, without the 8-byte GVCP header.

Also for the module's events (EVENT_MODULE) the GenTL Producer must provide two types of information about every single event, so that it can be connected to a module's nodemap:

- Event ID: queried through EventGetDataInfo(EVENT_DATA_ID). The ID is passed as a string representation of hexadecimal form, for example "CF51" without the leading '0x'. The ID can be also queried directly in numeric form using EventGetDataInfo (EVENT_DATA_NUMID).
- Event data: buffer containing the (optional) data accompanying the event. It must correspond with the data addressable from the module's nodemap, the beginning of the buffer must correspond with address 0 of the nodemap's event port, similar to way the EVENT_REMOTE_DEVICE is working.