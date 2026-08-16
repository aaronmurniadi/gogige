|  Name |  |  |  |  |  |   |
| --- | --- | --- | --- | --- | --- | --- |
|  oMACControlFunctionEntity | O | ICategory | R | - | G | Category that contains statistics pertaining to the MAC control PAUSE function of the device.  |
|  aPAUSEMACCtrlFramesTransmitted[GevInterfaceSelector] | O | IInteger | R | - | G | Reports the number of transmitted PAUSE frames.  |
|  aPAUSEMACCtrlFramesReceived[GevInterfaceSelector] | O | IInteger | R | - | G | Reports the number of received PAUSE frames.  |

# Camera Link

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  CameraLink | O | ICategory | R | - | B | Category that contains the features pertaining to the Camera Link transport layer of the device.  |
|  ClConfiguration | R | IEnumeration | R/(W) | - | B | This Camera Link specific feature describes the configuration used by the camera.  |
|  ClTimeSlotsCount | O | IEnumeration | R/(W) | - | E | This Camera Link specific feature describes the time multiplexing of the camera link connection to transfer more than the configuration allows, in one single clock.  |

# CoaXPress

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  CoaXPress | O | ICategory | R | - | B | Category that contains the features pertaining to the CoaXPress transport layer of the device.  |
|  CxpLinkConfigurationStatus | R | IEnumeration | R | - | B | This feature indicates the current and active Link configuration used by the Device.  |
|  CxpLinkConfigurationPreferred | R | IEnumeration | R | - | E | Provides the Link configuration that allows the Transmitter Device to operate in its default mode.  |
|  CxpLinkConfiguration | R | IEnumeration | R/W | - | B | This feature allows specifying the Link configuration for the communication between the Receiver and Transmitter Device.  |