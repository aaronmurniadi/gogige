|  GEN<i>CAM |   | ![img-11.jpeg](img-11.jpeg) emva  |
| --- | --- | --- |
|  V1.2 | CLProtocol Standard Module  |   |

#### Summary

The following list summarizes the steps a client program has to take in order to select a CLProtocol driver library and identify a camera connected to a frame grabber port.

1. Retrieve a list of PortIDs
2. Present the list of PortIDs to the user to select a frame grabber port for configuration
3. Retrieve a list of DeviceID templates for the selected port
4. Present the list of DeviceID templates to the user to select the best matching template
5. Probe the camera using the selected DeviceID template as a hint. If the camera is recognized a DeviceID is returned unambiguously identifying the camera attached to the selected port
6. Connect to the camera using the DeviceID as identifier.
7. Store the DeviceID for later re-connection.

## 4 Retrieving an XML File for a Camera

Once the CLProtocol driver library is set up and the connection to the camera is established a XML camera description must be retrieved either from the camera or from the file system.

Because there could be more than one matching XML description, e.g. referring to different GenApi schema versions, the standard provides a two step approach for retrieving the XML code: First a sorted list of possible XML descriptions is created, with the best matching description coming first.

Users relying on the automatic just always take the first description to create the GenApi XML node map and configure the camera. If the user wants more control however he can select another XML description manually thus overriding the automatic.

### XML IDs

Each XML description is identified by a XML ID which has the following form:

"SchemaVersion.1.0@<shortDeviceID>@XMLVersion.1.2.3"

The XML ID is composed of three tokens delimited by an at ("@") sign.

The first token describes the version number of the GenApi schema the XML description uses. It has the form

"SchemaVersion.<versionmajor>.<versionminor>"</versionminor></versionmajor>

where <VersionMajor> and <VersionMinor> are integers.

The second token is a short DeviceID template. It thus can have one of the following forms

"Manufacturer"
"Manufacturer#Family"
"Manufacturer#Family#Model"
"Manufacturer#Family#Model#Version"
"Manufacturer#Family#Model#Version#SerialNumber"

The third token describes the version number given in the XML description file for the device. It has the form

"XMLVersion.<versionmajor>.<versionminor>.<versionsubminor>"</versionsubminor></versionminor></versionminor>

where <versionmajor>, <versionminor>, and <versionminor> are integers. Note that the Version from the DeviceID string is an arbitrary CName and not necessarily identical to the version given in the XML. This makes for example sense if a XML file for an existing camera is created stepwise, each step covering more for the camera's functionality while the camera itself is not changing.</versionminor></versionminor>

Here is an example for a XML ID denoting a XML description which is valid for a whole family of cameras

"SchemaVersion.1.1@MyVendor#MyFamily1@XMLVersion.1.2.3.xml"

The list of XML IDs is assembled from the following sources: