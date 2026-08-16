|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

# Parameters

[in] hDataStream

Data Stream module to work on.

[in] iOperation

Flush operation to perform as defined in

ACQ QUEUE TYPE.

# Returns

GC_ERR_SUCCESS

Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED

No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE

The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream.

GC_ERR_NOT_IMPLEMENTED

iOperation is not implemented.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

#### 6.3.5.6 DSGetBufferID

GC ERROR DSGetBufferID

( DS_HANDLE hDataStream,
uint32_t iIndex,
BUFFER_HANDLE * phBuffer )

DSGetBufferID queries the buffer handle for a given index iIndex. The buffer handle phBuffer works as a unique ID of an instance of the Buffer module. The relation between an index iIndex and a particular buffer stays valid until a buffer revoked. The index reflects the order in which buffers are announced. If new buffers are announced they are to be appended at the end. If buffers “in the middle” are revoked the sequentially following buffers move into that position. The index stays continuous. So for example if you have 10 buffers announced and you remove the buffer with the id of index 5 you still have the index range from 0 to 8.

Note that the relation between index and buffer handle might change with revoked buffers. As long as no buffers are revoked this relation must not change.

The number of announced buffers can be queried with the DSGetInfo function.

# Parameters

[in] hDataStream

Data Stream module to work on.

[in] iIndex

Zero-based index of the buffer on this data stream.

[in,out] phBuffer

Buffer module handle of the given iIndex.

# Returns

GC_ERR_SUCCESS

Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED

No preceding call to GCInitLib.