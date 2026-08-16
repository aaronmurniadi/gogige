|  ![img-258.jpeg](img-258.jpeg) CAM |   | ![img-259.jpeg](img-259.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Value | Description  |
| --- | --- | --- |
|   |  | This event type used to be called EVENT_FEATURE_DEVEVENT but has been renamed for a more intuitive understanding.  |
|  EVENT_MODULE | 5 | Notification that one GenTL Producer module wants to inform the GenICam GenApi instance of this module that a GenApi compatible event was fired. This Event is to be registered on any module handle except on the Remote Device. Values that can be retrieved are:EVENT_DATA_IDString representation of the EventID number in hexadecimal numbers with even number of digits and without the leading ‘0x’.Data type: STRING (Event ID)EVENT_DATA_VALUECorresponds to the data addressable through the module’s nodemap event port, beginning of the buffer corresponding to address 0.Data type: BUFFER (optional data)  |
|  EVENT_CUSTOM_ID | 1000 | Starting value for GenTL Producer custom events which are implementation specific. If a generic GenTL Consumer is using custom EVENT_TYPEs provided through a specific GenTL Producer implementation it must differentiate the handling of different GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |

### 6.5 Structures

Structures are byte aligned. The size of pointers as members is platform dependent.

#### 6.5.1 Data Stream Structures

##### 6.5.1.1 SINGLE_CHUNK_DATA

struct SINGLE_CHUNK_DATA

Layout of the array elements being used in the function DSGetBufferChunkData to carry information about individual chunks present in the buffer.