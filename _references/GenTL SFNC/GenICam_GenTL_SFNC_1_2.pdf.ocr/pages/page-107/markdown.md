|  GEN<I>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Visibility | Expert  |
| --- | --- |
|  Values | ≥0  |

Number of frames started in the acquisition engine. This number is incremented every time in case of a new buffer is started and then to be filled (data written to) regardless even if the buffer is later delivered to the user or discarded for any reason. This number is initialized with 0 at the time the stream is opened. It is not reset until the stream is closed.

Corresponds to the STREAM_INFO_NUM_STARTED command of DSGetInfo function.

### 3.4.3.10 PayloadSize

|  Name | PayloadSize  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Byte  |
|  Visibility | Expert  |
|  Values | >0  |

Size of the expected data in bytes. Note that this feature "overwrites" the PayloadSize of the remote device, see also sections "Data Payload Delivery" and "Allocate Memory" of the GenICam GenTL standard.

Corresponds to the STREAM_INFO_PAYLOAD_SIZE command of DSGetInfo function.

### 3.4.3.11 StreamIsGrabbing

|  Name | StreamIsGrabbing  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Recommended  |
|  Interface | IBoolean  |
|  Access | Read  |
|  Unit |   |
|  Visibility | Expert  |