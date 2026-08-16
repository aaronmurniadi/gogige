|  GevSCPHostPort[GevStreamChannelSelector] | O | IInteger | R/W | - | G | Controls the port of the selected channel to which a GVSP transmitter must send data stream or the port from which a GVSP receiver may receive data stream.  |
| --- | --- | --- | --- | --- | --- | --- |
|  GevSCPSFireTestPacket[GevStreamChannelSelector] | O | IBoolean | R/W | - | G | Sends a test packet.  |
|  GevSCPSDoNotFragment[GevStreamChannelSelector] | O | IBoolean | R/W | - | G | The state of this feature is copied into the "do not fragment" bit of IP header of each stream packet.  |
|  GevSCPSBigEndian[GevStreamChannelSelector] | O | IBoolean | R/W | - | I | This feature is deprecated (See DeviceStreamChannelEndianness).  |
|  GevSCPSPacketSize[GevStreamChannelSelector] | R | IInteger | R/(W) | B | E | This GigE Vision specific feature corresponds to DeviceStreamChannelPacketSize and should be kept in sync with it.  |
|  GevSCPD[GevStreamChannelSelector] | R | IInteger | R/W | - | E | Controls the delay (in GEV timestamp counter unit) to insert between each packet for this stream channel.  |
|  GevSCDA[GevStreamChannelSelector] | O | IInteger | R/W | - | G | Controls the destination IP address of the selected stream channel to which a GVSP transmitter must send data stream or the destination IP address from which a GVSP receiver may receive data stream.  |
|  GevSCSP[GevStreamChannelSelector] | O | IInteger | R | - | G | Indicates the source port of the stream channel.  |
|  GevSCZoneCount[GevStreamChannelSelector] | O | IInteger | R | - | G | Reports the number of zones per block transmitted on the selected stream channel.  |
|  GevSCZoneDirectionAll[GevStreamChannelSelector] | O | IInteger | R | - | G | Reports the transmission direction of each zone transmitted on the selected stream channel.  |
|  GevSCZoneConfigurationLock[GevStreamChannelSelector] | O | IBoolean | R/W | - | G | Controls whether the selected stream channel multi-zone configuration is locked.  |

# Network Statistics

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  NetworkStatistics | O | ICategory | R | - | G | Category that contains statistics pertaining to various modules of the GigE Vision transport layer.  |
|   | Level | Interface | Access | Unit | Visibility | Description  |