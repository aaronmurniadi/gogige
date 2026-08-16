|  ![img-266.jpeg](img-266.jpeg)CAN |   | ![img-267.jpeg](img-267.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

#### 6.5.3 Port Structures

##### 6.5.3.1 PORT_REGISTER_STACK_ENTRY

struct PORT_REGISTER_STACK_ENTRY

Layout of the array elements being used in the function GCWritePortStacked and GCReadPortStacked to accomplish a stacked register read/write operations.

|  Member | Type | Description  |
| --- | --- | --- |
|  Address | uint64_t | Register address  |
|  Buffer | void * | Pointer to the buffer receiving the data being read/containing the data to write.  |
|  Size | size_t | Number of bytes to read / write. The provided Buffer must be at least that size.  |

### 6.6 String Constants

#### 6.6.1 Transport Layer Types

String constants for transport layer technologies that are supported.to be used with the module info commands xxx_INFO_TLTYPE (for example TL_INFO_TLTYPE) inquiry commands.

|  Transport Technology Standard | String Constant  |
| --- | --- |
|  GigE Vision | “GEV”  |
|  Camera Link | “CL”  |
|  IIDC 1394 | “IIDC  |
|  USB video class | “UVC”  |
|  CoaXPress | “CXP”  |
|  Camera Link HS | “CLHS”  |
|  USB3 Vision Standard | “U3V”  |
|  Generic Ethernet | “Ethernet”  |
|  PCI / PCIe | “PCI”  |
|  Mixed | “Mixed”This type is only valid for the System module in case the different Interface modules with a single system are of different types. All other modules must be of a defined type.  |
|  Non standard transport technology, not covered by other constants. | “Custom“  |