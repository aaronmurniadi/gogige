![img-8.jpeg](img-8.jpeg)

|  BufferXOffset[BufferPartSelector] | O | All | IInteger | R | - | E | XOffset of the data in the buffer in number of pixels from the image origin to handle areas of interest.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  BufferYOffset[BufferPartSelector] | O | All | IInteger | R | - | E | YOffset of the data in the buffer in number of lines from the image origin to handle areas of interest.  |
|  BufferXPadding[BufferPartSelector] | O | All | IInteger | R | Byte | E | XPadding of the data in the buffer in number of bytes.  |
|  BufferYPadding | O | All | IInteger | R | Byte | E | YPadding of the data in the buffer in number of bytes.  |
|  BufferFrameID | R | All | IInteger | R | - | E | A sequentially incremented number of the frame.  |
|  BufferImagePresent | O | All | IBoolean | R | - | E | Flag to indicate if the current data in the buffer contains image data.  |
|  BufferImageOffset | O | All | IInteger | R | Byte | E | Offset of the image data from the beginning of the delivered buffer in bytes.  |
|  BufferPixelFormat[BufferPartSelector] | O | All | IEnumeration | R | - | E | Format of the pixels provided by the buffer.  |
|  BufferDeliveredImageHeight[BufferPart Selector] | O | All | IInteger | R | - | E | The number of lines in the current buffer part as delivered by the transport mechanism.  |
|  BufferDeliveredChunkPayloadSize | O | All | IInteger | R | - | E | Size of the valid chunk payload data delivered in the buffer.  |
|  BufferChunkLayoutID | O | All | IInteger | R | - | E | ID of the chunk data layout delivered in the buffer.  |
|  BufferFileName | O | All | IString | R | - | E | Filename for the file payload data delivered in the buffer.  |

### 2.5.3 GenICam Control

Contains the features related to GenICam control and access of a specific Buffer module.

Table 2-22: GenICam Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  Root | O | All | ICategory | R | - | B | Provides the Root of the GenICam features tree.  |