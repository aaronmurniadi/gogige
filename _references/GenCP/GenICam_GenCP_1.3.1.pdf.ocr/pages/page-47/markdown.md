|  GEN<i>CAM |   | ![img-50.jpeg](img-50.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

5.4. Register Map

|  Width (Bytes) | Offset (Bytes) | Support | Access | Description  |
| --- | --- | --- | --- | --- |
|  4 | 0x00000 | M | R | GenCP VersionComplying GenCP specification Version  |
|  64 | 0x00004 | M | R | Manufacturer NameString containing the self-describing name of the manufacturer  |
|  64 | 0x00044 | M | R | Model NameString containing the self-describing name of the device model  |
|  64 | 0x00084 | CM | R | Family NameString containing the name of the family of this device  |
|  64 | 0x000C4 | M | R | Device VersionString containing the version of this device  |
|  64 | 0x00104 | M | R | Manufacturer InfoString containing additional manufacturer information  |
|  64 | 0x00144 | M | R | Serial NumberString containing the serial number of the device  |
|  64 | 0x00184 | CM | RW | User Defined NameString containing the user defined name of the device  |
|  8 | 0x001C4 | M | R | Device CapabilityBit field describing the device's capabilities  |
|  4 | 0x001CC | M | R | Maximum Device Response TimeMaximum response time in ms  |
|  8 | 0x001D0 | M | R | Manifest Table AddressPointer to the Manifest Table  |
|  8 | 0x001D8 | CM | R | SBRM AddressPointer to the Technology Specific Bootstrap Register Map  |
|  8 | 0x001E0 | M | RW | Device ConfigurationBit field describing the device's configuration  |