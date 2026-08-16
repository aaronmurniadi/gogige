|  ![img-177.jpeg](img-177.jpeg) CAM |   | ![img-178.jpeg](img-178.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Value | Description  |
| --- | --- | --- |
|   |  | and anything for true.  |
|  INFO_DATATYPE_SIZET | 12 | Platform dependent unsigned integer (32 bit on 32 bit platforms)  |
|  INFO_DATATYPE_BUFFER | 13 | Like a INFO_DATATYPE_STRING but with arbitrary data and no 0 termination.  |
|  INFO_DATATYPE_PTRDIFF | 14 | The type ptrdiff_t is a type that can hold the result of subtracting two pointers.  |
|  INFO_DATATYPE_CUSTOM_ID | 1000 | Starting value for Custom IDs which are implementation specific.If a generic GenTL Consumer is using custom data types provided through a specific GenTL Producer implementation it must differentiate the handling of GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |

##### 6.4.1.2 TL_CHAR_ENCODING_LIST

enum TL_CHAR_ENCODING_LIST

Char encoding schemata.

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  TL_CHAR_ENCODING_ASCII | 0 | Char encoding of the GenTL Producer is ASCII.  |
|  TL_CHAR_ENCODING_UTF8 | 1 | Char encoding of the GenTL Producer is UTF8.  |

##### 6.4.1.3 TL_INFO_CMD

enum TL_INFO_CMD

System module information commands for the TLGetInfo and GCGetInfo functions. The reported information through these two functions must be in sync.

The column labeled “Impl” in the following table lists if the implementation of a given command is mandatory (M) or optional (O).

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|  TL_INFO_ID | M | 0 | Unique ID identifying a GenTL Producer. For example the filename of the GenTL Producer implementation including its path.  |