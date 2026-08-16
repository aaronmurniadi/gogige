|  GEN<I>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

In case a given GenTL Producer does not provide a data stream it must return “0” for the number of available stream channels. In this case a call to DevOpenDataStream and all data stream related functions which start with a DS in the name will fail. This is then called a “Non Streaming Implementation”. It only covers the control layer which is responsible for enumeration and communication with the device.

If a device is not needed anymore call the DevClose function to free the Device module's resources and its depending child Data Streams if they are still open.

After a Device module has been closed it may be opened again and the handle to the module may be different from the first instantiation.

### 3.5 Data Stream

The Data Stream module does not enumerate its child modules. Main purpose of this module is the acquisition which is described in detail in chapter 5 Acquisition Engine (page 42ff). Buffers are introduced by the calling GenTL Consumer and thus it is not necessary to enumerate them.

Every stream is identified not by an index but by a Device module wide unique ID. The content of this ID is up to the GenTL Producer and is only interpreted by it and not by any GenTL Consumer.

When a Data Stream module is not needed anymore the DSClose function must be called to free its resources. This automatically stops a running acquisition, flushes all buffers and revokes them.

Access from a different process space is not recommended. The module does no reference counting. That means that even if a Data Stream module handle is requested a second time from within one process space the second call will return the error GC ERR RESOURCE IN USE. The first call from within that process to the close function will free all resources and shut down the module in that process.

After a Data Stream module has been closed it may be opened again and the handle to the module may be different from the first instantiation.

### 3.6 Buffer

A buffer acts as the destination for the data from the acquisition engine.

Every buffer is identified not by an index but by a unique handle returned from the DSAnnounceBuffer, DSAllocAndAnnounceBuffer or DSAnnounceCompositeBuffer functions.

A buffer can be allocated either by the GenTL Consumer or by the GenTL Producer. Buffers allocated by the GenTL Consumer are made known to the Data Stream module by a call to DSAnnounceBuffer or DSAnnounceCompositeBuffer which returns a BUFFER_HANDLE for this buffer. Buffers allocated by the GenTL Producer are retrieved by a call to DSAllocAndAnnounceBuffer which also returns a BUFFER_HANDLE. The two methods (consumer vs. producer allocated buffers) must not be mixed on a single Data