|  GEN<i>CAM |   | ![img-15.jpeg](img-15.jpeg) emva  |
| --- | --- | --- |
|  V1.2 | CLProtocol Standard Module  |   |

Is the connected camera and the CLProtocol driver library capable to deal with device events the function clpGetEventData is used. The library has to put received device events into an event queue. A call to clpGetEventData delivers the next event from that queue and copies the event data into a user allocated buffer.

The CLProtocol module (CCLPort wrapper class) must provide a check if a CLProtocol driver library supports this functionality.

#### Error Handling

Each call to one of the CLProtocol driver library functions returns an error code which normally will be CL_ERR_NO_ERR (=0). The error codes can origin from different places each living in a separate number range. A negative number indicates an error, a positive number a success.

■ Standard error codes from the CLSerXXX interface definition: ±10***
■ Standard error codes from the CLProtocol interface definition: ±20***
■ Custom error codes from the CLProtocol implementations: ±30***

All other numbers are reserved.

The CLProtocol driver library implements the function clpGetErrorText which when given a custom error code must return an error description message in English language. A similar function is also implemented by the CLSerXXX modules. If the CLProtocol module (CCLPort wrapper class) receives a negative return code it first asks the CLProtocol driver library for an error description text. If that call does not return a valid error message, it calls CLAllSerial module which in turn asks the CLSerXXX module and finally it tries to look-up a message text in a list of standard error messages.

#### Interface Version

In order to prepare for future extensions of the CLProtocol driver library the function clpGetCLProtocolVersion must be implemented returning the major and minor version number of the interface. Different major version numbers make two protocols incompatible. A higher minor version number makes the interface backwards compatible to one with a lower minor version.

The current version number is major.minor = 1.2.

#### Setting and Getting Properties of the CLProtocol driver library (v1.1)

The CLPRotocol driver library itself can have properties which can be get and set by the function pair clpGetParam and clpSetParam. The properties available are defined by the enumeration CLP_PARAMS.

CLP_LOG_LEVEL and CLP_LOG_CALLBACK are used to set a logging target and a log level. This allows feeding debug messages into GenICam's standard logging system.

CLP_STOP_PROBE_DEVICE is used to abort a running probe procedure (v1.2).

CLP_DEVICE_BAUDERATE and CLP_DEVICE_SUPPORTED_BAUDERATES allow setting the baud rate of the camera as well as the frame grabber (see ISerial interface).

#### Initializing the CLProtocol driver library (v1.1)

The functions clpInitLib and clpCloseLib are called after loading the CLProtocol driver library and shortly before unloading it.