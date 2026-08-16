|  ![img-93.jpeg](img-93.jpeg) CAM |   | ![img-94.jpeg](img-94.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Value | Description  |
| --- | --- | --- |
|   |  | Producer internal object is out of bounds.  |
|  GC_ERR_PARSING_CHUNK_DATA | -1018 | An error occurred parsing a buffer containing chunk data.  |
|  GC_ERR_INVALID_VALUE | -1019 | A register write function was trying to write an invalid value.  |
|  GC_ERR_RESOURCE_EXHAUSTED | -1020 | A requested resource is exhausted. This is a rather general error which might for example refer to a limited number of available handles being available.  |
|  GC_ERR_OUT_OF_MEMORY | -1021 | The system and/or other hardware in the system (frame grabber) ran out of memory.  |
|  GC_ERR_BUSY | -1022 | The required operation cannot be executed because the responsible module/entity is busy executing actions that cannot be performed concurrently with the requested operation.  |
|  GC_ERR_AMBIGUOUS | -1023 | The required operation cannot be executed unambiguously in given context.  |
|  GC_ERR_CUSTOM_ID | -10000 | Any error smaller or equal than -10000 is implementation specific. If a GenTL Consumer receives such an error number it should react as if it would be a generic runtime error.  |

To get a detailed descriptive text about the error reason call the GCGetLastError function.

Some error codes might be returned by any function and are therefore not explicitly listed in the function's error code table. These error codes are:

- GC_ERR_ERROR
- GC_ERR_IO
- GC_ERR_RESOURCE_EXHAUSTED
- GC_ERR_OUT_OF_MEMORY

#### 6.1.6 Software Interface Versions

The software interface evolves over the individual versions of the GenTL specification. In particular, between two versions of the interface, new functions (and corresponding data structures) and enumerations might be introduced. In rare cases, existing functions or commands might be conversely deprecated. Interface versions are indicated by a major