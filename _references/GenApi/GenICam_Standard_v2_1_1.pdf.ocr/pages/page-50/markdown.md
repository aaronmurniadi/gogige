|  GENICAM |   | ![img-66.jpeg](img-66.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

therefore needs to get notified about the schema version associated with the XML being used to access the camera.

If the two components are part of the same software package, the client can easily inform the IPort implementation about the schema version used.

If the two components are communicating through the GenTL interface (IPort implementation resides in GenTL Producer, XML retrieval and parsing in the GenTL Consumer), they need to cooperate according following rules:

- GenTL producers handling other than GigE Vision cameras don't need any specific considerations.
- GenTL producers (GenTL version 1.1 and newer) handling GigE Vision cameras must publish enumeration feature "DeviceEndianessMechanism" in the XML file associated with the GenTL device module. This enumeration must provide two entries, "Standard" and "Legacy". When the consumer selects "Standard", the producer has to implement the GCReadPort/GCWritePort functions of the associated remote device port in the standard way corresponding with schema version 1.1. When the consumer selects "Legacy", the producer has to implement the GCReadPort/GCWritePort of the associated remote device port in the legacy way corresponding with schema version 1.0.
- GenTL consumers (GenTL version 1.1 and newer) accessing a GigE Vision device must instantiate a node map of the GenTL device module and set the DeviceEndianessMechanism feature properly before any access (read/write) to the port of the remote device. The DeviceEndianessMechanism feature must be set to "Standard" when the XML file used for the remote device is based on schema version 1.1 or newer. It has to be set to "Legacy" when the XML file is based on schema version 1.0.
- The enumeration feature DeviceEndianessMechanism and corresponding rules will be standardized by the next GenTL standard version 1.1 (not yet released when publishing this document). GenTL 1.0 does not address this issue. However, GenTL 1.0 compliant producers and consumers willing to fully support little endian GigE Vision cameras are free to implement the same functionality.

### 3.2 Default values of optional Node elements and attributes

The table below contains all node elements and attributes which are an optional part of the available node types in the device XML file. These node elements or attributes needs a definition of its default value even if they are not present in the device XML file to ensure the same behavior in different GenApi implementations.