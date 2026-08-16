|  ![img-19.jpeg](img-19.jpeg) CAM |   | ![img-20.jpeg](img-20.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

### 3.2 System

The System module is always the entry point for the calling GenTL Consumer to the GenTL Producer. With the functions present here, all available hardware interfaces in the form of an Interface module can be enumerated.

By calling the TLOpen function the TL_HANDLE to work on the System module's functions can be retrieved. The TL_HANDLE obtained from a successful call to the TLOpen function will be needed for all successive calls to other functions belonging to the System module.

Before doing that, the GCGetInfo function might be called to retrieve the basic information about the GenTL Producer implementation without opening the system module.

Each GenTL Producer driver exposes only a single System instance in an operating system process space. If a GenTL Producer allows access from multiple processes it has to take care of the inter-process-communication and must handle the book-keeping of instantiated system modules. If it does not allow this kind of access it must return an appropriate error code whenever an attempt is made to create a second System module instance from another operating system process.

The System module does no reference counting within a single process. Thus even when a System module handle is requested twice from within a single process space, the second call will return the error GC ERR RESOURCE IN USE. The first call to the close function from within that process will free all resources and shut down the module.

Prior to the enumeration of the child interfaces the TLUpdateInterfaceList function must be called. The list of interfaces held by the System module must not change its content unless this function is called again. Any call to TLUpdateInterfaceList does not affect instantiated interface handles. It may only change the order of the internal list accessed via TLGetInterfaceID. The instantiation of a child interface with a known id is possible without a previous enumeration. It is recommended to call TLUpdateInterfaceList after reconfiguration of the System module to reflect possible changes.

The GenTL Consumer must make sure that calls to the TLUpdateInterfaceList function and the functions accessing the list are not made concurrent from multiple threads and that all threads are aware of the update operation, when performed. The GenTL Producer must make sure that any list access is done in a thread safe way.

After the list of available interfaces has been generated internally the TLGetNumInterfaces function retrieves the number of present interfaces known to this system. The list contains not the IF_HANDLES itself but their unique IDs of the individual interfaces. To retrieve such an ID the TLGetInterfaceID function must be called. This level of indirection allows the enumeration of several interfaces without the need to open them which can save resources and time.

If additional information is needed to be able to decide which interface is to be opened, the TLGetInterfaceInfo function can be called. This function enables the GenTL Consumer to query information on a single interface without opening it.

To open a specific interface the unique ID of that interface is passed to the TLOpenInterface function. If an ID is known prior to the call this ID can be used to