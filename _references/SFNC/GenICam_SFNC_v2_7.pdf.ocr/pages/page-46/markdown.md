|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

## 2 Features Summary

This chapter provides a comprehensive summary of the standard features covered by this document. The following chapters provide more detailed explanation of each feature.

In case of discrepancy, the sections describing the features in detail prevail.

### 2.1 Device Control

Contains the features related to the control and information of the device (See the Device Control chapter for details).

Table 2-1: Device Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  DeviceControl | R | ICategory | R | - | B | Category for device information and control.  |
|  DeviceType | O | IEnumeration | R | - | G | Returns the device type.  |
|  DeviceScanType | R | IEnumeration | R/(W) | - | E | Scan type of the sensor of the device.  |
|  DeviceVendorName | R | ISString | R | - | B | Name of the manufacturer of the device.  |
|  DeviceModelName | R | ISString | R | - | B | Model of the device.  |
|  DeviceFamilyName | O | ISString | R | - | B | Identifier of the product family of the device.  |
|  DeviceManufacturerInfo | R | ISString | R | - | B | Manufacturer information about the device.  |
|  DeviceVersion | R | ISString | R | - | B | Version of the device.  |
|  DeviceFirmwareVersion | R | ISString | R | - | B | Version of the firmware in the device.  |
|  DeviceSerialNumber | R | ISString | R | - | E | Device's serial number.  |
|  DeviceID | R | ISString | R | - | I | This feature is deprecated (See DeviceSerialNumber).  |
|  DeviceUserID | O | ISString | R/W | - | B | User-programmable device identifier.  |
|  DeviceSFNCVersionMajor | R | IInteger | R | - | B | Major version of the Standard Features Naming Convention that was used to create the device's GenICam XML.  |
|  DeviceSFNCVersionMinor | R | IInteger | R | - | B | Minor version of the Standard Features Naming Convention that was  |