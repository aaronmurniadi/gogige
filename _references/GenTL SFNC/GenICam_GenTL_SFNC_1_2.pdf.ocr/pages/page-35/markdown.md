|  StreamStartedFrameCount | R | All | Integer | R | - | E | Number of frames started in the acquisition engine.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  PayloadSize | R | All | Integer | R | Byte | E | Size of the expected data in bytes.  |
|  StreamIsGrabbing | R | All | Boolean | R |  | E | Flag indicating whether the acquisition engine is started or not.  |
|  StreamChunkCountMaximum | R | All | Integer | R |  | E | Maximum number of chunks to be expected in a buffer (can be used to allocate the array for the DSGetBufferChunkData function).  |
|  StreamBufferAlignment | R | All | Integer | R | Byte | E | Alignment size in bytes of the buffers passed to DSAnnounceBuffer.  |

#### 2.4.4 GenICam Control

Category that contains Event Control features.

Table 2-18: Event Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  Root | M | All | ICategory | R | - | B | Provides the Root of the GenICam features tree.  |
|  StreamPort | M | All | IPort | R/W | - | I | The GenICam port through which the Data Stream module is accessed.  |

#### 2.4.5 Event Control

Contains the features related to the Event Buffers Discarded.

Table 2-19: Buffer Discarded Event Summary