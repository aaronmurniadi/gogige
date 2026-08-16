|  GEN<i>CAM |   | ![img-30.jpeg](img-30.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

#### 4.3.2. Acknowledge Packet Layout

|  Width (Bytes) | Offset (Bytes) | Description  |
| --- | --- | --- |
|  Prefix  |   |   |
|  2 | 0 | status codeStatus code, indicating the result of the operation.See chapter 4.3.2.1 for a list of codes.  |
|  2 | 2 | command_idCommand id as specified in the command_id chapter 4.3.3  |
|  2 | 4 | lengthLength of the Specific Command Data depending on the command in bytes.  |
|  2 | 6 | request_idSequential number used to identify a single acknowledge. This id is provided by the command sender and incremented every time a new command is issued.  |
|  SCD  |   |   |
|  Postfix  |   |   |

Table 5 – Acknowledge layout