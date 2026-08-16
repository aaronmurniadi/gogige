|  ![img-127.jpeg](img-127.jpeg) CAM |   | ![img-128.jpeg](img-128.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

# Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream.  |
|  GC_ERR_INVALID_PARAMETER | Parameters pBuffer and/or phBuffer are invalid pointers (NULL or ~0x0).  |
|  GC_ERR_RESOURCE_IN_USE | The specified pBuffer is already announced to this Data Stream module or, depending on the implementation of the GenTL Producer, it has already been announced to another instance of the Data Stream module (see chapter 3.6).  |
|  GC_ERR_BUSY | The acquisition has been started and the GenTL Producer does not support announcing buffers while the acquisition is active.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.5.3 DSAnnounceCompositeBuffer

|  GC_ERROR | DSAnnounceCompositeBuffer ( DS_HANDLE hDataStream, size_t iNumSegments void ** ppSegments, size_t * piSizes, void * pPrivate, BUFFER_HANDLE * phBuffer )  |
| --- | --- |

This announces GenTL Consumer allocated memory to the Data Stream associated with the hDataStream handle and returns a buffer handle which references that composite buffer until the buffer is revoked. This will allocate internal resources which will be freed upon a call to DSRevokeBuffer.

In contrast to DSAnnounceBuffer, this function allows to announce a buffer not referring a single contiguous block of memory, but consisting of multiple segments. This allows routing logically independent portions of the acquired data to separate memory locations, for example when streaming data over multiple flows mechanism. The segments may correspond with different memory areas as well as form a contiguous block, but must not overlap.

All the segments of the announced composite buffer are treated as a single entity by all functions operating on a buffer handle.