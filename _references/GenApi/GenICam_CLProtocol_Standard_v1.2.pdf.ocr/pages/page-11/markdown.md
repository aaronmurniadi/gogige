|  GEN<i>CAM |   | ![img-13.jpeg](img-13.jpeg) emva  |
| --- | --- | --- |
|  V1.2 | CLProtocol Standard Module  |   |

GenICam therefore defines the BaudRate as standard feature which must be implemented by the CLProtocol driver library. Besides of the standard baud rates 9600, 19200 etc. a special AutoMax baud rate can be optionally implemented which is the maximum baud rate the camera and the frame grabber can run with.

The CLProtocol driver library should implement baud rate auto detection, i.e. being able to identify the camera's baud rate during probing.

There must be two ways to access the baud rate.

1. The CLProtocol driver library must implement pseudo registers which are accessible via the IPort interface so the baud rate is exposed via the camera's GenICam XML file.
2. The CLProtocol driver library must also expose the baud rate via the DLL properties which have been introduced in v1.1. This is required so that the CLProtocol module (CCLPort wrapper class) can boost the baud rate while downloading the XML file from the camera.

## 6 Standardized Programming Interfaces

The CLProtocol driver library must implement a set of C functions. The necessary header files are part of the standard.

- CLProtocol.h – declares the C functions to be implemented by the CLProtocol driver library
- CLSerialTypes.h – declares some types and constants
- ISerial.h – declares an abstract C++ interface ISerial which is used by the CLProtocol driver library to access the serial port. A C alias of the virtual function table formed by the C++ interface is also given so an implementation of the CLProtocol driver library in pure C is possible.

These header files contain a detailed description of the functions and their parameters which can be extracted using DoxyGen \( ^{\dagger} \) . This section gives an overview and explains how the functions are used.

### 6.1 ISerial Interface

The CLProtocol driver library needs to have access to the frame grabber's serial port. This is given by a pointer to an ISerial interface which contains the following methods:

- clSerialRead – use this method to retrieve an array of bytes from the camera with timeout. The functionality and parameters are the same as with the corresponding function of the Camera Link standard.
- clSerialWrite – use this method to send an array of bytes to the camera with timeout.
- clGetSupportedBaudRates – this method provides the set of baud rates supported by the frame grabber board in form of a bit field.
- clSetBaudRate – this method sets the baud rate of the frame grabber board

The functionality and parameters of the four methods listed above are the same as with the corresponding function of the Camera Link standard. Because boards supporting only Camera Link v1.0 must be supported no advanced functions like GetNumBytesAvail can be supported.

### 6.2 CLProtocol Interface

The functions to be implemented by the CLProtocol driver library are explained along the use cases introduced in the previous sections. Note that the client of the interface described here is a CLProtocol module. The wrapper class CCLPort contained in the GenICam reference implementation is an example implementation of a CLProtocol module and not the end user's code. However since the wrapper class is not part of the standard it would be possible by a user to write their own client code from scratch.

#### Retrieving a List of DeviceID Templates