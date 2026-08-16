|  GEN<i>CAM |   | ![img-31.jpeg](img-31.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

##### 4.3.2.1. Status Codes

This section lists status codes that can be returned through an acknowledge packet. Each status code has 16 bits. The bits within the Status Code have the following meanings:

|  Bit offset (lsb << x) | Width (bits) | Description  |
| --- | --- | --- |
|  0 | 12 | Status Code  |
|  12 | 1 | ReservedSet to 0  |
|  13 | 2 | Namespace0 = GenCP Status Code1 = Technology specific Code2 = Device specific Code  |
|  15 | 1 | Severity0 = Warning/Info1 = Error  |

Warning and Info Status Codes indicate that the command was correctly executed and that the device resumes operation. For example, if a float value needed to be rounded it would be a warning but the rounded value has been set.