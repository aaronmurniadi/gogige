|  ![img-175.jpeg](img-175.jpeg) CAM |   | ![img-176.jpeg](img-176.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_NOT_INITIALIZED

The event has not previously been registered through GCRegisterEvent or no preceding call to GCInitLib has been made.

GC_ERR_INVALID_HANDLE

The handle hModule is invalid (NULL) or does not reference a previously instantiated module.

GC_ERR_NOT_IMPLEMENTED

The specified event type is not implemented in the provided module of the GenTL Producer.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

### 6.4 Enumerations

Enumeration values are signed 32 bit integers.

#### 6.4.1 Library and System Enumerations

##### 6.4.1.1 INFO_DATATYPE

enum INFO_DATATYPE

Defines the data type possible for the various Info functions. The data type itself may define its size. For buffer or string types the piSize parameter must be used to query the actual amount of data being written.

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  INFO_DATATYPE_UNKNOWN | 0 | Unknown data type. This value is never returned from a function but can be used to initialize the variable to inquire the type.  |
|  INFO_DATATYPE_STRING | 1 | 0-terminated C string (encoding according to the TL_INFO_CHAR_ENCODING info command).  |
|  INFO_DATATYPE_STRINGLIST | 2 | Concatenated INFO_DATATYPE_STRING list. End of list is signaled with an additional 0.  |
|  INFO_DATATYPE_INT16 | 3 | Signed 16 bit integer.  |
|  INFO_DATATYPE_UINT16 | 4 | Unsigned 16 bit integer.  |
|  INFO_DATATYPE_INT32 | 5 | Signed 32 bit integer.  |
|  INFO_DATATYPE_UINT32 | 6 | Unsigned 32 bit integer.  |
|  INFO_DATATYPE_INT64 | 7 | Signed 64 bit integer.  |
|  INFO_DATATYPE_UINT64 | 8 | Unsigned 64 bit integer.  |
|  INFO_DATATYPE_FLOAT64 | 9 | Signed 64 bit floating point number.  |
|  INFO_DATATYPE_PTR | 10 | Pointer type (void*). Size is platform dependent (32 bit on 32 bit platforms)  |
|  INFO_DATATYPE_BOOL8 | 11 | Boolean value occupying 8 bit. 0 for false  |