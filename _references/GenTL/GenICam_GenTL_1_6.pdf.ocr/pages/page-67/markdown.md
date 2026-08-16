|  GEN<I>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 6.3.1.3 GCGetLastError

|  GC_ERROR | GCGetLastError | ( GC_ERROR * char * size_t * | piErrorCode, sErrorText, piSize )  |
| --- | --- | --- | --- |

Returns a readable text description of the last error occurred in the local thread context.

If multiple threads are supported on a platform this function must store this information thread local. In case an error occurs and after that several other function calls return without error the last error value and description is returned and the successful calls are ignored. If there has not been any error in the given thread context since startup the function will return GC_ERR_SUCCESS with *piErrorCode also set to GC_ERR_SUCCESS and sErrorText containing "No Error". In case GCGetLastError itself generates an error it will return the according error code but it will not store the error internally so that succeeding calls to GCGetLastError will still be able to report the stored error code.

##### Parameters

|  [out] piErrorCode | Error code of the last error.  |
| --- | --- |
|  [in,out] sErrorText | Pointer to a user allocated C string buffer to receive the last error text. If this parameter is NULL, piSize will contain the needed size of sErrorText in bytes. The size includes the terminating 0.  |
|  [in,out] piSize | sErrorText equal NULL:out: minimal size of sErrorText in bytes to hold all information.sErrorText unequal NULL:in: size of the provided sErrorText in bytesout: number of bytes filled by the function.  |

##### Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_PARAMETER | Parameters piSize and/or piErrorCode are invalid pointers (NULL or ~0x0).  |
|  GC_ERR_BUFFER_TOO_SMALL | sErrorText is not NULL and the value of *piSize is too small to receive the expected amount of data.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.