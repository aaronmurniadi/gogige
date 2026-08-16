|  ![img-131.jpeg](img-131.jpeg) CAM |   | ![img-132.jpeg](img-132.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_RESOURCE_IN_USE

the ppSegments array is invalid pointer (NULL or ~0x0) or iNumSegments is zero.

One or more of the specified ppSegments is already announced to this Data Stream module or, depending on the implementation of the GenTL Producer, it has already been announced to another instance of the Data Stream module (see chapter 3.6).

GC_ERR_BUSY

The acquisition has been started and the GenTL Producer does not support announcing buffers while the acquisition is active.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.5.4 DSClose

GC ERROR DSClose

( DS HANDLE

hDataStream )

Closes the Data Stream module associated with the given hDataStream handle. This frees all resources of the Data Stream module, discards and revokes all buffers.

##### Parameters

[in] hDataStream

Data Stream module handle to close.

##### Returns

GC_ERR_SUCCESS

Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED

No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE

The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.5.5 DSFlushQueue

GC ERROR DSFlushQueue

( DS HANDLE

hDataStream,

ACQ_QUEUE_TYPE

iOperation )

Flushes the by iOperation defined internal buffer pool or queue to another one as defined in ACQ_QUEUE_TYPE.