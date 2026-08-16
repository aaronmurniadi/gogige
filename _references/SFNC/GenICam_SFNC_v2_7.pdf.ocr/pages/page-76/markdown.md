|  ChunkEncoderSelector | R | IEnumeration | R/W | - | E | Selects which Encoder to retrieve data from.  |
| --- | --- | --- | --- | --- | --- | --- |
|  ChunkEncoderValue[ChunkEncoderSelector][ChunkScanLineSelector] | R | IInteger | R | - | E | Returns the counter's value of the selected Encoder at the time of the FrameStart in area scan mode or the counter's value at the time of the LineStart selected by ChunkScanLineSelector in Linescan mode.  |
|  ChunkEncoderStatus[ChunkEncoderSelector][ChunkScanLineSelector] | O | IEnumeration | R | - | E | Returns the motion status of the selected encoder.  |
|  ChunkExposureTimeSelector | O | IEnumeration | R/W | - | E | Selects which exposure time is read by the ChunkExposureTime feature.  |
|  ChunkExposureTime[ChunkExposureTimeSelector] | R | IFloat | R | us | E | Returns the exposure time used to capture the image.  |
|  ChunkGainSelector | R | IEnumeration | R/W | - | E | Selects which Gain to return.  |
|  ChunkGain[ChunkGainSelector] | R | IFloat | R | - | E | Returns the gain used to capture the image.  |
|  ChunkBlackLevelSelector | R | IEnumeration | R/W | - | E | Selects which Black Level to return.  |
|  ChunkBlackLevel[ChunkBlackLevelSelector] | R | IFloat | R | - | E | Returns the black level used to capture the image included in the payload.  |
|  ChunkLinePitch | R | IInteger | R | B | E | Returns the LinePitch of the image included in the payload.  |
|  ChunkFrameID | R | IInteger | R | - | E | Returns the unique Identifier of the frame (image or data container) included in the payload.  |
|  ChunkSourceSelector | O | IEnumeration | R/W | - | E | Selects which Source to retrieve data from.  |
|  ChunkSourceID | O | IEnumeration | R | - | E | Returns the Identifier of Source that the image comes from.  |
|  ChunkSourceIDValue[ChunkSourceSelector] | R | IInteger | R | - | E | Returns the unique integer Identifier value of the Source that the image comes from.  |
|  ChunkTransferBlockID | R | IInteger | R | - | E | Returns the unique identifier of the transfer block used to transport the payload.  |
|  ChunkTransferStreamID | R | IEnumeration | R | - | E | Returns identifier of the stream that generated this block.  |
|  ChunkTransferQueueCurrentBlockCount | O | IInteger | R | - | E | Returns the current number of blocks in the transfer queue.  |
|  ChunkStreamChannelID | R | IInteger | R | - | E | Returns identifier of the stream channel used to carry the block.  |