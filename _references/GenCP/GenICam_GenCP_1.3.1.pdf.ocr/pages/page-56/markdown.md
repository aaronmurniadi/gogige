|  GEN<i>CAM |   | ![img-60.jpeg](img-60.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

#### 5.4.11. Manifest Table Address

Pointer to the Manifest table, containing the URLs for the GenICam files for this device. (See chapter 5.5.1)

|  Offset | Hex 1D0  |
| --- | --- |
|  Length | 8  |
|  Access Type | R  |
|  Support | M  |
|  Data Type | UINT64  |
|  Factory Default | Implementation specific  |

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 64 | Manifest Table Address64-bit register address of the Manifest Table  |

Table 23 – Register Manifest Table Offset

#### 5.4.12. SBRM Address

The register contains a pointer to the Technology Specific Bootstrap Register Map. The SBRM Supported bit in the Device Capability register indicates if this register is present or not.

|  Offset | Hex 1D8  |
| --- | --- |
|  Length | 8  |
|  Access Type | R  |
|  Support | CM  |
|  Data Type | UINT64  |
|  Factory Default | Implementation Specific  |