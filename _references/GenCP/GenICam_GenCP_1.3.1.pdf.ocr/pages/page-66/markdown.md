|  GEN<i>CAM |   | ![img-70.jpeg](img-70.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

5.5.1.1. Manifest Table

|  Width (Bytes) | Offset (Bytes) | Support | Access | Description  |
| --- | --- | --- | --- | --- |
|  8 | 0 | M | R | MT Entry CountNumber of entries in the Manifest Table  |
|  64 | 8 | M | R | Manifest Entry 0First entry in the Manifest Table  |
|  64 | 8 + 64 | O | R | Manifest Entry 1Second entry in the Manifest Table  |
|  ... | ... |  |  | ...  |
|  64 | 8 + n*64 | O | R | Manifest Entry n(N+1)th entry in the Manifest Table  |

Table 33 – Manifest Table Layout