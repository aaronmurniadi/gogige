|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_INVALID_HANDLE

The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream.

GC_ERR_INVALID_PARAMETER

Parameter phDevice is an invalid pointer (NULL or ~0x0).

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.5.15 DSGetNumBufferParts

GC_ERROR DSGetNumBufferParts ( DS_HANDLE hDataStream,
    BUFFER_HANDLE hBuffer,
    uint32_t * piNumParts )

Inquires the number of independent data parts in the buffer. The GenTL Producer may return 0 in *piNumParts in case the buffer payload is multipart but does not contain any parts. For example in case the individual parts of a multipart buffer can be enabled or disabled in the camera it can happen that a camera sends a multipart payload with no parts enabled and maybe only chunk data is being delivered. As described in this case *piNumParts would report 0. Detailed information about the individual parts can be queried using function DSGetBufferPartInfo.

If the buffer content can be fully described using the information available through DSGetBufferInfo queries, it is not split into parts and the buffer payload is not multi-part the GenTL Producer must return the error GC_ERR_NO_DATA. The GenTL Consumer would in this case avoid querying information about buffer parts using DSGetBufferPartInfo.

If the reported payload is multi-part the GenTL Producer must use DSGetNumBufferParts and DSGetBufferPartInfo to provide information about the buffer.

##### Parameters

[in] hDataStream
[in] hBuffer
[out] piNumParts

Data Stream module to work on.

Buffer handle to retrieve information about.

Number of independent data parts in the buffer. The reported number may be 0 in case the referenced buffer carries a multipart buffer payload but for some reason the number of parts is buffer is 0.

##### Returns

GC_ERR_SUCCESS:

Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED:

No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE:

The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through