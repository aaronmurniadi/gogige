|  ![img-25.jpeg](img-25.jpeg) CAM |   | ![img-26.jpeg](img-26.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

features is specific to each respective technology and in many cases no specific configuration is required at all.

When an interface is not needed anymore it must be closed with the IFClose function. This frees the resources of this Interface and all child Device modules still open.

After an Interface module has been closed it may be opened again and the handle to the module may be different from the first instantiation.

### 3.4 Device

A Device module represents the GenTL Producer driver's view on a remote device. If the Device is able to output streaming data, this module is used to enumerate the available data streams. The number of available data streams is limited first by the remote device and second by the GenTL Producer implementation. Dependent on the implementation it might be possible that only one of multiple stream channels can be acquired.

If a GenTL Consumer requests a Device that has been instantiated from within the same process beforehand and has not been closed, the Interface returns an error. If the instance was created in another process space and the GenTL Producer explicitly wants to grant access to the Device this access should be restricted to read only. The module does no reference counting within one process space. If a Device module handle is requested a second time from within one process space, the second call will return the error GC ERR RESOURCE IN USE. The first call from within that process to the DevClose function will free all resources and shut down the module including all child modules in that process.

Every device is identified not by an index but by an Interface module wide unique ID. It is recommended to have a general unique identifier for a specific device. The ID of the GenTL Device module should be different to the remote device ID. The content of this ID is up to the GenTL Producer and is only interpreted by it and not by any GenTL Consumer.

For convenience a GenTL Producer may allow opening a device not only by its unique ID. Other representations may be a user defined name or a transport layer technology dependent ID like for example an IP address for IP-based devices.

To get the number of available data streams the DevGetNumDataStreams function is called using the DEV_HANDLE returned from the Interface module. As with the Interface and the Device lists, this list holds the unique IDs of the available streams. The number of data streams or the data stream IDs may not change during a single session. The IDs of the data streams must be fix between all sessions.

To get access to the Port object associated with a Device the function DevGetPort must be called.

A Data Stream module can be instantiated by using the DevOpenDataStream function. As with the IDs of the modules discussed before a known ID can be used to open a Data Stream directly. The ID must not change between different sessions. To obtain a unique ID for a Data Stream call the DevGetDataStreamID function.