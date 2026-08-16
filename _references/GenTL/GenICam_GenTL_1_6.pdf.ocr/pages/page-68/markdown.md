|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 6.3.1.4 GCInitLib

GC ERROR GCInitLib (void)

This function must be called prior to any other function call to allow global initialization of the GenTL Producer driver. This function is necessary since automated initialization functionality like within DllMain on MS Windows platforms is very limited. Multiple calls to GCInitLib without accompanied calls to GCCloseLib will return the error GC_ERR_RESOURCE_IN_USE.

##### Returns

GC_ERR_SUCCESS Operation was successful; no error occurred.

GC_ERR_RESOURCE_IN_USE GCInitLib already called without accompanied call to GCCloseLib.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

#### 6.3.2 System Functions

##### 6.3.2.1 TLClose

GC ERROR TLClose ( TL_HANDLE hSystem )

Closes the System module associated with the given hSystem handle. This closes the whole GenTL Producer driver and frees all resources. Call the GCCloseLib function afterwards if the library is not needed anymore.

##### Parameters

[in] hSystem System module handle to close.

##### Returns

GC_ERR_SUCCESS Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE The handle hSystem is invalid (NULL) or does not reference an open System module retrieved through a call to TLOpen.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.