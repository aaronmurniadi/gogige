|  ![img-103.jpeg](img-103.jpeg)CAN |   | ![img-104.jpeg](img-104.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 6.3.2.6 TLOpen

GC ERROR TLOpen

( TL_HANDLE * phSystem )

Opens the System module and puts the instance in the phSystem handle. This allocates all system wide resources. Call the GCInitLib function before this function. A System module can only be opened once.

##### Parameters

[out] phSystem

System module handle of the newly opened system. It is recommended to initialize *phSystem to GENTL_INVALID_HANDLE before calling TLOpen to indicate an invalid handle.

##### Returns

GC_ERR_SUCCESS

Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED

No preceding call to GCInitLib

GC_ERR_RESOURCE_IN_USE

The TL module has already been instantiated through a previous call to TLOpen.

GC_ERR_INVALID_PARAMETER

Parameter phSystem is an invalid pointer (NULL or ~0x0).

GC_ERR_ACCESS_DENIED

The access to the requested System module is denied. This may be because it is already opened by another process but it might have other reasons as well.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.2.7 TLOpenInterface

GC ERROR TLOpenInterface

( TL_HANDLE hSystem,
const char * sIfaceID,
IF_HANDLE * phIface )

Opens the given sIfaceID on the given hSystem.

Any subsequent call to TLOpenInterface with an sIfaceID which has already been opened will return the error GC_ERR_RESOURCE_IN_USE.

The interface ID need not match the one returned from TLGetInterfaceID. As long as the GenTL Producer knows how to interpret that ID it will return a valid handle. For example, if in a specific implementation the interface has a user-defined name, this function will return a valid handle as long as the provided name refers to an internally known interface.