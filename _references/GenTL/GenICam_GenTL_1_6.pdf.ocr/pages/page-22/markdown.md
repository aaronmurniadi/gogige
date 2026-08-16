|  ![img-21.jpeg](img-21.jpeg) CAM |   | ![img-22.jpeg](img-22.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

directly open an interface without inquiring the list of available interfaces via TLUpdateInterfaceList. That implies that the IDs must stay the same in-between two sessions. This is only guaranteed when the hardware does not change in any way. The TLUpdateInterfaceList function may be called nevertheless for the creation of the System's internal list of available interfaces. A GenTL Producer may call TLUpdateInterfaceList at module instantiation if needed. TLUpdateInterfaceList must be called by the GenTL Consumer before any call to TLGetNumInterfaces or TLGetInterfaceID. After successful module instantiation the TLUpdateInterfaceList function may be called by the GenTL Consumer so that it is aware of any change in this list. For convenience the GenTL Producer implementation may allow opening an Interface module not only using its unique ID but also with any other defined name. If the GenTL Consumer requests the ID of a module, the GenTL Producer must return its unique ID and not the convenience-name used to request the module's handle initially. This allows a GenTL Consumer, for example, to use the IP address of a network interface (in case of a GigE Vision GenTL Producer driver) to instantiate the module instead of using the unique ID.

When the GenTL Producer driver is not needed anymore the TLClose function must be called to close the System module and all other modules which are still open and relate to this System.

After a System module has been closed it may be opened again and the handle to the module may be different from the first instantiation.

### 3.3 Interface

An Interface module represents a specific hardware interface like a network interface card or a frame grabber. The exact definition of the meaning of an interface is left to the GenTL Producer implementation. After retrieving the IF_HANDLE from the System module all attached devices can be enumerated.

The size and order of the interface list provided by the System module can change during runtime only as a result of a call to the TLUpdateInterfaceList function. Interface modules may be closed in a random order which can differ from the order they have been instantiated in. The module does no reference counting. If an Interface module handle is requested a second time from within one process space the second call will return the error GC_ERR_RESOURCE_IN_USE. A single call from within that process to the IFClose function will free all resources and shut down the module in that process.

Every interface is identified by a System module wide unique ID and not by the index. The content of this ID is up to the GenTL Producer and is only interpreted by it and must not be interpreted by the GenTL Consumer.

In order to create or update the internal list of all available devices the IFUpdateDeviceList function may be called. The internal list of devices must not change its content unless this function is called again. It is recommended to call IFUpdateDeviceList regularly from time to time and after reconfiguration of the Interface module to reflect possible changes.