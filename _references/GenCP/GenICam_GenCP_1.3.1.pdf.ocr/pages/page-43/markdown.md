|  GEN<i>CAM |   | ![img-46.jpeg](img-46.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

|  2 | 4 | reservedReserved, set to 0  |
| --- | --- | --- |
|  2 | 6 | length 1 written (Len1)Number of bytes successfully written to the remote device's register map. For WRITEMEM_STACKED_ACK it is mandatory to report the length written (different than with the WRITEMEM_ACK).  |
|   |  | ..  |
|  2 | (n-1)*4 | reservedReserved, set to 0  |
|  2 | 2+(n-1)*4 | length n-1 written (Lenn-1)Number of bytes successfully written to the remote device's register map. For WRITEMEM_STACKED_ACK it is mandatory to report the length written (different than the WRITEMEM_ACK).  |
|  Postfix  |   |   |

Table 16 – WriteMemStacked Ack SCD-Fields

The writes are executed sequentially. In case of an error during a write command, subsequent writes are not executed and the WRITEMEM_STACKED_ACK returns the status. The length x written fields within the WRITEMEM_STACKED_ACK reflect the successful written bytes.