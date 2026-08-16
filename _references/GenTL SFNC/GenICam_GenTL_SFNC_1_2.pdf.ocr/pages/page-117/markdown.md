|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Access | Read  |
| --- | --- |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Flag to indicate if the buffer is in the input pool or output buffer queue.

Corresponds to the BUFFER_INFO_IS_QUEUED command of DSGetBufferInfo function.

### 3.5.2.6 BufferIsAcquiring

|  Name | BufferIsAcquiring  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Flag to indicate that the buffer is currently being filled with data.

Corresponds to the BUFFER_INFO_IS_ACQUIRING command of DSGetBufferInfo function.

### 3.5.2.7 BufferIsIncomplete

|  Name | BufferIsIncomplete  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |