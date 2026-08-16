|  ![img-105.jpeg](img-105.jpeg) CAM |   | ![img-106.jpeg](img-106.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

# Parameters

|  [in] | hSystem  |
| --- | --- |
|  [in] | sIfaceID  |
|  [out] | phIface  |

System module to work on.

Unique interface ID to open as a 0-terminated C string.

Interface handle of the newly created interface. It is recommended to initialize *phIface to

GENTL_INVALID_HANDLE before calling

TLOpenInterface to indicate an invalid handle.

# Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_RESOURCE_IN_USE | The Interface module has already been instantiated through a previous call to TLOpenInterface.  |
|  GC_ERR_INVALID_HANDLE | The handle hSystem is invalid (NULL) or does not reference an open System module retrieved through a call to TLOpen.  |
|  GC_ERR_INVALID_ID | The GenTL Producer is unable to interpret the provided ID string sIfaceID or is not able to match it to an existing Interface.  |
|  GC_ERR_INVALID_PARAMETER | Parameters phIface and/or sIfaceID are invalid pointers (NULL or ~0x0).  |
|  GC_ERR_ACCESS_DENIED | The access to the requested Interface is denied. This may be because it is already opened by another Process but it might have other reasons as well.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.2.8 TLUpdateInterfaceList

|  GC_ERROR | TLUpdateInterfaceList ( TL_HANDLE bool8_t * uint64_t | hSystem, pbChanged, iTimeout )  |
| --- | --- | --- |

Updates the internal list of available interfaces. This may change the connection between a list index and an interface ID. It is recommended to call TLUpdateInterfaceList after reconfiguration of the System module to reflect possible changes.

A call to this function has implications on the thread safety of

- TLGetNumInterfaces
- TLGetInterfaceID