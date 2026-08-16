|  GEN<i>CAM |   | ![img-35.jpeg](img-35.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

### 4.4. Command Specific Data

#### 4.4.1. ReadMem Command

Start address and length of any read access is byte aligned unless the underlying technology states different rules.

|  Width (Bytes) | Offset (Bytes) | Description  |
| --- | --- | --- |
|  Prefix  |   |   |
|  CCD (command_id = READMEM_CMD)  |   |   |
|  8 | 0 | register address64 bit register address.  |
|  2 | 8 | reservedReserved, set to 0  |
|  2 | 10 | read lengthNumber of bytes to read.  |
|  Postfix  |   |   |

Table 8 – ReadMem SCD-Fields

#### 4.4.2. ReadMem Acknowledge

|  Width (Bytes) | Offset (Bytes) | Description  |
| --- | --- | --- |
|  Prefix  |   |   |
|  CCD-ACK (command_id = READMEM_ACK)  |   |   |
|  x | 0 | DataData read from the remote device's register map. If the number of bytes read is different than specified in the relating READMEM_CMD the status of the READMEM_ACK must indicate the reason.  |
|  Postfix  |   |   |

Table 9 – ReadMem Ack SCD-Fields