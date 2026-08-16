|  ![img-89.jpeg](img-89.jpeg) CAM |   | ![img-90.jpeg](img-90.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

If a platform supports independent processes the GenTL Producer implementation may establish interprocess communication. Minimal requirement is that other processes are not allowed to use an opened Device module. It is recommended though that a GenTL Producer implementation is multiprocess capable to the point where:

- Access rights for the Modules are checked
An open Device module should be locked against multiple process write access. In that case an error should be returned. Read access may be granted though.
• Data/state safety is ensured
Reference counting must be done so that if, e.g., the System module of one process is closed the resources of another process are not automatically freed.
- Different processes can communicate with different devices
Each process should be able to communicate with one or multiple devices. Furthermore different processes should be able to communicate with different devices.

#### 6.1.5 Error Handling

Every function has as its return value a GC_ERROR. This value indicates the status of the operation. Functions must give strong exception safety. With an exception not a language dependent exception object is meant, but an execution error in the called function with a return code other than GC_ERR_SUCCESS. No exception objects may be thrown of any exported function. Strong exception safety means:

• Data validity is preserved
No data becomes corrupted or leaked.

• State is unchanged

First the internal state must stay consistent and it must be as if the function encountering the error was never called. Therefore the output values of a function are to be handled as if being invalid if the function returns an error code.

This ensures that calling GenTL Consumers always can expect a known state in the GenTL Producer implementation: either it is the desired state when a function call was successful or it is the state the GenTL Producer implementation had before the call.

The following values are defined:

Table 6-4: C interface error codes

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  GC_ERR_SUCCESS | 0 | Operation was successful; no error occurred.  |
|  GC_ERR_ERROR | -1001 | Unspecified runtime error.  |
|  GC_ERR_NOT_INITIALIZED | -1002 | Module or resource not initialized; e.g., GCInitLib was not called .  |
|  GC_ERR_NOT_IMPLEMENTED | -1003 | Requested operation not implemented; e.g., no Port functions on a Buffer module.  |
|  GC_ERR_RESOURCE_IN_USE | -1004 | Requested resource is already in use.  |