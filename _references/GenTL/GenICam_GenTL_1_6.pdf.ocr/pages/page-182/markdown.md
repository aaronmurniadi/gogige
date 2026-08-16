|  ![img-274.jpeg](img-274.jpeg)CAN |   | ![img-275.jpeg](img-275.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Name | Interface | Access | Description  |
| --- | --- | --- | --- |
|  InterfaceID | IString | R | GenTL Producer wide unique identifier of the selected interface.  |
|  InterfaceType | IEnumeration | R | Identifies the transport layer technology of the interface. See chapter 6.6.1 for possible values.  |

Table 7-8: Device enumeration features

|  Name | Interface | Access | Description  |
| --- | --- | --- | --- |
|  DeviceUpdateList | ICommand | (R)/W | Updates the internal device list.This feature should be readable if the execution cannot be performed immediately. The command then returns and the status can be polled. This function interacts with the IFUpdateDeviceList function of the GenTL Producer. It is up to the GenTL Consumer to handle access in case both methods are used.  |
|  DeviceSelector | IInteger | R/W | Selector for the different devices on this interface.This value only changes on execution of “DeviceUpdateList”.The selector is 0 based in order to match the index of the C interface.  |
|  DeviceID[DeviceSelector] | IString | R | Interface wide unique identifier of the selected device.This value only changes on execution of “DeviceUpdateList”.  |
|  DeviceVendorName[DeviceSelector] | IString | R | Name of the device vendor.This value only changes on execution of “DeviceUpdateList”.  |
|  DeviceModelName[DeviceSelector] | IString | R | Name of the device model.This value only changes on execution of “DeviceUpdateList”.  |
|  DeviceAccessStatus[DeviceSelector] | IEnumeration | R | Returns the device's access status.Possible values are:“ReadWrite”The device is available to be opened with full access. As soon as the device is open the value should change to “OpenReadWrite” or “OpenRead”Corresponds toDEVICE ACCESS STATUS REA  |