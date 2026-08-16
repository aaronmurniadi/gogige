|  ![img-37.jpeg](img-37.jpeg) CAM |   | ![img-38.jpeg](img-38.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 4.1.2.1 Module Register Map (Recommended)

A URL in the form “local:[///]filename.extension;address;length[?SchemaVersion=x.x.x]” indicates that the XML description file is located in the module’s virtual register map. The square brackets are optional. The “x.x.x” stands for the SchemaVersion the referenced XML complies to in the form major.minor.subminor. If the SchemaVersion is omitted the URL references to an XML referring to SchemaVersion 1.0.0. This optional SchemaVersion is only to be used with the legacy function GCGetPortURL. For current implementations the GCGetPortURLInfo function is used to obtain the SchemaVersion for a specific XML file. Optionally the “///” behind “local:” can be omitted to be compatible to the GigE Vision local format.

If the XML description is stored in the local register map the document can be read by calling the GCReadPort function.

Entries in italics must be replaced with actual values as follows:

Table 4-1: Local URL definition for XML description files in the module register map

|  Entry | Description  |
| --- | --- |
|  local | Indicates that the XML description file is located in the virtual register map of the module.  |
|  filename | Information file name. It is recommended to put the vendor, model/device and version information in the file name separated by an underscore. For example: ‘tlguru_system_rev1’ for the first version of the System module file of the GenTL Producer company TLGuru.  |
|  extension | Indicates the file type. Allowed types are• ‘xml’ for an uncompressed XML description file.• ‘zip’ for a zip-compressed XML description file.  |
|  address | Start address of the file in the virtual register map. It must be expressed in hexadecimal form without a prefix.  |
|  length | Length of the file in bytes. It must be expressed in hexadecimal form without a prefix.  |
|  SchemaVersion | Version the referenced XML complies to. The version is specified in major.minor.subminor format. This only concerns the legacy GCGetPortURL function since the legacy mechanism has no other ways to report a SchemaVersion for the XML file. For the new GCGetPortURLInfo function the SchemaVersion should be retrieved through the info commands.  |

A complete local URL would look like this:

local:tlguru_system_rev1.xml;F0F00000;3BF?SchemaVersion=1.0.0