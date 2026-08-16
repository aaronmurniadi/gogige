|  GevDeviceModeCharacterSet | O | IEnumeration | R | - | I | This feature is deprecated (See DeviceCharacterSet).  |
| --- | --- | --- | --- | --- | --- | --- |
|  GevPhysicalLinkConfiguration | O | IEnumeration | R/W | - | E | Controls the principal physical link configuration to use on next restart/power-up of the device.  |
|  GevCurrentPhysicalLinkConfiguration | O | IEnumeration | R | - | E | Indicates the current physical link configuration of the device.  |
|  GevActiveLinkCount | O | IInteger | R | - | I | Indicates the current number of active logical links.  |
|  GevSupportedOptionSelector | O | IEnumeration | R/W | - | E | Selects the GEV option to interrogate for existing support.  |
|  GevSupportedOption[GevSupportedOptionSelector] | O | IBoolean | R | - | E | Returns if the selected GEV option is supported.  |
|  GevInterfaceSelector | O | IInteger | R/W | - | B | Selects which logical link to control.  |
|  GevLinkSpeed[GevInterfaceSelector] | O | IInteger | R | Mbps | I | This feature is deprecated (See DeviceLinkSpeed).  |
|  GevMACAddress[GevInterfaceSelector] | O | IInteger | R | - | B | MAC address of the logical link.  |
|  GevPAUSEFrameReception[GevInterfaceSelector] | O | IBoolean | R/(W) | - | E | Controls whether incoming PAUSE Frames are handled on the given logical link.  |
|  GevPAUSEFrameTransmission[GevInterfaceSelector] | O | IBoolean | R/(W) | - | E | Controls whether PAUSE Frames can be generated on the given logical link.  |
|  GevCurrentIPConfigurationLLA[GevInterfaceSelector] | O | IBoolean | R/W | - | B | Controls whether the Link Local Address IP configuration scheme is activated on the given logical link.  |
|  GevCurrentIPConfigurationDHCP[GevInterfaceSelector] | O | IBoolean | R/W | - | B | Controls whether the DHCP IP configuration scheme is activated on the given logical link.  |
|  GevCurrentIPConfigurationPersistentIP[GevInterfaceSelector] | O | IBoolean | R/W | - | B | Controls whether the PersistentIP configuration scheme is activated on the given logical link.  |
|  GevCurrentIPAddress[GevInterfaceSelector] | O | IInteger | R | - | B | Reports the IP address for the given logical link.  |
|  GevCurrentSubnetMask[GevInterfaceSelector] | O | IInteger | R | - | B | Reports the subnet mask of the given logical link.  |
|  GevCurrentDefaultGateway[GevInterfaceSelector] | O | IInteger | R | - | B | Reports the default gateway IP address of the given logical link.  |
|  GevIPConfigurationStatus[GevInterface | O | IEnumeration | R | - | B | Reports the current IP configuration status.  |