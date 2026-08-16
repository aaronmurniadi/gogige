|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

- Dynamic loading of libraries: it is easily possible to dynamically load and call C style functions. This enables the implementation of a GenTL Consumer dynamically loading one or more GenTL Producers at runtime.
- Upgradeability: a C library can be designed in a way that it is binary compatible to earlier versions. Thus the GenTL Consumer does not need to be recompiled if a version change occurs.

Although a C interface was chosen because of the reasons mentioned above, the actual GenTL Producer implementation can be done in an object-oriented language. Except for the global functions, all interface functions work on handles which can be mapped to objects.

Any programming language that can export a library with a C interface can be used to implement a GenTL Producer.

To guarantee interchangeability of GenTL Producers and GenTL Consumers no language specific feature except the ones compatible to ANSI C may be used in the interface of the GenTL Producer.

#### 2.3.2 Configuration

Each module provides GenTL Port functionality so that the GenICam GenApi (or any other similar, non-reference implementations) can be used to access a module's configuration. The basic operations on a GenTL Producer implementation can be done with the C interface without using a specific module configuration. More complex or implementation-specific access can be done via the flexible GenApi Feature interface using the GenTL Port functionality and the provided GenApi XML description.

Each module brings this XML description along with which the module's port can be used to read and/or modify settings in the module. To do that each module has its own virtual register map that can be accessed by the Port functions. Thus, the generic way of accessing the configuration of a remote device has been extended to the transport layer modules themselves.

#### 2.3.3 Signaling (Events)

Each module provides the possibility to notify the GenTL Consumer of certain events. As an example, a New Buffer event can be raised/signaled if new image data has arrived from a remote device. The number of events supported for a specific module depends on the module and its implementation.

The C interface enables the GenTL Consumer to register events on a module. The event object used is platform and implementation dependent, but is encapsulated in the C interface.