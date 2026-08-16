|  GEN<I>CAM |   | ![img-14.jpeg](img-14.jpeg) emva  |
| --- | --- | --- |
|  V1.2 | CLProtocol Standard Module  |   |

The function clpGetShortDeviceIDTemplates is used to collect a list of DeviceID templates. The environment variable GENICAM_CLPROTOCOL contains a list of locations were CLProtocol driver libraries are stored. The CLProtocol module (CCLPort wrapper class) loads each of those driver libraries and calls clpGetShortDeviceIDTemplates retrieving for each driver library a list of short DeviceID templates the respective driver library will understand. The combined short DeviceID templates decorated with the location of the driver library where they originate from for the desired list of DeviceID templates.

Note that for calling clpGetShortDeviceIDTemplates no ISerial interface needs to be supplied.

#### Probing, Identifying and Re-Connecting a Camera

The function clpProbeDevice is called with an ISerial interface and a DeviceID template as input parameter. The function attempts to identify the camera attached to the respective frame grabber port using the DeviceID template as hint. If the function is successful it returns a DeviceID as well as a Cookie (see below).

If the DeviceID is already known the function clpProbeDevice can also be used to re-connect the camera. This is simply done by handing in the DeviceID instead of a DeviceID template. It is the responsibility of the CLProtocol driver library to distinguish between the two use cases. Re-connecting instead of probing again makes sense because re-connecting is normally much faster than probing. Generally an application should probe and identify a camera only once and then store the DeviceID for re-connect.

By calling clpProbeDevice a connection to the camera is opened. This connection is identified by the Cookie which must be handed in for any subsequent calls to the CLProtocol driver library. The driver library may use the Cookie to persist any data while the connection is open. Note that the Cookie value must not be 0.

#### Closing a Connection to the Camera

In order to close the connection to the camera call clpDisconnect handing in the Cookie. On this call the CLProtocol driver library must free all persistent data attached to the connection and the Cookie becomes invalid.

#### Retrieving the XML Description for a Camera

Calling clpGetXMLIDs returns a list of XML ID's exposing what kind of XML descriptions the camera and/or the CLProtocol driver library itself is able to provide. Note that this does not include XML descriptions stored on the file system beside the driver library. These XML IDs belonging to these XML files are added by the CLProtocol module (CCLPort wrapper class).

The XML IDs returned from calling clpGetXMLIDs are not sorted. This is also done by the CLProtocol module (CCLPort wrapper class).

If the user has selected a XML ID originating from the call to clpGetXMLIDs he can retrieve the actual XML description by calling clpGetXMLDescription handing in the XML ID as a parameter. Again, the CLProtocol module (CCLPort wrapper class) handles XML ID's belonging to XML files stored on the file system.

The best matching XML ID is typically only determined once and then stored along with the DeviceID for reconnection. The CLProtocol module (CCLPort wrapper class) caches XML descriptions retrieved during the first connect and thus makes sure that unnecessary XML file downloads from the camera are avoided.

#### Accessing Camera Registers

Camera Registers are read and written to using the functions clpReadRegister and clpWriteRegister. Both function calls require an ISerial interface, a Cookie and a timeout which should by default be set to 500ms.

For commands taking a longer time to complete than the typical timeout the function clpWriteRegister can return a special value CL_ERR_PENDING_WRITE. In this case the client code must call clpContinueWriteRegister which either completes the call or returns CL_ERR_PENDING_WRITE again. The function clpContinueWriteRegister can be called with a cancel flag thus abandoning the command processing. After cancelling a call the camera must be in a state to accept further commands without problem. The wrapper class handles the whole pending business under the hood so the customer will normally not notice it.

#### Event Handling