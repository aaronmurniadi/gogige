|  ![img-150.jpeg](img-150.jpeg) CAM |   | ![img-151.jpeg](img-151.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 6.3.5.21 DSGetNumBufferSegments

GC_ERROR DSGetNumBufferSegments ( DS_HANDLE hDataStream, BUFFER_HANDLE hBuffer, uint32_t * piNumSegments )

Inquires the number of segments in a buffer.

If hBuffer references a composite buffer (announced using DSAnnounceCompositeBuffer), the function reports the number of “real” segments, as announced in DSAnnounceCompositeBuffer.

If hBuffer references a contiguous buffer (announced using DSAllocAndAnnounceBuffer or DSAnnounceBuffer), the function reports the number of “virtual” segments created within the buffer by the acquisition engine when it was last time filled. When flows were used for the acquisition, the virtual segments correspond to the flow structure  \( (5.7.2) \) , otherwise the buffer is assumed to consist of a single segment.

##### Parameters

[in] hDataStream Data Stream module to work on.
[in] hBuffer Buffer handle to retrieve information about.
[out] piNumSegments Number of segments the composite buffer contains. The reported number is expected to be non-zero.

##### Returns

GC_ERR_SUCCESS: Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED: No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE: The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream or the handle hBuffer is invalid (NULL) or does not reference an announced Buffer.

GC_ERR_NOT_IMPLEMENTED: The GenTL implementation does not support composite buffers.

GC_ERR_INVALID_PARAMETER: Parameter piNumSegments is an invalid pointer (NULL or ~0x0).

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.