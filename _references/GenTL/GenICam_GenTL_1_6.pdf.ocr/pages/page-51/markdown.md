|  ![img-68.jpeg](img-68.jpeg) CAM |   | ![img-69.jpeg](img-69.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

if ( PayloadType == PAYLOAD_TYPE_CHUNK_DATA )
{
    ChunkListSize = 0;
    DSGetBufferChunkData( hStream, hBuffer, 0, ChunkListSize )
    {
    // Alternatively it would be possible to inquire the max number of
    // chunks per buffer through STREAM_INFO_NUM_CHUNKS_MAX

    DSGetInfo( hStream, STREAM_INFO_NUM_CHUNKS_MAX, Type, ChunkListSize, sizeof(ChunkListSize));

    // In this case the consumer needs error checking in case the
    // GenTL Producer cannot provide that information
    }

    // Allocate array of SINGLE_CHUNK_DATA structures
    DSGetBufferChunkData( hStream, hBuffer, ChunkArray, ChunkListSize )

    // Pass Chunk Array to GenApi Port
    // Free ChunkArray.
}

### 5.5 Data Payload Delivery

The GenTL Producer is allowed to modify the image data acquired from the remote device if needed or if it is convenient for the user. An examples of such modifications can be a PixelFormat conversion (e.g., when decoding a Bayer encoded color image) or LinePitch adjustment (elimination of the line padding produced on the remote device).

Whenever a modification leads to a change of basic parameters required to interpret the image, the GenTL Producer must inform the GenTL Consumer about the modifications. It is mandatory to report the modified values through the BUFFER_INFO_CMD or BUFFER_PART_INFO_CMD commands of the C interface. The tables listing the values for BUFFER_INFO_CMD and BUFFER_PART_INFO_CMD also list which commands are optional and and which are mandatory.

If a given BUFFER_INFO_CMD command is not available, the GenTL Consumer assumes, that the GenTL Producer did not modify the corresponding parameter and that it corresponds to the settings on the remote device. For example, if the query for BUFFER_INFO_PIXELFORMAT returns an error, meaning that the BUFFER_INFO_PIXELFORMAT command is not available, the GenTL Consumer should assume that the GenTL Producer did