|  GENICAM |   | ![img-243.jpeg](img-243.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

command is implemented or not. Conditional Mandatory means that command is to be implemented if possible.

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|  PORT_INFO_ID | M | 0 | Unique ID of the module the port references.In case of the remote device module PORT_INFO_ID returns the same ID as for the local device module.In case of a buffer PORT_INFO_ID returns the address of the buffer as hex string without the leading ‘0x’. For composite buffers the address of the first segment is used.Data type: STRING  |
|  PORT_INFO_VENDOR | M | 1 | Port vendor name.In case the underlying module has no explicit vendor the vendor of the GenTL Producer is to be used. In case of a Buffer or a Data Stream the GenTL Producer vendor and model are to be used.Data type: STRING  |
|  PORT_INFO_MODEL | M | 2 | Port model name.The port model references the model of the underlying module. For example if the port is for the configuration of a TLSystem module the PORT_INFO_MODEL returns the model of the TLSystem Module.In case the underlying module has no explicit model, the model of the GenTL Producer is to be used. So in case of a Buffer or a Data Stream the GenTL Producer model is to be used.Data type: STRING  |
|  PORT_INFO_TLTYPE | M | 3 | Transport layer technology that is supported. See string constants in chapter 6.6.1.Data type: STRING  |
|  PORT_INFO_MODULE | M | 4 | GenTL Module the port refers to:“TLSystem” for the System module.“TLInterface” for the Interface module.“TLDevice” for the Device module.  |