|  GEN<I>CAM |   | ![img-96.jpeg](img-96.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

expected to have 8 bits. On platforms/compilers where this is not the case a byte like primitive data type must be used.

String encoding is by default ASCII (characters with numerical values up to and including 127) unless stated different through the TL_INFO_CHAR_ENCODING command. A string as an input value without a size parameter must be 0-terminated. Strings with a size parameter must include the terminating 0.

##### 6.2.1.5 Primitive Data Types

The size_t type indicates that a buffer size is represented. This is a platform dependent unsigned integer (e.g., 32 bit on 32 bit platforms).

The ptrdiff_t is a signed type which indicates that its value relates to an arithmetic operation with a memory pointer, usually a buffer. Its size is platform dependent (e.g., 32 bit on 32 bit platforms and 64Bit on 64Bit platforms).

The other functions use a notation defining its base type and size. uint8_t stands for an unsigned integer with the size of 8 bits. int32_t is a signed integer with 32 bits size.

### 6.3 Function Declarations

#### 6.3.1 Library Functions

##### 6.3.1.1 GCCloseLib

GC_ERROR GCCloseLib (void)

This function must be called after no function of the GenTL library is needed anymore to clean up the resources from the GCInitLib function call. Each call to GCCloseLib has to be accompanied by a preceding call to GCInitLib.

GCGetLastError must not be called after the call of this function!

#### Returns

GC_ERR_SUCCESS

Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED

No preceding call to GCInitLib or library has already been closed through a call to GCCLoseLib.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.