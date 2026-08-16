|  GEN<i>CAM |   | ![img-5.jpeg](img-5.jpeg) emva  |
| --- | --- | --- |
|  V1.2 | CLProtocol Standard Module  |   |

also uses the CLAllSerial / CLSerXXX mechanism* defined in chapter 3 to communicate with the camera. Note that the wrapper class is not part of the standard. For more details refer to the CLProtocol tutorial coming with the reference implementation.

The C-interface of the CLProtocol driver library is not operating system dependent, however the compiled library is. Currently the following operating systems are supported:

■ Windows XP (Win32) or higher : this operating system MUST be supported
■ Windows XP (Win64) or higher : this operating system SHOULD be supported
■ Linux 32 bit : this operating system CAN be supported
■ Linux 64 bit : this operating system CAN be supported

In order to setup and use a CLProtocol driver library the following steps must be performed:

1. The CLProtocol driver libraries and any accompanying XML files must be installed and registered in the system
2. For each frame grabber port the right CLProtocol driver library must be selected and the camera must be identified
3. A camera description XML file must be retrieved

These steps are described in the following sections. In addition the ISerial interface and the C functions forming the interface of the CLProtocol driver library are explained. Finally, the properties CLP_DEVICE_BAUDERATE and CLP_DEVICE_SUPPORTED_BAUDERATES are defined which must be implemented by the CLProtocol driver library in order to allow a generic mechanism to connect to setup the baud rate and the camera.

## 2 Installing and Registering CLProtocol Driver Libraries

The CLProtocol driver libraries and any corresponding XML files can be installed by the camera vendor's setup program to an arbitrary location on the target machine, e.g.:

c:\program files\MyVendor\CLProtocol

XML files accompanying the driver libraries are installed directly to that location. For each supported operating system there is a separate sub-directory with a name defined by this standard were the corresponding driver library must be installed to. The following sub-directory names are defined:

■ For Win32 the sub-directory is Win32_i86
■ For Win64 the sub-directory is Win64_x64
■ For Linux 32 bit the sub-directory is Linux32_i86
■ For Linux 64 bit the sub-directory is Linux64_x64

Here is a Windows example:

c:\program files\MyVendorDir\CLProtocol    # install XML files here
c:\program files\MyVendorDir\CLProtocol\Win32_i86    # install Win32 DLL here
c:\program files\MyVendorDir\CLProtocol\Win64_x64    # install Win64 DLL here

Multiple driver libraries with different names can reside in one sub-directory. The driver library name must be of the form *.dll for Windows and lib*.so for Linux.

The registration is performed by adding the location (i.e. the directory name without trailing backslash) to a list of locations given in the environment variable GENICAM_CLPROTOCOL, for example:

GENICAM_CLPROTOCOL=c:\program files\MyVendorDir\CLProtocol;c:\temp\MyTest