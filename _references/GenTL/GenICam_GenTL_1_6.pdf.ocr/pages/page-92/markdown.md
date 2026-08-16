|  ![img-129.jpeg](img-129.jpeg) CAM |   | ![img-130.jpeg](img-130.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

Announcing a buffer to a data stream does not mean that this buffer will be automatically queued for acquisition. This is done through a separate call to DSQueueBuffer.

The memory referenced in ppSegments must stay valid until the buffer is revoked with DSRevokeBuffer. Every call of this function must be matched with a call of DSRevokeBuffer.

A composite buffer (and any of its segments) can only be announced once to a given stream. If a GenTL Consumer tries to announce an already announced buffer the function will return the error GC_ERR_RESOURCE_IN_USE. A buffer (with its segments) may additionally be announced to one or more other data stream(s) which will then result in one or more additional handles. The GenTL Consumer needs to take care about synchronisation between these streams.

Refer to chapter 5.7 in order to determine the buffer structure.

Note: there is no alloc-and-announce version of this function.

Parameters

|  [in] | hDataStream | Data Stream module to work on.  |
| --- | --- | --- |
|  [in] | iNumSegments | Number of segments constituting the composite buffer.  |
|  [in] | ppSegments | Pointers to memory of individual segments of the composite buffer to announce. The array must contain iNumSegments items.  |
|  [in] | piSizes | Size of the segments in bytes. The array must contain iNumSegments items, one for each segment in ppSegments.  |
|  [in] | pPrivate | Pointer to private data which will be passed to the GenTL Consumer on New Buffer events. This parameter may be NULL.  |
|  [out] | phBuffer | Buffer module handle of the newly announced composite buffer. It is recommended to initialize *phBuffer to GENTL_INVALID_HANDLE before calling DSAnnounceCompositeBuffer to indicate an invalid handle.  |

Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream.  |
|  GC_ERR_NOT_IMPLEMENTED: | The GenTL implementation does not support composite buffers.  |
|  GC_ERR_INVALID_PARAMETER | Parameters ppSegments and/or piSizes and/or phBuffer are invalid pointers or any of the segment pointers in  |