|  ![img-87.jpeg](img-87.jpeg) CAM |   | ![img-88.jpeg](img-88.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Entry | Description  |
| --- | --- |
|   | Close to close a moduleGet to query information about a module or object  |
|  Specifier | The specifier is optional. If an operation needs additional information it is provided by the Specifier.Values (choice):xxxInfo to retrieve xxx-object specific informationNumxxx to retrieve the number of xxx-objects  |

For example the function TLGetNumInterfaces works on the System module's TL_HANDLE and queries the number of available interfaces. TLClose for instance closes the System module.

#### 6.1.3 Memory and Object Management

The interface is designed in a way that objects and data allocated in the GenTL Producer implementation are only freed or reallocated inside the library. Vice versa all objects and data allocated by the calling GenTL Consumer must only be reallocated and freed by the calling GenTL Consumer. No language specific features except the ones allowed by ANSI C and published in the interface are allowed.

The function names of the exported functions must be undecorated. The function calling convention is stdcall for x86 platforms and architecture dependent for other platforms.

This ensures that the GenTL Producer implementation and the calling GenTL Consumer can use different heaps and different memory allocation strategies. Also language interchangeability is easier handled this way.

For functions filling a buffer (e.g., a C string) the function can be called with a NULL pointer for the char* parameter (buffer). The piSize parameter is then filled with the size of buffer needed to hold the information in bytes. For C strings that does incorporate the terminating 0 character. A function expecting a C string as its parameter not containing a size parameter for it expects a 0-terminated C string. Queries are not allowed for event data.

Objects that contain the state of one module's instance are referenced by handles (void*). If a module has been instantiated before and is opened a second time from within a single process the error GC ERR RESOURCE IN USE has to be returned. A close on the module will free the resource of the closed module and all underlying or depending child modules. This is true as long as these calls are in the same process space (see below). Thus if a Interface module is closed all attached Device, Data Stream and Buffer modules are also closed.

#### 6.1.4 Thread and Multiprocess Safety

If the platform supports threading, all functions must be thread safe to ensure data integrity when a function is called from different threads in one process. Certain restrictions apply for all list functions like TLUpdateInterfaceList and IFUpdateDeviceList since results are cached inside the module.