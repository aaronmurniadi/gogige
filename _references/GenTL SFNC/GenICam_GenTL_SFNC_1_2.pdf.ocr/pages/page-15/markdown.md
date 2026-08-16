|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|   | communicate events to the Device.  |
| --- | --- |
|  Control Channel | A Communication Channel used to configure and control a Device. For a Control Channel the Device acts as a server that provides the initial point of Communication for the Host System that acts as a Client. The Communication on a Control Channel is bidirectional and initiated by the Host System.  |
|  Event | An asynchronous notification of the occurrence of a fact. Events are transmitted on an Event Channel.  |

### 1.3.1 Events in GenTL

Events in GenICam are used for asynchronous signaling between entities, such as the device signaling to the host application. This is described and exemplified in the GenICam and SFNC documents.

In GenTL each module in the producer has the ability to implement events to the application (consumer). Therefore the feature lists in this document includes description of the event mechanism for each module, even if for some modules no predefined events are included.

Typical events from producer to consumer in GenTL give information about the device(s), e.g. new devices are available or a connected device becomes unavailable, or data stream information such as the arrival of a new buffer.

### 1.3.2 Feature Persistence in GenTL

GenICam Feature Persistence is handled outside of the module whos features to persist. In devices this use the defined standard feature DeviceFeaturePersistenceStart to announce that features are to be read from the device, and the feature DeviceFeaturePersistenceEnd to announce that reading of features for persistence has ended. Between these, the persistence algorithm should read all streamable features.

Likewise, DeviceRegistersStreamingStart is used to announce writing of streamed features without validation, and DeviceRegistersStreamingEnd to end this mode, validate the current feature set and update DeviceRegistersValid. The current persistence algorithm in the GenAPI reference implementation uses these standard features.

These features can be used inside GenTL modules to facilitate persistence even though the GenTL modules are not devices.. The persistence features are not included in the features listed in this specification.