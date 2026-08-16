|  ![img-27.jpeg](img-27.jpeg) CAM |   | ![img-28.jpeg](img-28.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

Stream module. A GenTL Producer must implement both methods even if one of them is of lesser performance. The simplest implementation for DSAllocAndAnnounceBuffer would be a malloc from the platform SDK.

If the same buffer is announced twice on a single stream via a call to DSAnnounceBuffer the error GC_ERR_RESOURCE_IN_USE is returned. A buffer may be announced to multiple streams. In this case individual handles for each stream will be returned. In general there is no synchronization or locking mechanism between two streams defined. A GenTL Producer may though provide special functionality to prevent data loss. In case a GenTL Producer is not able to handle buffers announced to multiple streams it may refuse the announcement and return GC_ERR_RESOURCE_IN_USE.

The required size of the buffer must be retrieved either from the Data Stream module the buffer will be announced to or from the associated remote device (see chapter 5.2.1 for further details).

To allow the acquisition engine to stream data into a buffer it has to be placed into the Input Buffer Pool by calling the DSQueueBuffer function with the BUFFER_HANDLE retrieved through buffer announcement functions.

A BUFFER_HANDLE retrieved either by DSAnnounceBuffer, DSAllocAndAnnounceBuffer or DSAnnounceCompositeBuffer can be released through a call to DSRevokeBuffer. A buffer which is still in the Input Buffer Pool or the Output Buffer Queue of the acquisition engine cannot be revoked and an error is returned when tried. A memory buffer must only be announced once to a single stream.

The more advanced method of announcing “composite” buffers using DSAnnounceCompositeBuffer is described in 5.7.1.

### 3.7 Enumerated modules list overview

The purpose of this chapter is to highlight possible issues related to the maintenance of the list of GenTL modules (interfaces, devices) available in a system. It provides a summary of principles listed in other chapters of the specification.

While the set of Data Stream modules implemented by a device is static and stays fixed throughout the lifetime of the local GenTL Device module, the lists of interfaces within a system and devices discovered on an interface are dynamic and might be updated on request by the GenTL Consumer.

The explicit request to update the list might be issued through the C interface (TLUpdateInterfaceList and IFUpdateDeviceList functions) or through corresponding commands (InterfaceUpdateList, DeviceUpdateList) of the parent module.

It's important to remark that there might be multiple different views of the list of “currently available” modules, which we’ll demonstrate on an example of devices discovered on an interface:

- Real devices that are physically connected to the interface. If a new device is connected at runtime (or powered up), the GenTL Producer might or might not be aware of it. This depends on whether it actively monitors the interface. But it will not be published to the