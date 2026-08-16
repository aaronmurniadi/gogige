|  GEN<i>CAM |   | ![img-10.jpeg](img-10.jpeg) emva  |
| --- | --- | --- |
|  V1.2 | CLProtocol Standard Module  |   |

This key contains a value named “CLSERIALPATH” with type string (REG_SZ) which contains the actual path to the directory. The path should be:

%ProgramFiles%\CameraLink\Serial

If the key/value already exist and point to a different location this location must be used.

For 64-bit Windows the Win32 version of CLSerXXX module should be in the directory defined in the registry key:

HKEY_LOCAL_MACHINE\software\Wow6432Node\cameraLink

NOTE: CLAllSerial module always uses HKEY_LOCAL_MACHINE\software\cameraLink to retrieve the CLSerXXX modules. The Windows Registry Redirector makes sure that the application sees only one of the two registry entries depending on the application being built for Win32 or Win64.

This key contains a value named “CLSERIALPATH” with type string (REG_SZ) which contains the actual path to the directory. The path should be

%ProgramFiles(x86)%\CameraLink\Serial

If the key/value already exist and point to a different location this location must be used. You must not change any existing value.

You must not change any existing value. If the keys/values/directories do not exist they must be created.

The value “CLSERIALPATH” must only contain one path.

When CLAllSerial module loads under Linux, it will search for the CLSerXXX modules in the directory of the CLALLSerial module itself. To work with existing installations of frame grabber manufacturers it is recommended to create a soft link in the directory of the CLAllSerial module which points to the real (lib)clserxxx.so laying in the installation path of the frame grabber software.

Example when using it with soft links:

- /opt/ManufacturerXXX/libclserXXX.so → CLSerXXX module of ManufacturerXXX
- ~/example/clserXXX.so → soft link which points to /opt/ManufacturerXXX/libclserXXX.so
- ~/example/libCLAllSerial_gcc48_v3_0.so → CLAllSerial module
- ~/example/libCLProtocol_gcc48_v3_0.so → CLProtocol module

A PortID is a string of the following form:

"FrameGrabberManufacturer#PortName"

The token on the left of the hash (‘#’) sign is the frame grabber’s manufacturer name and the token to the right the port name. Both strings are retrieved via the clGetPortInfo function defined in the Camera Link standard.

If a CompanyZ has for example two frame grabbers installed in a system with two serial ports each the following list of PortIDs would be result:

"CompanyZ#BoardAPort1"
"CompanyZ#BoardAPort2"
"CompanyZ#BoardBPort1"
"CompanyZ#BoardBPort2"

The standard COM ports of a PC are available via a pseudo frame grabber manufacturer called "COM_Port" enumerating PortIDs of the following form:

"COM_Port#COM1"
"COM_Port#COM2"
etc.

The COM_Port frame grabber library comes as part of the reference implementation.

Another pseudo frame grabber is available named “Local” which is used for ISerial implementations provided statically without using the enumeration mechanism of the CLAllSerial module. This may for example be used in embedded systems. In this case a PortID could for example look like this:

"Local#TheOneAndOnlyPort"