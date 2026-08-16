|  ![img-278.jpeg](img-278.jpeg)CAN |   | ![img-279.jpeg](img-279.jpeg)  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Name | Interface | Access | Description  |
| --- | --- | --- | --- |
|  StreamType | IEnumeration | R | Identifies the transport layer technology of the stream. See chapter 6.6.1 for possible values.  |

#### 7.1.5 Buffer Module

All features that must be accessible on a buffer if a Port access is provided are listed here. Port functions use the BUFFER_HANDLE to access these features. The Port access for the BUFFER_HANDLE is not mandatory. Thus all features listed here need not be implemented. If a Port access is implemented on the handle though, all mandatory features must be present.

Table 7-12: Buffer information features

|  Name | Interface | Access | Description  |
| --- | --- | --- | --- |
|  BufferData | IRegister | R/(W) | Entire buffer data.  |
|  BufferUserData | IInteger | R | Pointer to user data (pPrivate) casted to an integer number referencing GenTL Consumer specific data. It is reflecting the pointer provided by the user data pointer (pPrivate) at buffer announcement.(see chapter 6.3.5 Data Stream Functions page 89ff). This allows the GenTL Consumer to attach information to a buffer.  |