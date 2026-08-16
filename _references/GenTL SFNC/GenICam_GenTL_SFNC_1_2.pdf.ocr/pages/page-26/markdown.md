|   |  |  |  |  |  |  | remote device.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  GevDeviceSubnetMask[DeviceSelector] | M | GEV | IInteger | R | - | E | Current subnet mask of the GVCP interface of the selected remote device.  |
|  GevDeviceGateway[DeviceSelector] | R | GEV | IInteger | R | - | E | Current gateway IP address of the GVCP interface of the selected remote device.  |
|  GevDeviceIPConfigurationStatus[DeviceSelector] | R | GEV | IEnum | R/W | - | E | Device IP configuration of the GVCP interface of the selected remote device.  |
|  GevDeviceMACAddress[DeviceSelector] | M | GEV | IInteger | R | - | E | 48-bit MAC address of the GVCP interface of the selected remote device.  |
|  GevDeviceCurrentControlMode[DeviceSelector] | O | GEV | IEnum | R/W | - | E | The current control mode of the device.  |
|  GevApplicationSwitchoverKey[DeviceSelector] | O | GEV | IInteger | W | - | E | Application switchover key to use when requesting ControlAccess switchover.  |
|  GevDeviceForceIP[DeviceSelector] | R | GEV | ICommand | (R)/W | - | E | Apply the force IP settings (GevDeviceForceIPAddress, GevDeviceForceSubnetMask and GevDeviceForceGateway) in the Device using ForceIP command.  |
|  GevDeviceForceIPAddress[DeviceSelector] | R | GEV | IInteger | R/W | - | E | Static IP address to set for the GVCP interface of the remote device.  |
|  GevDeviceForceSubnetMask[DeviceSelector] | R | GEV | IInteger | R/W | - | E | Static subnet mask to set for GVCP interface of the remote device.  |
|  GevDeviceForceGateway[DeviceSelector] | R | GEV | IInteger | R/W | - | E | Static gateway IP address to set for the GVCP interface of the remote device.  |