|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

## 4 Configuration and Signaling

Every module from the System to the Data Stream supports a GenTL Port for the configuration of the module's internal settings and the Signaling to the calling GenTL Consumer. For the Buffer module the GenTL Port is optional.

For a detailed description of the C function interface and data types see chapter 6 Software Interface (page 59ff). Before a module can be configured or an event can be registered the module to be accessed must be instantiated. This is done through module enumeration as described in chapter 3 Module Enumeration and Instantiation (page 20ff).

### 4.1 Configuration

To configure a module and access transport layer technology specific settings a GenTL Port with a GenApi compliant XML description is used. The module specific functions' concern is the enumeration, instantiation, configuration and basic information retrieval. Configuration is done through a virtual register map and a GenApi XML description for that register map.

For a GenApi reference implementation's IPort interface the TLI publishes Port functions. A GenApi IPort expects a Read and a Write function which reads or writes a block of data from the associated device. Regarding the GenTL Producer's feature access each module acts as a device for the GenApi implementation by implementing a virtual register map. When certain registers are written or read, implementation dependent operations are performed in the specified module. Thus the abstraction made for camera configuration is transferred also to the GenTL Producer.

The memory layout of that virtual register map is not specified and thus it is up to the GenTL Producer's implementation. A certain set of mandatory features must be implemented which are described in chapter 7, Standard Features Naming Convention for GenTL (GenTL SFNC) (page 180ff).

The Port functions of the C interface include a GCReadPort function and a GCWritePort function which can be used to implement an IPort object for the GenApi implementation. These functions resemble the IPort Read and Write functions in their behavior.

Register access through the Port functions is always byte aligned. In case the underlying technology does not allow byte aligned access the GenTL Producer must simulate that by reading more bytes than requested and returning only the requested bytes and by doing a read/modify/write access to the ports register map.

#### 4.1.1 Modules

Every GenTL module except the Buffer module must support the Port functions of the TLI. The Buffer module can support these functions. To access the registers of a module the GCReadPort and GCWritePort functions are called on the module's handle, for example on the TL_HANDLE for the System module. A GenApi XML description file and the GenApi Module of GenICam is used to access the virtual register map in the module using GenApi features.