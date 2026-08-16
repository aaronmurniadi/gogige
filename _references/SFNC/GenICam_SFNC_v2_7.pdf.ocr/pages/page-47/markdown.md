|   |  |  |  |  |  | used to create the device's GenICam XML.  |
| --- | --- | --- | --- | --- | --- | --- |
|  DeviceSFNCVersionSubMinor | R | Integer | R | - | B | Sub minor version of Standard Features Naming Convention that was used to create the device's GenICam XML.  |
|  DeviceManifestEntrySelector | O | Integer | R/W | - | G | Selects the manifest entry to reference.  |
|  DeviceManifestXMLMajorVersion [DeviceManifestEntrySelector] | O | Integer | R | - | G | Indicates the major version number of the GenICam XML file of the selected manifest entry.  |
|  DeviceManifestXMLMinorVersion [DeviceManifestEntrySelector] | O | Integer | R | - | G | Indicates the minor version number of the GenICam XML file of the selected manifest entry.  |
|  DeviceManifestXMLSubMinorVersion [DeviceManifestEntrySelector] | O | Integer | R | - | G | Indicates the subminor version number of the GenICam XML file of the selected manifest entry.  |
|  DeviceManifestSchemaMajorVersion [DeviceManifestEntrySelector] | O | Integer | R | - | G | Indicates the major version number of the schema file of the selected manifest entry.  |
|  DeviceManifestSchemaMinorVersion [DeviceManifestEntrySelector] | O | Integer | R | - | G | Indicates the minor version number of the schema file of the selected manifest entry.  |
|  DeviceManifestPrimaryURL [DeviceManifestEntrySelector] | O | String | R | - | G | Indicates the first URL to the GenICam XML device description file of the selected manifest entry.  |
|  DeviceManifestSecondaryURL [DeviceManifestEntrySelector] | O | String | R | - | G | Indicates the second URL to the GenICam XML device description file of the selected manifest entry.  |
|  DeviceTLType | R | Enumeration | R | - | B | Transport Layer type of the device.  |
|  DeviceTLVersionMajor | R | Integer | R | - | B | Major version of the Transport Layer of the device.  |
|  DeviceTLVersionMinor | R | Integer | R | - | B | Minor version of the Transport Layer of the device.  |
|  DeviceTLVersionSubMinor | R | Integer | R | - | B | Sub minor version of the Transport Layer of the device.  |
|  DeviceGenCPVersionMajor | R | Integer | R | - | B | Major version of the GenCP protocol supported by the device.  |
|  DeviceGenCPVersionMinor | R | Integer | R | - | B | Minor version of the GenCP protocol supported by the device.  |
|  DeviceMaxThroughput | O | Integer | R | Bps | E | Maximum bandwidth of the data that can be streamed out of the device.  |
|  DeviceConnectionSelector | R | Integer | R/(W) | - | B | Selects which Connection of the device to control.  |
|  DeviceConnectionSpeed [DeviceConnectionSelector] | O | Integer | R | Bps | E | Indicates the speed of transmission of the specified Connection.  |