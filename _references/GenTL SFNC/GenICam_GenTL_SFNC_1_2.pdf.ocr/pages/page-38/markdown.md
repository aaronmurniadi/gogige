|  BufferData | O | All | IRegister | R/(W) | - | E | Entire buffer data.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  BufferTimeStamp | O | All | IInteger | R | - | E | Timestamp the buffer was acquired.  |
|  BufferNewData | O | All | IBoolean | R | - | E | Flag to indicate that the buffer contains new data since the last delivery.  |
|  BufferIsQueued | O | All | IBoolean | R | - | E | Flag to indicate if the buffer is in the input pool or output buffer queue.  |
|  BufferIsAcquiring | O | All | IBoolean | R | - | E | Flag to indicate that the buffer is currently being filled with data.  |
|  BufferIsIncomplete | O | All | IBoolean | R | - | E | Flag to indicate that a buffer was filled but an error occurred during that process.  |
|  BufferPayloadType | O | All | IEnumeration | R | - | E | Payload type of the data.  |
|  BufferNumberOfParts | O | All | IInteger | R | - | E | The number of parts in the current buffer as delivered by the transport mechanism.  |
|  BufferPartSelector | O | All | IInteger | R | - | E | The buffer part to extract information from.  |
|  BufferSizeFilled | O | All | IInteger | R | Byte | E | Number of bytes written into the buffer last time it was filled.  |
|  BufferPartDataType[BufferPartSelector] | O | All | IEnumeration | R | - | E | The data type of the part.  |
|  BufferPartSourceIDValue[BufferPartSelector] | O | All | IInteger | R | - | E | The Source ID type of the part.  |
|  BufferPartRegionIDValue[BufferPartSelector] | O | All | IInteger | R | - | E | The Region ID type of the part.  |
|  BufferPartComponentIDValue[BufferPartSelector] | O | All | IInteger | R | - | E | The Component ID type of the part.  |
|  BufferWidth[BufferPartSelector] | O | All | IInteger | R | - | E | Width of the data in the buffer in number of pixels.  |
|  BufferHeight[BufferPartSelector] | O | All | IInteger | R | - | E | Height of the data in the buffer in number of pixels as configured.  |