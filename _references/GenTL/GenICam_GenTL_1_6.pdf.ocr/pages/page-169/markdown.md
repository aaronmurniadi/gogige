|  ![img-248.jpeg](img-248.jpeg) CAM |   | ![img-249.jpeg](img-249.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|  URL_INFO_FILE_SIZE | CM | 8 | File size of the XML-File in bytes in the register map. For URLs starting with ‘file:’ or ‘http:’ the according info function should return a GC_ERR_NOT_AVAILABLE. Data type: UINT64  |
|  URL_INFO_SCHEME | CM | 9 | Scheme of the URL. Possible values are defined in URL_SCHEME_IDS. Data type: INT32  |
|  URL_INFO_FILENAME | CM | 10 | Filename in case the scheme of the URL is URL_SCHEME_FILE or as a hint if the scheme is URL_SCHEME_LOCAL. Data type: STRING  |
|  URL_INFO_CUSTOM_ID | O | 1000 | Starting value for GenTL Producer custom IDs which are implementation specific. If a generic GenTL Consumer is using custom URL_INFO_CMDs provided through a specific GenTL Producer implementation it must differentiate the handling of different GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |

#### 6.4.5.3 URL_SCHEME_IDS

enum URL_SCHEME_IDS

This enumeration defines the values to be retrieved through a call to GCGetPortURLInfo with the command URL_INFO_SCHEME.

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  URL_SCHEME_LOCAL | 0 | The XML-File is to be retrieved from the local register map. The address and size where it can be read can be queried using theGCGetPortURLInfo function with theURL_INFO_FILE_REGISTER_ADDRESSandURL_INFO_FILE_SIZE command.  |