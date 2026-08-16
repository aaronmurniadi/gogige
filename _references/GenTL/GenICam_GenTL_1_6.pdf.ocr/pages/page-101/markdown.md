|  ![img-141.jpeg](img-141.jpeg) CAM |   | ![img-142.jpeg](img-142.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_INVALID_HANDLE

The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream.

GC_ERR_NOT_IMPLEMENTED

One or more flags set in iStopFlags referencing functionality which is not implemented.

GC_ERR_RESOURCE_IN_USE

The Acquisition has already been terminated or it has not been started.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.5.13 DSGetBufferChunkData

GC_ERROR DSGetBufferChunkData( DS_HANDLE hDataStream,
BUFFER_HANDLE hBuffer,
SINGLE_CHUNK_DATA * pChunkData,
size_t * piNumChunks )

DSGetBufferChunkData parses the transport layer technology dependent chunk data info in the buffer. The layout of the chunk data present in the buffer is returned in the pChunkData array, one entry per chunk. Every single chunk is described using its ChunkID, offset in the buffer and chunk data size.

Note that for composite buffers, the individual chunk offsets reported in pChunkData are linear offsets within entire payload (see also 5.4.1.1).

When dealing with buffers containing standard self-described containers (such as GenDC), the GenTL Consumer must parse and interpret the chunk data directly according to the given container type specification without using this function.

##### Parameters

[in] hDataStream

Data Stream module to work on.

[in] hBuffer

Buffer handle to parse.

[out] pChunkData

GenTL Consumer allocated array of structures to receive the chunk layout information. If this parameter is NULL, piNumChunks will contain the number of chunks in the buffer, e.g., the minimal number of entries in the pChunkData array.

[in,out] piNumChunks

pChunkData equal NULL:

out: number of chunks in the buffer (minimal number of entries in the pChunkData array to hold all information).

pChunkData unequal NULL:

in: number of entries in the provided pChunkData array.

out: number of entries successfully written to the pChunkData array.