|  ![img-53.jpeg](img-53.jpeg) CAM |   | ![img-54.jpeg](img-54.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

Note: to improve interoperability, it is recommended that for device events based on "standard" event data formats, the buffer delivered through EventGetData is directly the buffer that can be fed to the corresponding standard GenApi event adapter. For example, in case of GigE Vision it would be the entire EVENTDATA packet including the header.

If queued event data is not needed anymore the queue can be emptied by calling the EventFlush function on the associated EVENT_HANDLE. To inquire the queue state of a buffer the GenTL Consumer can call DSGetBufferInfo with the info command BUFFER_INFO_IS_QUEUED.

Signals that occur without a corresponding event object being registered using GCRRegisterEvent are silently discarded.

A single event notification carries one event and its data.

For example, a GigE Vision device event sent through the message channel carrying multiple EventIDs in a single packet must result in multiple GenTL Producer events. Each GenTL Producer event will then provide a single GigE Vision EventID.

#### 4.2.4 Example

This sample shows how to register a New Buffer event.

{
    GCRRegisterEvent(hDS, EVENT_NEW_BUFFER, hNewBufferEvent);
    CreateThread(AcqFunction);
}

##### 4.2.4.1 AcqFunction

This sample shows the wait loop to retrieve new buffers.

{
    while (!EndRun)
    {
    EventGetData(hNewBufferEvent, EventData);
    if (successful)
    {
    // Do something with the new buffer
    }
    }
}