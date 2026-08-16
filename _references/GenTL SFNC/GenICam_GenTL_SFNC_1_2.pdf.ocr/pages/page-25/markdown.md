![img-6.jpeg](img-6.jpeg)

|  ubnetSelector] |  |  |  |  |  |  |   |
| --- | --- | --- | --- | --- | --- | --- | --- |

### 2.2.2 Device Enumeration

Contains the features related to the enumeration of available Device modules within a specific Interface module.

Table 2-6: Device Enumeration Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  DeviceEnumeration | R | All | ICategory | R | - | E | Category that contains all Device Enumeration features of the Interface module.  |
|  DeviceUpdateList | M | All | ICommand | (R)/W | - | E | Updates the internal device list.  |
|  DeviceUpdateTimeout | R | All | IInteger | R/W | ms | E | Specifies timeout for the DeviceUpdateList Command.  |
|  DeviceSelector | M | All | IInteger | R/W | - | E | Selector for the different devices on this interface.  |
|  DeviceID[DeviceSelector] | M | All | IString | R | - | E | Interface wide unique identifier of the selected device.  |
|  DeviceVendorName[DeviceSelector] | M | All | IString | R | - | E | Name of the device vendor.  |
|  DeviceModelName[DeviceSelector] | M | All | IString | R | - | E | Name of the device model.  |
|  DeviceAccessStatus[DeviceSelector] | M | All | IEnumeration | R | - | E | Gives the device's access status at the moment of the last execution of the DeviceUpdateList command.  |
|  DeviceSerialNumber[DeviceSelector] | R | All | IString | R | - | E | Serial number of the remote device.  |
|  DeviceUserID[DeviceSelector] | O | All | IString | R | - | E | User-programmable device identifier of the remote device.  |
|  DeviceTLVersionMajor[DeviceSelector] | M | All | IInteger | R | - | E | Major version number of the transport layer specification the remote device complies with.  |
|  DeviceTLVersionMinor[DeviceSelector] | M | All | IInteger | R | - | E | Minor version number of the transport layer specification the remote device complies with.  |
|  GevDeviceIPAddress[DeviceSelector] | M | GEV | IInteger | R | - | E | Current IP address of the GVCP interface of the selected  |