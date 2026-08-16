|  GEN<i>CAM |   | ![img-52.jpeg](img-52.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

#### 5.4.1. GenCP Version

Version of the GenCP specification this Bootstrap Register Map complies with.

|  Offset | Hex 0  |
| --- | --- |
|  Length | 4  |
|  Access Type | R  |
|  Support | M  |
|  Data Type | 2 x 16bit fields  |
|  Factory Default | Implementation specific  |

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 16 | Minor VersionMinor Version of the Standard this BRM and the protocol the device's implementation complies to.  |
|  16 | 16 | Major VersionMajor Version of the Standard this BRM and the protocol the device's implementation complies to.  |

Table 20 – Register GenCP Version

#### 5.4.2. Manufacturer Name

Manufacturer Name is a string containing a human readable manufacturer name.

|  Offset | Hex 4  |
| --- | --- |
|  Length | 64  |
|  Access Type | R  |
|  Support | M  |
|  Data Type | String  |
|  Factory Default | Device specific  |