|  ![img-66.jpeg](img-66.jpeg)CAN |   | ![img-67.jpeg](img-67.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

SINGLE CHUNK DATA structures. This information is sufficient to connect the chunk to the remote device's nodemap (for example through the generic chunk adapter of GenApi reference implementation).

The acquired buffer might contain only the chunk data or the data might be mixed within the same buffer with an image or other data. To query, if a given buffer contains chunk data, the BUFFER_INFO_CONTAINS_CHUNKDATA command may be used which will return true in case the buffer contains chunk data or the function DSGetBufferChunkData can be queried which, in case the buffer contains accessible chunk data, would return the number of chunks available.

There are other chunk data related buffer info commands, such as BUFFER_INFO_IMAGEPRESENT (indicating that the buffer contains also an image) or BUFFER_INFO_CHUNKLAYOUTID (can help to check, if the chunk structure has changed since the last delivered buffer and if it is necessary to parse it again). The STREAM_INFO_NUM_CHUNKS_MAX command reports the maximum number of chunks to be expected in a buffer acquired through a given stream (if that maximum is known a priori).

If the GenTL Consumer knows the chunk data structure, such as accessing a device of a known standard technology, it is not necessary to use the DSGetBufferChunkData function to parse the buffer. The GenTL Consumer can use a more direct approach to extract the data (by using a standard chunk adapter in GenApi reference implementation).

##### 5.4.1.1 Chunk data in composite buffer

The DSGetBufferChunkData function was introduced in GenTL with traditional contiguous buffers in mind, before introduction of composite buffers (announced using DSAnnounceCompositeBuffer, 5.7.1). Because it treats the buffer as a single entity, it does not report the chunk offsets corresponding to a particular buffer segment, instead it reports linear offsets within the entire payload. The GenTL Consumer has to take this into account. If the format of the chunk data within the composite buffer is well known to the GenTL Consumer, it must parse the chunk data directly without help of DSGetBufferChunkData, the GenTL Producer must return an error from that function.

##### 5.4.1.2 Chunk data in a generic container (for example GenDC)

When the acquired payload contains data in a standard self-described container format (such as for example GenDC), the consumer must parse and interpret the chunk data directly according to the given container type specification, without help of the GenTL interface. GenTL Producer must return error (GC_ERR_NO_DATA) for all chunk data related queries, including DSGetBufferChunkData.

#### 5.4.2 Example

This sample shows how to retrieve chunk data from a buffer.

{
    // Check if the buffer contains chunk data
    DSGetBufferInfo (hStream, hBuffer, BUFFER_INFO_PAYLOADTYPE, Type, PayloadType, SizeOfPayloadType);