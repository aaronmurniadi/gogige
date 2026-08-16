|  GEN<i>CAM |   | ![img-54.jpeg](img-54.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

#### 5.4.5. Device Version (Manufacturer specific)

A string containing a Device Version.

An application must NOT make any assumptions based on the content of this string. Its content is purely manufacturer specific and may or may not change in case of e.g. a firmware update. See Device Software Interface Version for a defined way to deal with changes that affect the behavior of the device.

|  Offset | Hex C4  |
| --- | --- |
|  Length | 64  |
|  Access Type | R  |
|  Support | M  |
|  Data Type | String  |
|  Factory Default | Device specific  |

#### 5.4.6. Manufacturer Info

Manufacturer Info is a string containing manufacturer specific information. If there is none, this field should be all 0.

|  Offset | Hex 104  |
| --- | --- |
|  Length | 64  |
|  Access Type | R  |
|  Support | M  |
|  Data Type | String  |
|  Factory Default | Device specific  |