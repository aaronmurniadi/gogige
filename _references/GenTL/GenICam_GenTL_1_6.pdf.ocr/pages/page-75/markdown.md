|  GEN<ICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

# Parameters

[in] hSystem

[out] pbChanged

[in] iTimeout

System module to work on.

Contains true if the internal list was changed and false otherwise. If set to NULL nothing is written to this parameter.

Timeout in ms. If set to GENTL_INFINITE the timeout is infinite and the function will only return after the operation is completed. In any case the GenTL Producer must make sure that this operation is completed in a reasonable amount of time depending on the underlying technology. Please be aware that there is no defined way of terminating such an update operation. On the other hand it is the GenTL

Consumer's responsibility to call this function with a reasonable timeout.

# Returns

GC_ERR_SUCCESS

GC_ERR_NOT_INITIALIZED

GC_ERR_INVALID_HANDLE

GC_ERR_TIMEOUT

Operation was successful; no error occurred.

No preceding call to GCInitLib

The handle hSystem is invalid (NULL) or does not reference an open System module retrieved through a call to TLOpen.

The specified iTimeout expired before the Producer was able to completely update the list. In this case the “old” list stays valid.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

#### 6.3.3 Interface Functions

##### 6.3.3.1 IFClose

GC ERROR IFClose

( IF HANDLE

hIface )

Closes the Interface module associated with the given hIface handle. This closes all dependent Device modules and frees all interface related resources.

# Parameters

[in] hIface

Interface module handle to close.

# Returns

GC_ERR_SUCCESS

GC_ERR_NOT_INITIALIZED

Operation was successful; no error occurred.

No preceding call to GCInitLib.