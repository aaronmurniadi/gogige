|  DeviceInformation | R | All | ICategory | R | - | B | Category that contains all Device Information features of the Device module.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  DeviceID | M | All | IString | R | - | E | Interface-wide unique identifier of this device.  |
|  DeviceSerialNumber | R | All | IString | R | - | E | Serial number of the remote device.  |
|  DeviceUserID | O | All | IString | R/W | - | E | User-programmable device identifier of the remote device.  |
|  DeviceVendorName | M | All | IString | R | - | B | Name of the remote device vendor.  |
|  DeviceModelName | M | All | IString | R | - | B | Name of the remote device model.  |
|  DeviceFamilyName | R | All | IString | R | - | B | Name of the product family of the remote device model.  |
|  DeviceVersion | R | All | IString | R | - | B | The version of the remote device model.  |
|  DeviceManufacturerInfo | R | All | IString | R | - | B | Manufacturer information about the remote device.  |
|  DeviceType | M | All | IEnumeration | R | - | E | Transport layer type of the device.  |
|  DeviceDisplayName | R | All | IString | R | - | E | User readable name of the device.  |
|  DeviceTimestampFrequency | R | All | IInteger | R | - | B | The tick-frequency of the time stamp clock.  |
|  DeviceAccessStatus | M | All | IEnumeration | R | - | E | Gives the device's access status at the moment of the last execution of the DeviceUpdateList command.  |
|  DeviceChunkDataFormat | R | All | IEnumeration | R | - | E | Chunk data format used by the device.  |
|  DeviceEventDataFormat | R | All | IEnumeration | R | - | E | Enumeration, informing about the event data format used by the device (meaning the "device events", see event type EVENT_REMOTE_DEVICE (named EVENT_FEATURE_DEVEVENT in GenTL up to version 1.  |
|  GevDeviceMACAddress | M | GEV | IInteger | R | - | E | 48-bit MAC address of the GVCP interface of the remote device.  |
|  GevDeviceIPAddress | M | GEV | IInteger | R | - | E | Current IP address of the GVCP interface of the remote device.  |