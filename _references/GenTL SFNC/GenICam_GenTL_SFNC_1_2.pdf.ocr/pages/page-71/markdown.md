|  GEN<i>CAM |   | ![img-19.jpeg](img-19.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Application switchover key to use when requesting ControlAccess switchover. Setup of the key for switchover is done via device features in the device by a host connected in ExclusiveAccess mode.

3.2.2.20 GevDeviceForceIP

|  Name | GevDeviceForceIP[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Apply the force IP settings (GevDeviceForceIPAddress, GevDeviceForceSubnetMask and GevDeviceForceGateway) in the Device using ForceIP command.

This command is only accepted by a device showing ReadWrite DeviceAccessStatus. The IP change is not persistent in the device.

3.2.2.21 GevDeviceForceIPAddress

|  Name | GevDeviceForceIPAddress[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Static IP address to set for the GVCP interface of the remote device.