|  ![img-43.jpeg](img-43.jpeg)CAN |   | ![img-44.jpeg](img-44.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

By using an event-object-based approach the acquisition engine for example only prepares the necessary data and then signals its availability to the GenTL Consumer through the previously registered event objects. The GenTL Consumer can decide in which thread context and with which priority the data processing is done. Thus processing of the event and the signal's generation are decoupled.

#### 4.2.1 Event Objects

Event objects allow asynchronous signaling to the calling GenTL Consumer.

Event objects have two states: signaled or not signaled. An EventGetData function blocks the calling thread until either a user defined timeout occurs, the event object is signaled or the wait is terminated by the GenTL Consumer. If the event object is signaled prior to the call of the EventGetData functions, the function returns immediately delivering the data associated with the event signaled.

Not every event type can be registered with every module and not every module needs to implement every possible event type. If a module is not listed for an event (in the table below), it does not support that event type.

The maximum size of the data delivered by an event is defined in the event description and can be retrieved through the EventGetInfo function. The actual size is returned by the EventGetData function, which retrieves the data associated with the event.

There are no mandatory event types. If an event type is not implemented in a GenTL Producer the GCRegisterEvent should return GC ERR NOT IMPLEMENTED. If an event type is implemented by a GenTL Producer module it is recommended to register an event object for that event type. The event types are described in the following table.

Table 4-2: Event types per module

|  Event Type | Modules | Description  |
| --- | --- | --- |
|  Error | All | A GenTL Consumer can get notified on asynchronous errors in a module. These are not errors due to function calls in the C interface or in the GenApi Feature access. These have their own error reporting. For example this event applies to an error while data is acquired in the acquisition engine of a Data Stream module.  |
|  New Buffer | Data Stream | New data is present in a buffer in the acquisition engine. In case the New Buffer event is implemented it must be registered on a Data Stream module. After registration the calling GenTL Consumer is informed about every new buffer in that stream. If the EventFlush function is called all buffers in the output buffer queue are discarded. If a DSFlushQueue is called all events from the event queue are removed as well. Please use the BUFFER_INFO_IS_QUEUED info command in order to inquire the queue state of a buffer.  |