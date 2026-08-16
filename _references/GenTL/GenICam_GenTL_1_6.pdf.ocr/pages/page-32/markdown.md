|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

The URL containing the location of the according GenICam XML description can be retrieved through calls to the GCGetNumPortURLs and GCGetPortURLInfo functions of the C interface.

Additional information about the actual port implementation in the GenTL Producer can be retrieved using GCGetPortInfo. The information includes for example the port endianness or the allowed access (read/write, read only,...).

Two modules are special in the way the Port access is handled and are decribed in the following chapters.

##### 4.1.1.1 Device Module

In the Device module two ports are available:

- First the Port functions can be used with a DEV_HANDLE giving access to the Device module's internal features.
- Second the GenTL Consumer can get the PORT_HANDLE of the remote device by calling the DevGetPort function.

Both Ports are mandatory for a GenTL Producer implementation.

##### 4.1.1.2 Buffer Module

The implementation of the Port functions is not mandatory for buffers. To check if an implementation is available call the GCGetPortInfo function with, e.g., the PORT_INFO_MODULE command. If no implementation is present the function's return value must be GC_ERR_NOT_IMPLEMENTED.

#### 4.1.2 XML Description

The last thing missing to be able to use the GenApi like feature access is the XML description. To retrieve a list with the possible locations of the XML the GCGetNumPortURLs function and the GCGetPortURLInfo function can be called. Three possible locations are defined in a URL like notation (for a definition on the URL see RFC 3986):

- Module Register Map (recommended for GenTL Producer)
- Local Directory
- Vendor Web Site

A GenTL Consumer is required to implement ‘Module Register Map’ and ‘Local Directory’. The download from a vendor’s website is optional.

Supported formats are:

• Uncompressed XML description files
- Zip-compressed XML description files. The compression methods used are DEFLATE and STORE as described in RFC 1951.