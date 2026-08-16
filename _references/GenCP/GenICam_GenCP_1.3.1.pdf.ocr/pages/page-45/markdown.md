|  GEN<i>CAM |   | ![img-48.jpeg](img-48.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

#### 4.4.11. Event Acknowledge

|  Width (Bytes) | Offset (Bytes) | Description  |
| --- | --- | --- |
|  Prefix  |   |   |
|  CCD-ACK (command_id = EVENT_ACK)  |   |   |
|  Postfix  |   |   |

Table 18 – Event Acknowledge SCD-Fields

### 4.5. Postfix

The Postfix carries data like a CRC in case the underlying protocol layers do not provide such services. The Postfix is conditional mandatory depending on the technology.