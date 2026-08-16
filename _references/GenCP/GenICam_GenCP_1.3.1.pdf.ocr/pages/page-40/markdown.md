|  GEN<i>CAM |   | ![img-38.jpeg](img-38.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

|  Prefix  |   |   |
| --- | --- | --- |
|  CCD (command_id = READMEM_STACKED_CMD)  |   |   |
|  8 | 0 | register address 064 bit register address of the first data block to read.  |
|  2 | 8 | reservedReserved, set to 0  |
|  2 | 10 | read length 0 (Len0)Number of bytes to read from address 0.  |
|  8 | (1*12) | register address 164 bit register address of the second data block.  |
|  2 | 8+(1*12) | reservedReserved, set to 0  |
|  2 | 10+(1*12) | read length 1 (Len1)Number of bytes to read from address 1.  |
|  ...  |   |   |
|  8 | ((n-1)*12) | register address n-164 bit register address of the last data block to read.  |
|  2 | 8+((n-1)*12) | reservedReserved, set to 0  |
|  2 | 10+((n-1)*12) | read length n-1 (Lenn-1)Number of bytes to read from address n-1.  |
|  Postfix  |   |   |

Table 13 – ReadMemStacked SCD-Fields

#### 4.4.7. ReadMemStacked Acknowledge

The ReadMemStacked acknowledge states the result of a ReadMemStacked command.

|  Width (Bytes) | Offset (Bytes) | Description  |
| --- | --- | --- |
|  Prefix  |   |   |
|  CCD-ACK (command_id = READMEM_STACKED_ACK)  |   |   |
|  Len0 | 0 | dataData read from the remote device's register map.  |