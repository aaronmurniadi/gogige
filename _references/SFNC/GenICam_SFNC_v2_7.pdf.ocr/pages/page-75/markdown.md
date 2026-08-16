|  ChunkOffsetY | R | Integer | R | - | E | Returns the OffsetY of the image included in the payload.  |
| --- | --- | --- | --- | --- | --- | --- |
|  ChunkWidth[ChunkRegionSelector] | R | Integer | R | - | E | Returns the Width of the image included in the payload.  |
|  ChunkHeight[ChunkRegionSelector] | R | Integer | R | - | E | Returns the Height of the image included in the payload.  |
|  ChunkPixelFormat | R | Enumeration | R | - | E | Returns the PixelFormat of the image included in the payload.  |
|  ChunkPixelDynamicRangeMin | R | Integer | R | - | E | Returns the minimum value of dynamic range of the image included in the payload.  |
|  ChunkPixelDynamicRangeMax | R | Integer | R | - | E | Returns the maximum value of dynamic range of the image included in the payload.  |
|  ChunkBinningHorizontal | O | Integer | R/W | - | E | Number of horizontal photo-sensitive cells combined together.  |
|  ChunkBinningVertical | O | Integer | R/W | - | E | Number of vertical photo-sensitive cells combined together.  |
|  ChunkDecimationHorizontal | O | Integer | R/W | - | E | Horizontal sub-sampling of the image.  |
|  ChunkDecimationVertical | O | Integer | R/W | - | E | Vertical sub-sampling of the image.  |
|  ChunkReverseX | R | Boolean | R/W | - | E | Flip horizontal of the image sent by the device.  |
|  ChunkReverseY | R | Boolean | R/W | - | E | Flip vertically of the image sent by the device.  |
|  ChunkTimestamp | R | Integer | R | - | E | Returns the Timestamp of the image included in the payload at the time of the FrameStart internal event.  |
|  ChunkTimestampLatchValue | R | Integer | R | ns | E | Returns the last Timestamp latched with the TimestampLatch command.  |
|  ChunkLineStatusAll | R | Integer | R | - | E | Returns the status of all the I/O lines at the time of the FrameStart internal event.  |
|  ChunkCounterSelector | R | Enumeration | R/W | - | E | Selects which counter to retrieve data from.  |
|  ChunkCounterValue[ChunkCounterSelector] | R | Integer | R | - | E | Returns the value of the selected Chunk counter at the time of the FrameStart event.  |
|  ChunkTimerSelector | R | Enumeration | R/W | - | E | Selects which Timer to retrieve data from.  |
|  ChunkTimerValue[ChunkTimerSelector] | R | Float | R | us | E | Returns the value of the selected Timer at the time of the FrameStart internal event.  |
|  ChunkScanLineSelector | O | Integer | R/W | - | E | Index for vector representation of one chunk value per line in an image.  |