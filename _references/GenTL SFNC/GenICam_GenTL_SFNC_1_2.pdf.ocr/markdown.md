# GenICam GenTL

# Standard Features
# Naming Convention

Version 1.2.0

GENICAM

# Table of Contents

TABLE OF CONTENTS 2

HISTORY 9

1 INTRODUCTION 13

1.1 GENICAM REFERENCE DOCUMENTS 13

1.2 TL SPECIFIC FEATURES 13

1.3 STANDARD DEFINITIONS 14

1.3.1 Events in GenTL 15

1.3.2 Feature Persistence in GenTL 15

1.4 CONVENTIONS 16

1.5 STANDARD UNITS 18

1.6 ACRONYMS 19

2 FEATURES SUMMARY 20

2.1 SYSTEM MODULE 20

2.1.1 System Information 20

2.1.2 Interface Enumeration 21

2.1.3 GenICam Control 22

2.1.4 Event Control 23

2.2 INTERFACE MODULE 24

2.2.1 Interface Information 24

2.2.2 Device Enumeration 25

2.2.3 Action Control 27

2.2.4 GenICam Control 27

2.2.5 Event Control 28

2.3 DEVICE MODULE 28

2.3.1 Device Information 28

2.3.2 Device Control 31

2.3.3 Stream Enumeration 31

2.3.4 GenICam Control 32

|  ![img-0.jpeg](img-0.jpeg)CAM |   | ![img-1.jpeg](img-1.jpeg)emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

2.3.5 Event Control 32

2.4 DATA STREAM MODULE 33

2.4.1 Stream Information 33

2.4.2 Device Stream Channel Control 33

2.4.3 Buffer Handling Control 34

2.4.4 GenICam Control 35

2.4.5 Event Control 35

2.5 BUFFER MODULE 37

2.5.1 Buffer Information 37

2.5.2 Buffer Data Information 37

2.5.3 GenICam Control 37

3 GENERAL FEATURES 41

3.1 SYSTEM MODULE 41

3.1.1 System Information 41

3.1.1.1 SystemInformation 41

3.1.1.2 TLID 41

3.1.1.3 TLVendorName 42

3.1.1.4 TLModelName 42

3.1.1.5 TLVersion 43

3.1.1.6 TLFileName 43

3.1.1.7 TLDisplayName 43

3.1.1.8 TLPath 44

3.1.1.9 TLType 44

3.1.1.10 GenTLVersionMajor 46

3.1.1.11 GenTLVersionMinor 46

3.1.1.12 GenTLSFNCVersionMajor 47

3.1.1.13 GenTLSFNCVersionMinor 47

3.1.1.14 GevVersionMajor (Deprecated) 47

3.1.1.15 GevVersionMinor (Deprecated) 48

3.1.2 Interface Enumeration 48

3.1.2.1 InterfaceEnumeration 49

3.1.2.2 InterfaceUpdateList 49

3.1.2.3 InterfaceUpdateTimeout 49

3.1.2.4 InterfaceSelector 50

3.1.2.5 InterfaceID 50

|  GEN<i>CAM |   | ![img-2.jpeg](img-2.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

3.1.2.6 InterfaceDisplayName 51
3.1.2.7 GevInterfaceMACAddress 51
3.1.2.8 GevInterfaceDefaultIPAddress 51
3.1.2.9 GevInterfaceDefaultSubnetMask 52
3.1.2.10 GevInterfaceDefaultGateway 52

#### 3.1.3 GenICam Control 53

3.1.3.1 Root 53
3.1.3.2 TLPort 53

#### 3.1.4 Event Control 54

3.1.4.1 EventControl 54
3.1.4.2 EventSelector 54
3.1.4.3 EventNotification 55

### 3.2 INTERFACE MODULE 55

#### 3.2.1 Interface Information 55

3.2.1.1 InterfaceInformation 56
3.2.1.2 InterfaceID 56
3.2.1.3 InterfaceDisplayName 56
3.2.1.4 InterfaceType 57
3.2.1.5 InterfaceTLVersionMajor 58
3.2.1.6 InterfaceTLVersionMinor 58
3.2.1.7 GevInterfaceGatewaySelector 59
3.2.1.8 GevInterfaceGateway 59
3.2.1.9 GevInterfaceMACAddress 59
3.2.1.10 GevInterfaceSubnetSelector 60
3.2.1.11 GevInterfaceSubnetIPAddress 60
3.2.1.12 GevInterfaceSubnetMask 61

#### 3.2.2 Device Enumeration 61

3.2.2.1 DeviceEnumeration 61
3.2.2.2 DeviceUpdateList 62
3.2.2.3 DeviceUpdateTimeout 62
3.2.2.4 DeviceSelector 63
3.2.2.5 DeviceID 63
3.2.2.6 DeviceVendorName 63
3.2.2.7 DeviceModelName 64
3.2.2.8 DeviceAccessStatus 64
3.2.2.9 DeviceSerialNumber 65
3.2.2.10 DeviceUserID 66
3.2.2.11 DeviceTLVersionMajor 66

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

3.2.2.12 DeviceTLVersionMinor 67
3.2.2.13 GevDeviceIPAddress 67
3.2.2.14 GevDeviceSubnetMask 67
3.2.2.15 GevDeviceGateway 68
3.2.2.16 GevDeviceIPConfigurationStatus 68
3.2.2.17 GevDeviceMACAddress 69
3.2.2.18 GevDeviceCurrentControlMode 69
3.2.2.19 GevApplicationSwitchoverKey 70
3.2.2.20 GevDeviceForceIP 71
3.2.2.21 GevDeviceForceIPAddress 71
3.2.2.22 GevDeviceForceSubnetMask 72
3.2.2.23 GevDeviceForceGateway 72

### 3.2.3 Action Control 72

3.2.3.1 ActionControl 73
3.2.3.2 ActionCommand 73
3.2.3.3 ActionDeviceKey 73
3.2.3.4 ActionGroupKey 74
3.2.3.5 ActionGroupMask 74
3.2.3.6 ActionScheduledTimeEnable 75
3.2.3.7 ActionScheduledTime 75
3.2.3.8 GevActionDestinationIPAddress 75

### 3.2.4 GenICam Control 76

3.2.4.1 Root 76
3.2.4.2 InterfacePort 76

### 3.2.5 Event Control 77

3.2.5.1 EventControl 77
3.2.5.2 EventSelector 77
3.2.5.3 EventNotification 78

## 3.3 DEVICE MODULE 79

### 3.3.1 Device Information 79

3.3.1.1 Device Information 79
3.3.1.2 DeviceID 79
3.3.1.3 DeviceSerialNumber 80
3.3.1.4 DeviceUserID 80
3.3.1.5 DeviceVendorName 81
3.3.1.6 DeviceModelName 81
3.3.1.7 DeviceFamilyName 82
3.3.1.8 DeviceVersion 82

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

3.3.1.9 DeviceManufacturerInfo 83

3.3.1.10 DeviceType 83

3.3.1.11 DeviceDisplayName 84

3.3.1.12 DeviceTimestampFrequency 84

3.3.1.13 DeviceAccessStatus 85

3.3.1.14 DeviceChunkDataFormat 86

3.3.1.15 DeviceEventDataFormat 87

3.3.1.16 GevDeviceMACAddress 88

3.3.1.17 GevDeviceIPAddress 88

3.3.1.18 GevDeviceSubnetMask 89

3.3.1.19 GevDeviceGateway 89

3.3.2 Device Control 90

3.3.2.1 DeviceControl 90

3.3.2.2 DeviceEndianessMechanism 90

3.3.2.3 LinkCommandTimeout 91

3.3.2.4 LinkCommandRetryCount 91

3.3.3 Stream Enumeration 92

3.3.3.1 StreamEnumeration 92

3.3.3.2 StreamSelector 92

3.3.3.3 StreamID 92

3.3.4 GenICam Control 93

3.3.4.1 Root 93

3.3.4.2 DevicePort 93

3.3.5 Event Control 94

3.3.5.1 EventControl 94

3.3.5.2 EventSelector 94

3.3.5.3 EventNotification 95

3.4 DATA STREAM MODULE 96

3.4.1 Stream Information 96

3.4.1.1 Stream Information 96

3.4.1.2 StreamID 96

3.4.1.3 StreamType 97

3.4.2 Device Stream Channel Control 98

3.4.2.1 DeviceStreamChannelControl 98

3.4.2.2 DeviceStreamChannelPacketSize 98

3.4.2.3 DeviceStreamChannelPacketSizeMin 99

3.4.2.4 DeviceStreamChannelPacketSizeMax 99

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

3.4.2.5 DeviceStreamChannelPacketSizeInc 100
3.4.2.6 DeviceStreamChannelNegotiatePacketSize 100

### 3.4.3 Buffer Handling Control 101

3.4.3.1 BufferHandlingControl 101
3.4.3.2 StreamAnnouncedBufferCount 101
3.4.3.3 StreamBufferHandlingMode 101
3.4.3.4 StreamAnnounceBufferMinimum 104
3.4.3.5 StreamDeliveredFrameCount 105
3.4.3.6 StreamLostFrameCount 105
3.4.3.7 StreamInputBufferCount 105
3.4.3.8 StreamOutputBufferCount 106
3.4.3.9 StreamStartedFrameCount 106
3.4.3.10 PayloadSize 107
3.4.3.11 StreamIsGrabbing 107
3.4.3.12 StreamChunkCountMaximum 108
3.4.3.13 StreamBufferAlignment 108

### 3.4.4 GenICam Control 109

3.4.4.1 Root 109
3.4.4.2 StreamPort 109

### 3.4.5 Event Control 110

3.4.5.1 EventControl 110
3.4.5.2 EventSelector 110
3.4.5.3 EventNotification 111

### 3.5 BUFFER MODULE 112

#### 3.5.1 Buffer Information 112

3.5.1.1 BufferInformation 112
3.5.1.2 BufferUserData 113
3.5.1.3 BufferType 113
3.5.1.4 BufferSize 114

#### 3.5.2 Buffer Data Information 114

3.5.2.1 BufferDataInformation 115
3.5.2.2 BufferData 115
3.5.2.3 BufferTimeStamp 115
3.5.2.4 BufferNewData 116
3.5.2.5 BufferIsQueued 116
3.5.2.6 BufferIsAcquiring 117
3.5.2.7 BufferIsIncomplete 117

|  GEN<i>CAM |   | ![img-3.jpeg](img-3.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

3.5.2.8 BufferPayloadType 118
3.5.2.9 BufferNumberOfParts 119
3.5.2.10 BufferPartSelector 119
3.5.2.11 BufferSizeFilled 120
3.5.2.12 BufferPartDataType 120
3.5.2.13 BufferPartSourceIDValue 121
3.5.2.14 BufferPartRegionIDValue 121
3.5.2.15 BufferPartComponentIDValue 122
3.5.2.16 BufferWidth 122
3.5.2.17 BufferHeight 123
3.5.2.18 BufferXOffset 123
3.5.2.19 BufferYOffset 124
3.5.2.20 BufferXPadding 124
3.5.2.21 BufferYPadding 125
3.5.2.22 BufferFrameID 125
3.5.2.23 BufferImagePresent 126
3.5.2.24 BufferImageOffset 126
3.5.2.25 BufferPixelFormat 127
3.5.2.26 BufferDeliveredImageHeight 130
3.5.2.27 BufferDeliveredChunkPayloadSize 130
3.5.2.28 BufferChunkLayoutID 131
3.5.2.29 BufferFileName 131

#### 3.5.3 GenICam Control 132

3.5.3.1 Root 132

3.5.3.2 BufferPort 132

## 4 ACKNOWLEDGEMENTS 134

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

# History

|  Version | Date | Changed by | Change  |
| --- | --- | --- | --- |
|  0.1 | 04.10.2010 | Christoph Zierl, MVTec | First draft version based on Chapter 7 of the GenICam GenTL standard v1.2 and the feature collection at the GenICam Wiki.  |
|  0.2 | 29.09.2011 | Christoph Zierl, MVTec | - Adapted to changes in GenTL v1.3 RC2 - Added additional features corresponding to INFO_CMD enumerations - Added first round of CXP features  |
|  0.3 | 05.09.2012 | Christoph Zierl, MVTec Jan Becvar, Groget | - General review - Introduced feature categories - Reviewed feature visibility - Adapted CXP features to new proposal from CoaXPress group for SFNC 2.0 - Added all missing features corresponding to STREAM_INFO_CMD and BUFFER_INFO_CMD enumerations  |
|  RC1 | 05.12.2012 | Christoph Zierl, MVTec | - Fixed erroneous name of IFUpdateDeviceList function - Updated value list for TL/Interface/Device/Stream/BufferType features according to new value list defined in SFNC 2.0 - Updated interface type and value list for BufferPixelFormat feature according to new value list defined in SFNC 2.0 / PFNC 1.0 - Added features DeviceFamilyName, DeviceVersion, DeviceFirmwareVersion corresponding to SFNC 2.0 and GenCP 1.0 - Added features U3vVersionMajor and U3vVersionMinor - Renamed 'GenICam Access' categories to 'GenICam Control' - Updated text regarding buffer handling modes - Updated introduction text in Chapter 1  |
|  RC2 | 29.01.2013 | Christoph Zierl, MVTec | - Removed CxpVersionMajor/Minor and U3vVersionMajor/Minor features in accordance to SFNC 2.0 - Corrected entries in standard units table - Improved description of DeviceChunkDataFormat feature - Fixed typo in DeviceEventDataFormat - Updated acknowledgements table  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Version | Date | Changed by | Change  |
| --- | --- | --- | --- |
|  RC3 | 11.03.2013 | Christoph Zierl, MVTec Jan Becvar, Groget | - Removed erroneous '[DeviceSelector]' from DeviceSerialNumber and DeviceUserID features in section 3.3.1 'DeviceInformation' - Moved double occurrence of GevDeviceMACAddress feature from section 3.3.1 'Device Information' to section 3.2.2 'Device Enumeration' - Renamed CxpPoCxpAuto feature into CxpPoCxpSetAuto and CxpPoCxpOff feature into CxpPoCxpTurnOff following the input from CoaXPress liaison group - Revised description and fixed typos regarding CoaXPress features - Renamed 'Recommended Visibility' into 'Visibility' in all feature tables - Improved overall formatting, in particular to enable the automatic generation of the reference XML files - Updated acknowledgements table  |
|  RC4 | 25.03.2013 | Christoph Zierl, MVTec | - Fixed inconsistent naming of feature TLFileName - Removed all CXP features since it is not yet decided whether it actually makes sense to copy these feature definitions from the regular SFNC document to the GenTL SFNC document. Note that this does not affect the binding character of these features for the CoaXPress standard itself!  |
|  1.0 | 06.05.2013 | GenICam Committee | Final Release v1.0  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Version | Date | Changed by | Change  |
| --- | --- | --- | --- |
|  1.1 | 2007-02-03 | Mattias Johannesson, SICK | For GenTL 1.5. Collected from Discussion Topic #50. - Standards Definitions from SFNC 2.3 section 1.4 included. Acronyms not in document removed. - Timeout added for Interface/DeviceUpdateList commands. - Added GigE Vision IP & Control setup features - Added NewestOnly buffer mode - Added Generic TL version features, deprecated GigEs. - Added Action Command support - DeviceAccessStatus enum aligned with GenTL for device and interface - GigE specific parts moved to be in same position as generic for ease of use reading the document. - Added features for Multipart buffers introduced in GenTL 1.5 and proposed for GigE Vision 2.1. - InterfaceDisplayName Added - Interface Enumeration Category and some features changed to Beginner level. - TLType enum contents that were deprecated in 1.0 are kept to not break compatibility backwards. - Features for LinkCommand timeout and retries added. - Adjusted visibility levels for consistency to have category at same level as lowest visible feature. - Added DeviceTimestampFrequency to complete mapping of DEVICE_INFO_XX features, and references all of them correctly. - EventControl added as per ticket #1305. - Deprecated abbreviated TLTypes from 1.0 removed in this version are kept to not break compatibility backwards. EventControl added as per ticket #1305.a to not break compatibility backwards  |
|  1.1.1 | 2017-02-08 | Mattias Johannesson, SICK | - Fixing a few category typos and updating the Macros so that reference TXT/XML generation is working. - Untracked changes : typo fixes and missed references.  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  1.2.0 | 2020-06-17 | Mattias Johannesson, SICK | - GenTL 1.6 Partdatatype list extension, Region/ComponentID not tentative - Ticket #1305: Event control clarifications – made via more referencing to other standards documents and additional text. - Ticket #1942: Packet size renegotiation - Ticket #1985: Feature Persistence, added some text. - Added “Newest Only” buffer mode picture. - Fixing misaligned tables 2-19 and forward from 1.1.1. - Stream Information category Mandatory since it has mandatory features. - Added Optional “TLType” row for features, column for feature summary, allowing features to be Mandatory for specified TLs.  |
| --- | --- | --- | --- |

# 1 Introduction

The GenICam standards (see http://www.emva.org/standards-technology/genicam) define a generic standard software interface for industrial cameras. The GenICam standards are hosted by the EMVA. Part of the GenICam standards is GenTL, a generic Transport Layer interface on the host system, e.g. a PC. This document defines the Standard Features Naming Convention (SFNC) for the GenTL interface.

The GenICam GenTL standard provides a generic way to enumerate devices known to a system, communicate with one or more devices and, if possible, stream data from the device to the host independent from the underlying transport technology. This allows a third party software to use different technologies to control cameras and to acquire data in a transport layer agnostic way.

Besides the definition of a C interface with a defined behavior, the GenICam GenTL standard also defines a set of feature names and their meanings. To access these features the GenICam GenApi module is used.

The goal of the GenICam GenTL “Standard Features Naming Convention (GenTL SFNC)” is to standardize the features used in different GenTL Producer implementations. Thus, the GenICam GenTL standard should be decoupled as far as possible from the definition of specific feature names and their meaning. Note that the GenTL SFNC does not substitute or hide the features defined in the regular GenICam SFNC that defines the features for remote devices, but complement it by covering explicitly only the features of the GenTL Producer itself.

The GenTL Standard Features Naming Convention of GenICam is targeting maximum usability by existing and future transport layer technologies. It provides the definitions of a standard behavioral model and of standard features. The goal is to cover and to standardize the naming convention used in all the basic use cases where the implementation by different vendors would be very similar anyway.

## 1.1 GenICam Reference documents

It is recommended to study the GenICam Standard, the device-oriented GenICam Standard Feature Naming Convention (SFNC) and the GenICam GenTL Standard to understand this document.

The revisions relevant for this release are

|  Standard | Version | Date  |
| --- | --- | --- |
|  GenICam Standard | 2.1.1 | 2016-01-18  |
|  GenICam GenTL Standard | 1.6 | 2019-11-04  |
|  GenICam Standard Feature Naming Convention (SFNC) | 2.5 | 2019-05-07  |

## 1.2 TL specific features

All Transport Layer Specific features have a prefix. Currently only GigE Vision features are in this document and they all have the prefix “Gev”. The GigE Vision standard is hosted by AIA.

|  GEN<ICAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 1.3 Standard Definitions

This section defines the terms used in this document. See Transport layer specific standards, as well as GenICam and SFNC for detailed information.

|  Entity | An Entity is an end point located at either side (Host or Device) of a Communication.  |
| --- | --- |
|  Host System | The Host System is the Entity that takes control over a Device. A Host System can be the sink or the source for the data being streamed. Under GenICam the Host System must read and use the GenICam compliant XML file of the Device to control it.  |
|  Device | The Device is an Entity that is controlled by a Host System. A Device can be the source or the sink for streaming data. It can be remote (outside the Host System) or local (in the Host System). Under GenICam the Device must provide a GenICam compliant XML file and a register-based control access.  |
|  Link | A Link is the virtual binding between a Host System and a Device to establish a Communication. A Link is logical and may use one or more physical Connections.  |
|  Connection | A Connection is the physical binding between a Host System and a Device.  |
|  Interface | A: A virtual endpoint of the Link between a Device and a Host System. B: A GenICam programming interface class, e.g. Uint or Command.  |
|  Consumer | A library or application using an implementation of a GenTL Transport Layer Interface.  |
|  Producer | GenTL Transport Layer Interface implementation.  |
|  Adapter | A physical entity located in the Host System that has one or many Interfaces.  |
|  Communication | A Communication is an exchange of information between two Entities using a Link.  |
|  Channel | A logical point-to-point Communication over a Link. There may be multiple Channels on a single Link.  |
|  Transport Layer | The layer of Communication responsible to transport information between Entities.  |
|  Transmitter | An Entity that acts as the source for streaming data. This may apply to a Host System or a Device.  |
|  Receiver | An Entity that acts as the sink for streaming data. This may apply to a Host System or a Device.  |
|  Transceiver | An Entity that can receive and transmit streaming data. This may apply to a Host System or a Device.  |
|  Peripheral | An Entity that neither acts as a source nor as a sink for streaming data but can be controlled.  |
|  Stream | A flow of data that comes from a source and goes to a sink. A data Stream can be composed of images or chunk of data.  |
|  Stream Channel | A Communication Channel used to transmit a data Stream from a Transmitter (or Transceiver) to a Receiver (or Transceiver).  |
|  Event Channel | A Communication Channel used by the Device to notify the Host System asynchronously of Events. The Host System could also use an Event Channel to  |

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

|  GEN<ICAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 1.4 Conventions

# **Feature Name and Interface**

According to the GenICam standard, all the public features of a GenTL Producer must be included in the corresponding GenICam XML files following the GenTL module hierarchy and must use the SFNC Name and Interface type for those features if they exist. Other vendor specific or specialized features not mapping to existing SNFC features can be included but must be located in a vendor specific namespace in the GenICam XML and may use a vendor specific name.

This document lists for each feature the Name and Interface type that must be used.

# **Feature Category**

With the GenICam standard, each feature should be included in a "Category". The Category element defines in which group of features the feature will be located.

The Category does not affect the functionality of the features but is used by the GUIs to group the features when displaying them. The purpose is mainly to insure that the GUI can present features in a more organized way.

This document lists for each feature, a recommended Category that should be used.

# **Feature Level**

In this document, features are tagged according to the following requirement levels:

- **M: Mandatory** - Must be implemented to achieve compliance with the GenICam GenTL standard.
- **R: Recommended** - This feature adds important aspects to the use case and must respect the naming convention if used.
- **O: Optional** - This feature is less critical. Nevertheless, it is considered and must respect the naming convention if used.

For additional details about the mandatory general features, please refer to the GenICam GenTL standard. For additional details about the mandatory features to certain transport layers, please refer to the text of those standards.

# **Feature Visibility**

According to the GenICam standard, each feature can be assigned a "Visibility". The Visibility defines the type of user that should get access to the feature. Possible values are Beginner, Expert, Guru and Invisible. The latter is required to make features accessible from the API, but invisible in the GUI.

The visibility does not affect the functionality of the features but is used by the GUI to decide which features to display based on the current user level. The purpose is mainly to insure that the GUI is not cluttered with information that is not intended at the current user level.

The following criteria have been used for the assignment of the recommended visibility:

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

- **B: Beginner** – Features that should be visible for *all* users via the GUI and API. This is the default visibility in the GenICam XML files and will be used if the Visibility element is omitted for a feature. The number of features with "Beginner" visibility should be limited to all **basic** features of the GenTL Producer so the GUI display is well-organized and is easy to use.
- **E: Expert** – Features that require a more in-depth knowledge of the device functionality. This is the preferred visibility level for all advanced features in the devices.
- **G: Guru** – Advanced features that might bring the devices into a state where it will not work properly anymore if it is set incorrectly for the devices current mode of operation.
- **I: Invisible** – Features that should be kept hidden for the GUI users but still be available via the API.

This document lists for each feature, a recommended Visibility that should be used.

### Selector

A selector is used to index which instance of the feature is accessed in situations where multiple instances of a feature exist.

A selector is a separate feature that is typically an IEnumeration or an IInteger. Selectors must be used only to select the target features for subsequent changes. It is not allowed to change the behavior of a GenTL Producer in response to a change of a selector value.

If a selector has only one possible value, the selector relation can be omitted but it is recommended to leave the selector feature as read only for information purpose.

In this document, the features which potentially dependent on a selector are expressed using the C language convention for arrays: a pair of brackets follows the feature name, like in SelectedFeature[Selector]. When the Selector is not present, one must deduce the feature is not an array.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

## 1.5 Standard Units

The following abbreviations are used as standard units for features described in this document. Note that all units are using plain ASCII characters.

|  ns | nanoseconds  |
| --- | --- |
|  us | microseconds  |
|  ms | milliseconds  |
|  s | seconds  |
|  B | Bytes  |
|  Bps | Bytes per second  |
|  MBps | Mega Bytes per second  |
|  Mbps | Mega bits per second  |
|  Fps | Frames per second  |
|  dB | Decibels  |
|  C | Celsius  |
|  Hz | Hertz  |
|  % | Percent  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 1.6 Acronyms

|  AIA | Automated Imaging Association. See http://www.visiononline.org  |
| --- | --- |
|  DHCP | Dynamic Host Configuration Protocol  |
|  EMVA | European Machine Vision Association. See http://www.emva.org  |
|  ID | Identifier  |
|  IP | Internet Protocol  |
|  LLA | Link-Local Address  |
|  MAC | Media Access Control  |
|  R | Read (or Recommended, depends on the context)  |
|  R/W | Read and Write, if one of the letters is in brackets either read or write is optional, for example R(/W) means read and optionally write  |
|  W | Write  |
|  XML | eXtensible Markup Language  |

## 2 Features Summary

This section provides a comprehensive list of the standard features covered by this document. The following sections provide more detailed explanation of each feature.

### 2.1 System Module

#### 2.1.1 System Information

Contains the features related to general information about the GenTL Producer.

Table 2-1: System Information Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  SystemInformation | R | All | ICategory | R | - | B | Category that contains all System Information features of the System module.  |
|  TLID | M | All | IString | R | - | E | Unique identifier of the GenTL Producer like a GUID.  |
|  TLVendorName | M | All | IString | R | - | B | Name of the GenTL Producer vendor.  |
|  TLModelName | M | All | IString | R | - | B | Name of the GenTL Producer to distinguish different kinds of GenTL Producer implementations from one vendor.  |
|  TLVersion | M | All | IString | R | - | B | Vendor specific version string of the GenTL Producer.  |
|  TLFileName | R | All | IString | R | - | E | Filename including extension of the GenTL Producer.  |
|  TLDisplayName | R | All | IString | R/(W) | - | B | User readable name of the GenTL Producer.  |
|  TLPath | M | All | IString | R | - | E | Full path to the GenTL Producer including filename and extension.  |
|  TLType | M | All | IEnumeration | R | - | E | Transport layer type of the GenTL Producer implementation.  |

![img-4.jpeg](img-4.jpeg)

|  GenTLVersionMajor | M | All | IInteger | R | - | E | Major version number of the GenTL specification the GenTL Producer implementation complies with.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  GenTLVersionMinor | M | All | IInteger | R | - | E | Minor version number of the GenTL specification the GenTL Producer implementation complies with.  |
|  GenTLSFNCVersionMajor | R | All | IInteger | R | - | E | Major version number of the GenTL Standard Features Naming Convention that was used to create the GenTL Producer's XML.  |
|  GenTLSFNCVersionMinor | R | All | IInteger | R | - | E | Minor version number of the GenTL Standard Features Naming Convention that was used to create the GenTL Producer's XML.  |
|  GevVersionMajor | O | GEV | IInteger | R | - | E | This feature is deprecated (See InterfaceTLVersionMajor).  |
|  GevVersionMinor | O | GEV | IInteger | R | - | E | This feature is deprecated (See InterfaceTLVersionMinor).  |

#### 2.1.2 Interface Enumeration

Contains the features related to the enumeration of available Interface modules within the System module of a GenTL Producer.

Table 2-2: Interface Enumeration Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  InterfaceEnumeration | R | All | ICategory | R | - | B | Category that contains all Interface Enumeration features of the System module.  |
|  InterfaceUpdateList | M | All | ICommand | (R)/W | - | B | Updates the internal list of the interfaces.  |
|  InterfaceUpdateTimeout | R | All | IInteger | R/W | ms | E | Specifies timeout for the InterfaceUpdateList Command.  |
|  InterfaceSelector | M | All | IInteger | R/W | - | B | Selector for the different GenTL Producer interfaces.  |
|  InterfaceID[InterfaceSelector] | M | All | IString | R | - | B | GenTL Producer wide unique identifier of the selected interface.  |

![img-5.jpeg](img-5.jpeg)

|  InterfaceDisplayName[InterfaceSelector] | R | All | IString | R | - | B | A user-friendly name of the Interface.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  GevInterfaceMACAddress[InterfaceSelector] | M | GEV | IInteger | R | - | E | 48-bit MAC address of the selected interface.  |
|  GevInterfaceDefaultIPAddress[InterfaceSelector] | M | GEV | IInteger | R | - | E | IP address of the first subnet of the selected interface.  |
|  GevInterfaceDefaultSubnetMask[InterfaceSelector] | M | GEV | IInteger | R | - | E | Subnet mask of the first subnet of the selected interface.  |
|  GevInterfaceDefaultGateway[InterfaceSelector] | R | GEV | IInteger | R | - | E | Gateway of the selected interface.  |

#### 2.1.3 GenICam Control

Contains the features related to GenICam control and access of the System module.

Table 2-3: GenICam Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  Root | M | All | ICategory | R | - | B | Provides the Root of the GenICam features tree.  |
|  TLPort | M | All | IPort | R/W | - | I | The GenICam port through which the System module is accessed.  |

## 2.1.4 Event Control

Category that contains Event Control features.

Table 2-4: Event Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  EventControl | R | All | ICategory | R | - | E | Category that contains Event control features.  |
|  EventSelector | R | All | IEnumeration | R/W | - | E | Selects which Event to signal to the host application.  |
|  EventNotification[EventSelector] | R | All | IEnumeration | R/W | - | E | Activate or deactivate the notification to the host application of the occurrence of the selected Event.  |

### 2.2 Interface Module

#### 2.2.1 Interface Information

Contains the features related to general information about a specific Interface module.

Table 2-5: Interface Information Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  InterfaceInformation | R | All | ICategory | R | - | E | Category that contains all Interface Information features of the Interface module.  |
|  InterfaceID | M | All | IString | R | - | E | GenTL Producer wide unique identifier of the selected interface.  |
|  InterfaceDisplayName | R | All | IString | R | - | E | A user-friendly name of the Interface.  |
|  InterfaceType | M | All | IEnumeration | R | - | E | Transport layer type of the interface.  |
|  InterfaceTLVersionMajor | M | All | IInteger | R | - | E | Major version number of the transport layer specification the GenTL Producer interface complies with.  |
|  InterfaceTLVersionMinor | M | All | IInteger | R | - | E | Minor version number of the transport layer specification the GenTL Producer interface complies with.  |
|  GevInterfaceGatewaySelector | M | GEV | IInteger | R/W | - | E | Selector for the different gateway entries for this interface.  |
|  GevInterfaceGateway[GevInterfaceGatewaySelector] | M | GEV | IInteger | R | - | E | IP address of the selected gateway entry of this interface.  |
|  GevInterfaceMACAddress | M | GEV | IInteger | R | - | E | 48-bit MAC address of this interface.  |
|  GevInterfaceSubnetSelector | M | GEV | IInteger | R/W | - | E | Selector for the subnet of this interface.  |
|  GevInterfaceSubnetIPAddress[GevInterfaceSubnetSelector] | M | GEV | IInteger | R | - | E | IP address of the selected subnet of this interface.  |
|  GevInterfaceSubnetMask[GevInterfaceS | M | GEV | IInteger | R | - | E | Subnet mask of the selected subnet of this interface.  |

![img-6.jpeg](img-6.jpeg)

|  ubnetSelector] |  |  |  |  |  |  |   |
| --- | --- | --- | --- | --- | --- | --- | --- |

### 2.2.2 Device Enumeration

Contains the features related to the enumeration of available Device modules within a specific Interface module.

Table 2-6: Device Enumeration Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  DeviceEnumeration | R | All | ICategory | R | - | E | Category that contains all Device Enumeration features of the Interface module.  |
|  DeviceUpdateList | M | All | ICommand | (R)/W | - | E | Updates the internal device list.  |
|  DeviceUpdateTimeout | R | All | IInteger | R/W | ms | E | Specifies timeout for the DeviceUpdateList Command.  |
|  DeviceSelector | M | All | IInteger | R/W | - | E | Selector for the different devices on this interface.  |
|  DeviceID[DeviceSelector] | M | All | IString | R | - | E | Interface wide unique identifier of the selected device.  |
|  DeviceVendorName[DeviceSelector] | M | All | IString | R | - | E | Name of the device vendor.  |
|  DeviceModelName[DeviceSelector] | M | All | IString | R | - | E | Name of the device model.  |
|  DeviceAccessStatus[DeviceSelector] | M | All | IEnumeration | R | - | E | Gives the device's access status at the moment of the last execution of the DeviceUpdateList command.  |
|  DeviceSerialNumber[DeviceSelector] | R | All | IString | R | - | E | Serial number of the remote device.  |
|  DeviceUserID[DeviceSelector] | O | All | IString | R | - | E | User-programmable device identifier of the remote device.  |
|  DeviceTLVersionMajor[DeviceSelector] | M | All | IInteger | R | - | E | Major version number of the transport layer specification the remote device complies with.  |
|  DeviceTLVersionMinor[DeviceSelector] | M | All | IInteger | R | - | E | Minor version number of the transport layer specification the remote device complies with.  |
|  GevDeviceIPAddress[DeviceSelector] | M | GEV | IInteger | R | - | E | Current IP address of the GVCP interface of the selected  |

|   |  |  |  |  |  |  | remote device.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  GevDeviceSubnetMask[DeviceSelector] | M | GEV | IInteger | R | - | E | Current subnet mask of the GVCP interface of the selected remote device.  |
|  GevDeviceGateway[DeviceSelector] | R | GEV | IInteger | R | - | E | Current gateway IP address of the GVCP interface of the selected remote device.  |
|  GevDeviceIPConfigurationStatus[DeviceSelector] | R | GEV | IEnum | R/W | - | E | Device IP configuration of the GVCP interface of the selected remote device.  |
|  GevDeviceMACAddress[DeviceSelector] | M | GEV | IInteger | R | - | E | 48-bit MAC address of the GVCP interface of the selected remote device.  |
|  GevDeviceCurrentControlMode[DeviceSelector] | O | GEV | IEnum | R/W | - | E | The current control mode of the device.  |
|  GevApplicationSwitchoverKey[DeviceSelector] | O | GEV | IInteger | W | - | E | Application switchover key to use when requesting ControlAccess switchover.  |
|  GevDeviceForceIP[DeviceSelector] | R | GEV | ICommand | (R)/W | - | E | Apply the force IP settings (GevDeviceForceIPAddress, GevDeviceForceSubnetMask and GevDeviceForceGateway) in the Device using ForceIP command.  |
|  GevDeviceForceIPAddress[DeviceSelector] | R | GEV | IInteger | R/W | - | E | Static IP address to set for the GVCP interface of the remote device.  |
|  GevDeviceForceSubnetMask[DeviceSelector] | R | GEV | IInteger | R/W | - | E | Static subnet mask to set for GVCP interface of the remote device.  |
|  GevDeviceForceGateway[DeviceSelector] | R | GEV | IInteger | R/W | - | E | Static gateway IP address to set for the GVCP interface of the remote device.  |

### 2.2.3 Action Control

Category that contains Action Control features.

Table 2-7: Action Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  ActionControl | R | GEV | ICategory | R | - | E | Category that contains all Action Control features of the Interface module.  |
|  ActionCommand | R | GEV | ICommand | (R)/W | - | E | Send ActionCommand to device(s).  |
|  ActionDeviceKey | R | GEV | IInteger | R/W | - | E | The Action Command Device Key to use in the Action Command.  |
|  ActionGroupKey | R | GEV | IInteger | R/W | - | E | The Action Command Group Key to use in the Action Command.  |
|  ActionGroupMask | R | GEV | IInteger | R/W | - | E | The Action Command Group Mask to use in the Action Command.  |
|  ActionScheduledTimeEnable | R | GEV | IBoolean | R/W | - | E | Specifies if a time enabled Action Command should be given.  |
|  ActionScheduledTime | R | GEV | IInteger | R/W | - | E | Specifies the time a time enabled Action Command is scheduled.  |
|  GevActionDestinationIPAddress | R | GEV | IInteger | R/W | - | E | Specifies destination the IP address for the Action Command.  |

### 2.2.4 GenICam Control

Contains the features related to GenICam control and access of a specific Interface module.

Table 2-8: GenICam Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  Root | M | All | ICategory | R | - | B | Provides the Root of the GenICam features tree.  |
|  InterfacePort | M | All | IPort | R/W | - | I | The GenICam port through which the Interface module is accessed.  |

#### 2.2.5 Event Control

Category that contains Event Control features.

Table 2-9: Event Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  EventControl | R | All | ICategory | R | - | E | Category that contains Event control features.  |
|  EventSelector | R | All | IEnumeration | R/W | - | E | Selects which Event to signal to the host application.  |
|  EventNotification[EventSelector] | R | All | IEnumeration | R/W | - | E | Activate or deactivate the notification to the host application of the occurrence of the selected Event.  |

### 2.3 Device Module

#### 2.3.1 Device Information

Contains the features related to general information about a specific Device module.

Table 2-10: Device Information Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |

|  DeviceInformation | R | All | ICategory | R | - | B | Category that contains all Device Information features of the Device module.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  DeviceID | M | All | IString | R | - | E | Interface-wide unique identifier of this device.  |
|  DeviceSerialNumber | R | All | IString | R | - | E | Serial number of the remote device.  |
|  DeviceUserID | O | All | IString | R/W | - | E | User-programmable device identifier of the remote device.  |
|  DeviceVendorName | M | All | IString | R | - | B | Name of the remote device vendor.  |
|  DeviceModelName | M | All | IString | R | - | B | Name of the remote device model.  |
|  DeviceFamilyName | R | All | IString | R | - | B | Name of the product family of the remote device model.  |
|  DeviceVersion | R | All | IString | R | - | B | The version of the remote device model.  |
|  DeviceManufacturerInfo | R | All | IString | R | - | B | Manufacturer information about the remote device.  |
|  DeviceType | M | All | IEnumeration | R | - | E | Transport layer type of the device.  |
|  DeviceDisplayName | R | All | IString | R | - | E | User readable name of the device.  |
|  DeviceTimestampFrequency | R | All | IInteger | R | - | B | The tick-frequency of the time stamp clock.  |
|  DeviceAccessStatus | M | All | IEnumeration | R | - | E | Gives the device's access status at the moment of the last execution of the DeviceUpdateList command.  |
|  DeviceChunkDataFormat | R | All | IEnumeration | R | - | E | Chunk data format used by the device.  |
|  DeviceEventDataFormat | R | All | IEnumeration | R | - | E | Enumeration, informing about the event data format used by the device (meaning the "device events", see event type EVENT_REMOTE_DEVICE (named EVENT_FEATURE_DEVEVENT in GenTL up to version 1.  |
|  GevDeviceMACAddress | M | GEV | IInteger | R | - | E | 48-bit MAC address of the GVCP interface of the remote device.  |
|  GevDeviceIPAddress | M | GEV | IInteger | R | - | E | Current IP address of the GVCP interface of the remote device.  |

## GEN<|>CAM

Version 1.2.0

GenTL Standard Features Naming Convention

![img-7.jpeg](img-7.jpeg)

emva

|  GevDeviceSubnetMask | M | GEV | Integer | R | - | E | Current subnet mask of the GVCP interface of the remote device.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  GevDeviceGateway | M | GEV | Integer | R | - | E | Current gateway IP address of the GVCP interface of the remote device.  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

#### 2.3.2 Device Control

Contains the features related to configure a specific Device module.

Table 2-11: Device Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  DeviceControl | R | All | ICategory | R | - | E | Category that contains all Device Control features of the Device module.  |
|  DeviceEndianessMechanism | M | GEV | IEnumeration | R/W | - | E | Identifies the endianess handling mode.  |
|  LinkCommandTimeout | R | All | IFloat | R/W | us | G | Specifies application timeout for the control channel communication.  |
|  LinkCommandRetryCount | R | All | IInteger | R/W | - | G | Specifies maximum number of tries before failing the control channel commands.  |

#### 2.3.3 Stream Enumeration

Contains the features related to the enumeration of available Data Stream modules within a specific Device module.

Table 2-12: Stream Enumeration Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  StreamEnumeration | R | All | ICategory | R | - | B | Category that contains all Stream Enumeration features of the Device module.  |
|  StreamSelector | M | All | IInteger | R/W | - | B | Selector for the different stream channels.  |
|  StreamID[StreamSelector] | M | All | IString | R | - | B | Device unique ID for the stream.  |

|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

#### 2.3.4 GenICam Control

Contains the features related to GenICam control and access of a specific Device module.

Table 2-13: GenICam Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  Root | M | All | ICategory | R | - | B | Provides the Root of the GenICam features tree.  |
|  DevicePort | M | All | IPort | R/W | - | I | The GenICam port through which the Device module is accessed.  |

#### 2.3.5 Event Control

Category that contains Event Control features.

Table 2-14: Event Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  EventControl | R | All | ICategory | R | - | E | Category that contains Event control features.  |
|  EventSelector | R | All | IEnumeration | R/W | - | E | Selects which Event to signal to the host application.  |
|  EventNotification[EventSelector] | R | All | IEnumeration | R/W | - | E | Activate or deactivate the notification to the host application of the occurrence of the selected Event.  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

## 2.4 Data Stream Module

### 2.4.1 Stream Information

Contains the features related to general information about a specific Data Stream module.

Table 2-15: Stream Information Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  StreamInformation | M | All | ICategory | R | - | E | Category that contains all Stream Information features of the Data Stream module.  |
|  StreamID | M | All | IString | R | - | E | Device unique ID for the data stream.  |
|  StreamType | M | All | IEnumeration | R | - | E | Transport layer type of the Data Stream.  |

### 2.4.2 Device Stream Channel Control

Contains the features related to control the buffers within the acquisition engine of a specific Data Stream module.

Table 2-16: Buffer Handling Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  DeviceStreamChannelControl | R | GEV | ICategory | R | - | E | Category containing features to control the stream channel shared between the remote device and the GenTL Producer's data stream module.  |
|  DeviceStreamChannelPacketSize | R | GEV | IInteger | R/(W) | B | E | Specifies the stream packet size, in bytes, to send on the selected channel for a transmitter or specifies the maximum packet size supported by a receiver.  |
|  DeviceStreamChannelPacketSizeMin | O | GEV | IInteger | R/(W) | B | G | Controls desired minimum of the packet size feature to be  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|   |  |  |  |  |  |  | used for the stream channel.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  DeviceStreamChannelPacketSizeMax | O | GEV | IInteger | R/(W) | B | G | Controls desired maximum of the packet size feature to be used for the stream channel.  |
|  DeviceStreamChannelPacketSizeInc | O | GEV | IInteger | R/(W) | B | G | Controls desired increment of the packet size feature to be used for the stream channel.  |
|  DeviceStreamChannelNegotiatePacketSize | O | GEV | ICommand | (R)/W | - | E | Starts negotiation for the optimal packet size considering the remote device, host and their connection path.  |

### 2.4.3 Buffer Handling Control

Contains the features related to GenICam control and access of a specific Data Stream module.

Table 2-17: GenICam Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  BufferHandlingControl | R | All | ICategory | R | - | B | Contains all features of the Data Stream module that control the used buffers.  |
|  StreamAnnouncedBufferCount | M | All | IInteger | R | - | E | Number of announced (known) buffers on this stream.  |
|  StreamBufferHandlingMode | M | All | IEnumeration | R/(W) | - | B | Available buffer handling modes of this Data Stream.  |
|  StreamAnnounceBufferMinimum | M | All | IInteger | R | - | E | Minimal number of buffers to announce to enable selected buffer handling mode.  |
|  StreamDeliveredFrameCount | R | All | IInteger | R | - | E | Number of delivered frames since last acquisition start.  |
|  StreamLostFrameCount | R | All | IInteger | R | - | E | Number of lost frames due to queue underrun.  |
|  StreamInputBufferCount | O | All | IInteger | R | - | E | Number of buffers in the input buffer pool plus the buffers(s) currently being filled.  |
|  StreamOutputBufferCount | R | All | IInteger | R | - | E | Number of buffers in the output buffer queue.  |

|  StreamStartedFrameCount | R | All | Integer | R | - | E | Number of frames started in the acquisition engine.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  PayloadSize | R | All | Integer | R | Byte | E | Size of the expected data in bytes.  |
|  StreamIsGrabbing | R | All | Boolean | R |  | E | Flag indicating whether the acquisition engine is started or not.  |
|  StreamChunkCountMaximum | R | All | Integer | R |  | E | Maximum number of chunks to be expected in a buffer (can be used to allocate the array for the DSGetBufferChunkData function).  |
|  StreamBufferAlignment | R | All | Integer | R | Byte | E | Alignment size in bytes of the buffers passed to DSAnnounceBuffer.  |

#### 2.4.4 GenICam Control

Category that contains Event Control features.

Table 2-18: Event Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  Root | M | All | ICategory | R | - | B | Provides the Root of the GenICam features tree.  |
|  StreamPort | M | All | IPort | R/W | - | I | The GenICam port through which the Data Stream module is accessed.  |

#### 2.4.5 Event Control

Contains the features related to the Event Buffers Discarded.

Table 2-19: Buffer Discarded Event Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  EventControl | R | All | ICategory | R | - | E | Category that contains Event control features.  |
|  EventSelector | R | All | IEnumeration | R/W | - | E | Selects which Event to signal to the host application.  |
|  EventNotification[EventSelector] | R | All | IEnumeration | R/W | - | E | Activate or deactivate the notification to the host application of the occurrence of the selected Event.  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 2.5 Buffer Module

#### 2.5.1 Buffer Information

Contains the features related to general information about a specific Buffer module.

Table 2-20: Buffer Information Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  BufferInformation | O | All | ICategory | R | - | E | Category that contains all Buffer Information features of the Buffer module.  |
|  BufferUserData | O | All | IInteger | R | - | E | Pointer to user data casted to an integer number referencing GenTL Consumer specific data.  |
|  BufferType | O | All | IEnumeration | R | - | E | Transport layer type of the buffer.  |
|  BufferSize | O | All | IInteger | R | Byte | E | Size of the buffer in bytes.  |

#### 2.5.2 Buffer Data Information

Contains the features related to the currently filled data of a specific Buffer module.

Table 2-21: Buffer Data Information Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  BufferDataInformation | O | All | ICategory | R | - | E | Contains all Buffer Data Information features of the Buffer module.  |

|  BufferData | O | All | IRegister | R/(W) | - | E | Entire buffer data.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  BufferTimeStamp | O | All | IInteger | R | - | E | Timestamp the buffer was acquired.  |
|  BufferNewData | O | All | IBoolean | R | - | E | Flag to indicate that the buffer contains new data since the last delivery.  |
|  BufferIsQueued | O | All | IBoolean | R | - | E | Flag to indicate if the buffer is in the input pool or output buffer queue.  |
|  BufferIsAcquiring | O | All | IBoolean | R | - | E | Flag to indicate that the buffer is currently being filled with data.  |
|  BufferIsIncomplete | O | All | IBoolean | R | - | E | Flag to indicate that a buffer was filled but an error occurred during that process.  |
|  BufferPayloadType | O | All | IEnumeration | R | - | E | Payload type of the data.  |
|  BufferNumberOfParts | O | All | IInteger | R | - | E | The number of parts in the current buffer as delivered by the transport mechanism.  |
|  BufferPartSelector | O | All | IInteger | R | - | E | The buffer part to extract information from.  |
|  BufferSizeFilled | O | All | IInteger | R | Byte | E | Number of bytes written into the buffer last time it was filled.  |
|  BufferPartDataType[BufferPartSelector] | O | All | IEnumeration | R | - | E | The data type of the part.  |
|  BufferPartSourceIDValue[BufferPartSelector] | O | All | IInteger | R | - | E | The Source ID type of the part.  |
|  BufferPartRegionIDValue[BufferPartSelector] | O | All | IInteger | R | - | E | The Region ID type of the part.  |
|  BufferPartComponentIDValue[BufferPartSelector] | O | All | IInteger | R | - | E | The Component ID type of the part.  |
|  BufferWidth[BufferPartSelector] | O | All | IInteger | R | - | E | Width of the data in the buffer in number of pixels.  |
|  BufferHeight[BufferPartSelector] | O | All | IInteger | R | - | E | Height of the data in the buffer in number of pixels as configured.  |

![img-8.jpeg](img-8.jpeg)

|  BufferXOffset[BufferPartSelector] | O | All | IInteger | R | - | E | XOffset of the data in the buffer in number of pixels from the image origin to handle areas of interest.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  BufferYOffset[BufferPartSelector] | O | All | IInteger | R | - | E | YOffset of the data in the buffer in number of lines from the image origin to handle areas of interest.  |
|  BufferXPadding[BufferPartSelector] | O | All | IInteger | R | Byte | E | XPadding of the data in the buffer in number of bytes.  |
|  BufferYPadding | O | All | IInteger | R | Byte | E | YPadding of the data in the buffer in number of bytes.  |
|  BufferFrameID | R | All | IInteger | R | - | E | A sequentially incremented number of the frame.  |
|  BufferImagePresent | O | All | IBoolean | R | - | E | Flag to indicate if the current data in the buffer contains image data.  |
|  BufferImageOffset | O | All | IInteger | R | Byte | E | Offset of the image data from the beginning of the delivered buffer in bytes.  |
|  BufferPixelFormat[BufferPartSelector] | O | All | IEnumeration | R | - | E | Format of the pixels provided by the buffer.  |
|  BufferDeliveredImageHeight[BufferPart Selector] | O | All | IInteger | R | - | E | The number of lines in the current buffer part as delivered by the transport mechanism.  |
|  BufferDeliveredChunkPayloadSize | O | All | IInteger | R | - | E | Size of the valid chunk payload data delivered in the buffer.  |
|  BufferChunkLayoutID | O | All | IInteger | R | - | E | ID of the chunk data layout delivered in the buffer.  |
|  BufferFileName | O | All | IString | R | - | E | Filename for the file payload data delivered in the buffer.  |

### 2.5.3 GenICam Control

Contains the features related to GenICam control and access of a specific Buffer module.

Table 2-22: GenICam Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  Root | O | All | ICategory | R | - | B | Provides the Root of the GenICam features tree.  |

|  **GEN<I>CAM** |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  BufferPort | O | All | IPort | R/W | - | I | The GenICam port through which the Buffer module is accessed.  |
| --- | --- | --- | --- | --- | --- | --- | --- |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3 General Features

Contains all features that are independent from the underlying transport technology, in particular including all mandatory features for all GenTL Producer implementations.

### 3.1 System Module

Contains all features of the System module that are independent from the underlying transport technology.

#### 3.1.1 System Information

Features in this section provide basic information about the System Module and its identity. Note that all features in this section are defined as read-only.

3.1.1.1 SystemInformation

|  Name | SystemInformation  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category that contains all System Information features of the System module.

3.1.1.2 TLID

|  Name | TLID  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Mandatory  |
|  Interface | ISString  |
|  Access | Read  |
|  Unit | -  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Visibility | Expert  |
| --- | --- |
|  Values | Any NULL-terminated string  |

Unique identifier of the GenTL Producer like a GUID.

Corresponds to the TL_INFO_ID command of TLGetInfo function.

##### 3.1.1.3 TLVendorName

|  Name | TLVendorName  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Name of the GenTL Producer vendor.

Corresponds to the TL_INFO_VENDOR command of TLGetInfo function.

##### 3.1.1.4 TLModelName

|  Name | TLModelName  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Name of the GenTL Producer to distinguish different kinds of GenTL Producer implementations from one vendor.

|  GEN<i>CAM |   | ![img-9.jpeg](img-9.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Corresponds to the TL_INFO_MODEL command of TLGetInfo function.

3.1.1.5 TLVersion

|  Name | TLVersion  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Vendor specific version string of the GenTL Producer.

Corresponds to the TL_INFO_VERSION command of TLGetInfo function.

3.1.1.6 TLFileName

|  Name | TLFileName  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Recommended  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

Filename including extension of the GenTL Producer.

Corresponds to the TL_INFO_NAME command of TLGetInfo function.

3.1.1.7 TLDisplayName

|  Name | TLDisplayName  |
| --- | --- |
|  Category | SystemInformation  |

|  GEN<i>CAM |   | ![img-10.jpeg](img-10.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Level | Recommended  |
| --- | --- |
|  Interface | IString  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

User readable name of the GenTL Producer.

Corresponds to the TL_INFO_DISPLAYNAME command of TLGetInfo function.

3.1.1.8 TLPath

|  Name | TLPath  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

Full path to the GenTL Producer including filename and extension.

Corresponds to the TL_INFO_PATHNAME command of TLGetInfo function.

3.1.1.9 TLType

|  Name | TLType  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Mandatory  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Values | GigEVision CameraLink CameraLinkHS CoaXPress USB3Vision Mixed Custom CL (Deprecated) CLHS (Deprecated) CXP (Deprecated) Ethernet (Deprecated) IIDC (Deprecated) PCI (Deprecated) USB3 (Deprecated) UVC (Deprecated)  |
| --- | --- |

Transport layer type of the GenTL Producer implementation.

Note that these values already follow the updated value list of the "DeviceTLType" feature from GenICam SFNC 2.3. Depending on this value, the transport layer specific features for the chosen transport layer standard have to be considered.

- CameraLink: Camera Link
- CameraLinkHS: Camera Link High Speed
- CoaXPress: CoaXPress
- GigEVision: GigE Vision
- USB3Vision: USB3 Vision
- Mixed: Different Interface modules of the GenTL Producer are of different types
- Custom: Custom transport layer
- CL (Deprecated): Camera Link
- CLHS (Deprecated): Camera Link HS
- CXP (Deprecated): CoaXPress
- Ethernet (Deprecated): Generic Ethernet
- GEV (Deprecated): GigE Vision
- IIDC (Deprecated): IIDC 1394

|  GEN<i>CAM |   | ![img-11.jpeg](img-11.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

• PCI (Deprecated): PCI / PCIe
• USB3 (Deprecated): USB3 Vision
• UVC (Deprecated): USB video class

Corresponds to the TL_INFO_TLTYPE command of TLGetInfo function.

3.1.1.10 GenTLVersionMajor

|  Name | GenTLVersionMajor  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | >0  |

Major version number of the GenTL specification the GenTL Producer implementation complies with.

3.1.1.11 GenTLVersionMinor

|  Name | GenTLVersionMinor  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Minor version number of the GenTL specification the GenTL Producer implementation complies with.

|  GEN<i>CAM |   | ![img-12.jpeg](img-12.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

3.1.1.12 GenTLSFNCVersionMajor

|  Name | GenTLSFNCVersionMajor  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | >0  |

Major version number of the GenTL Standard Features Naming Convention that was used to create the GenTL Producer's XML.

3.1.1.13 GenTLSFNCVersionMinor

|  Name | GenTLSFNCVersionMinor  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Minor version number of the GenTL Standard Features Naming Convention that was used to create the GenTL Producer's XML.

3.1.1.14 GevVersionMajor (Deprecated)

|  Name | GevVersionMajor  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Optional  |
|  TLType | GigEVision  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Interface | IInteger  |
| --- | --- |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | >0  |

This feature is deprecated (See InterfaceTLVersionMajor).

Major version number of the GigE Vision specification the GenTL Producer implementation complies with.

If the value of the feature TLType is "Mixed" but supports GigE Vision interfaces this feature must be present.

3.1.1.15 GevVersionMinor (Deprecated)

|  Name | GevVersionMinor  |
| --- | --- |
|  Category | SystemInformation  |
|  Level | Optional  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

This feature is deprecated (See InterfaceTLVersionMinor).

Minor version number of the GigE Vision specification the GenTL Producer implementation complies with.

If the value of the feature TLType is "Mixed" but supports GigE Vision interfaces this feature must be present.

#### 3.1.2 Interface Enumeration

The Interface Enumeration section describes all features related to discovery and enumeration of interfaces belonging to the System module.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.1.2.1 InterfaceEnumeration

|  Name | InterfaceEnumeration  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category that contains all Interface Enumeration features of the System module.

### 3.1.2.2 InterfaceUpdateList

|  Name | InterfaceUpdateList  |
| --- | --- |
|  Category | InterfaceEnumeration  |
|  Level | Mandatory  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Updates the internal list of the interfaces. This feature should be readable if the execution cannot performed immediately. The command then returns and the status can be polled. This function interacts with the TLUpdateInterfaceList function of the GenTL Producer. It is up to the GenTL Consumer to handle access in case both methods are used.

### 3.1.2.3 InterfaceUpdateTimeout

|  Name | InterfaceUpdateTimeout  |
| --- | --- |
|  Category | InterfaceEnumeration  |
|  Level | Recommended  |
|  Interface | IInteger  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Access | Read/Write  |
| --- | --- |
|  Unit | ms  |
|  Visibility | Expert  |
|  Values | >0  |

Specifies timeout for the InterfaceUpdateList Command.

### 3.1.2.4 InterfaceSelector

|  Name | InterfaceSelector  |
| --- | --- |
|  Category | InterfaceEnumeration  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥0  |

Selector for the different GenTL Producer interfaces. This interface list only changes on execution of "InterfaceUpdateList". The selector is 0-based in order to match the index of the C interface.

### 3.1.2.5 InterfaceID

|  Name | InterfaceID[InterfaceSelector]  |
| --- | --- |
|  Category | InterfaceEnumeration  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

GenTL Producer wide unique identifier of the selected interface.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.1.2.6 InterfaceDisplayName

|  Name | InterfaceDisplayName[InterfaceSelector]  |
| --- | --- |
|  Category | InterfaceEnumeration  |
|  Level | Recommended  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

A user-friendly name of the Interface.

Corresponds to the TLGetInterfaceID function with the index corresponding to "InterfaceSelector".

### 3.1.2.7 GevInterfaceMACAddress

|  Name | GevInterfaceMACAddress[InterfaceSelector]  |
| --- | --- |
|  Category | InterfaceEnumeration  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

48-bit MAC address of the selected interface. Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory, and that the Representation element should be used in the XML to facilitate understanding the data.

### 3.1.2.8 GevInterfaceDefaultIPAddress

|  Name | GevInterfaceDefaultIPAddress[InterfaceSelector]  |
| --- | --- |
|  Category | InterfaceEnumeration  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Level | Mandatory  |
| --- | --- |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

IP address of the first subnet of the selected interface. Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

#### 3.1.2.9 GevInterfaceDefaultSubnetMask

|  Name | GevInterfaceDefaultSubnetMask[InterfaceSelector]  |
| --- | --- |
|  Category | InterfaceEnumeration  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Subnet mask of the first subnet of the selected interface. Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

#### 3.1.2.10 GevInterfaceDefaultGateway

|  Name | GevInterfaceDefaultGateway[InterfaceSelector]  |
| --- | --- |
|  Category | InterfaceEnumeration  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | IInteger  |

|  GEN<i>CAM |   | ![img-13.jpeg](img-13.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Access | Read  |
| --- | --- |
|  Unit | -  |
|  Visibility | Export  |
|  Values |   |

Gateway of the selected interface.

#### 3.1.3 GenICam Control

This section provides the necessary features to use the GenICam feature tree of the System module.

Note: In case of discrepancy between the features described in this section and the “GenICam Standard text” the GenTL SFNC document prevails.

3.1.3.1 Root

|  Name | Root  |
| --- | --- |
|  Category | None  |
|  Level | Mandatory  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Provides the Root of the GenICam features tree.

3.1.3.2 TLPort

|  Name | TLPort  |
| --- | --- |
|  Category | None  |
|  Level | Mandatory  |
|  Interface | IPort  |
|  Access | Read/Write  |
|  Unit | -  |

|  GEN<I>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Visibility | Invisible  |
| --- | --- |
|  Values | -  |

The GenICam port through which the System module is accessed.

Note that TLPort is a port node (not a feature node) and is generally not accessed by the end user directly.

#### 3.1.4 Event Control

Controls the generation of events for an instance of the interface module. An Event is a message that is sent to the host application to notify it of the occurrence of an internal event.

See GenICam SFNC for more details on event control.

EventSelector selects which particular Event to control.

##### 3.1.4.1 EventControl

|  Name | EventControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains Event control features.

##### 3.1.4.2 EventSelector

|  Name | EventSelector  |
| --- | --- |
|  Category | EventControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Visibility | Expert  |
| --- | --- |
|  Values | InterfaceListChanged  |

Selects which Event to signal to the host application.

Possible values are:

- InterfaceListChanged: the list of interfaces is updated.

### 3.1.4.3 EventNotification

|  Name | EventNotification[EventSelector]  |
| --- | --- |
|  Category | EventControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Off On Once  |

Activate or deactivate the notification to the host application of the occurrence of the selected Event.

Possible values are:

- Off: The selected Event notification is disabled.
- On: The selected Event notification is enabled.
- Once: The selected Event notification is enabled for one event then return to the Off state.

## 3.2 Interface Module

Contains all features of the Interface module that are independent from the underlying transport technology.

### 3.2.1 Interface Information

Features in this section provide basic information about the Interface Module and its identity. Note that all features in this section are defined read-only.

|  GEN<i>CAM |   | ![img-14.jpeg](img-14.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

3.2.1.1 InterfaceInformation

|  Name | InterfaceInformation  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains all Interface Information features of the Interface module.

3.2.1.2 InterfaceID

|  Name | InterfaceID  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

GenTL Producer wide unique identifier of the selected interface.

Corresponds to the INTERFACE_INFO_ID command of IFGetInfo function.

3.2.1.3 InterfaceDisplayName

|  Name | InterfaceDisplayName  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Recommended  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Interface | IString  |
| --- | --- |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

A user-friendly name of the Interface.

Corresponds to the INTERFACE_INFO_DISPLAYNAME command of IFGetInfo function.

##### 3.2.1.4 InterfaceType

|  Name | InterfaceType  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | GigEVisionCameraLinkCameraLinkHSCoaXPressUSB3VisionCustom  |

Transport layer type of the interface.

Note that these values already follow the updated value list of the "DeviceTLType" feature from GenICam SFNC 2.3. Depending on this value, the transport layer specific features for the chosen transport layer standard have to be considered.

• CameraLink: Camera Link
• CameraLinkHS: Camera Link High Speed
- CoaXPress: CoaXPress

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

- GigEVision: GigE Vision
- USB3Vision: USB3 Vision
- Custom: Custom transport layer

Corresponds to the INTERFACE_INFO_TLTYPE command of IFGetInfo function.

### 3.2.1.5 InterfaceTLVersionMajor

|  Name | InterfaceTLVersionMajor  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | >0  |

Major version number of the transport layer specification the GenTL Producer interface complies with. The TL version of the Interface can be compared with the TL version of the device to assure compatibility.

### 3.2.1.6 InterfaceTLVersionMinor

|  Name | InterfaceTLVersionMinor  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Minor version number of the transport layer specification the GenTL Producer interface complies with. The TL version of the Interface can be compared with the TL version of the device to assure compatibility.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.2.1.7 GevInterfaceGatewaySelector

|  Name | GevInterfaceGatewaySelector  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Selector for the different gateway entries for this interface. The selector is 0-based. Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

### 3.2.1.8 GevInterfaceGateway

|  Name | GevInterfaceGateway[GevInterfaceGatewaySelector]  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

IP address of the selected gateway entry of this interface. Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

### 3.2.1.9 GevInterfaceMACAddress

|  Name | GevInterfaceMACAddress  |
| --- | --- |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Category | InterfaceInformation  |
| --- | --- |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

48-bit MAC address of this interface. Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory, and that the Representation element should be used in the XML to facilitate understanding the data.

3.2.1.10 GevInterfaceSubnetSelector

|  Name | GevInterfaceSubnetSelector  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Selector for the subnet of this interface. Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

3.2.1.11 GevInterfaceSubnetIPAddress

|  Name | GevInterfaceSubnetIPAddress[GevInterfaceSubnetSelector]  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  TLType | GigEVision  |

|  GEN<i>CAM |   | ![img-15.jpeg](img-15.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Interface | IInteger  |
| --- | --- |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

IP address of the selected subnet of this interface. Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

3.2.1.12 GevInterfaceSubnetMask

|  Name | GevInterfaceSubnetMask[GevInterfaceSubnetSelector]  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Subnet mask of the selected subnet of this interface. Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

#### 3.2.2 Device Enumeration

The Device Enumeration section describes all features related to discovery and enumeration of devices belonging to the Interface module.

3.2.2.1 DeviceEnumeration

|  Name | DeviceEnumeration  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Access | Read  |
| --- | --- |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains all Device Enumeration features of the Interface module.

### 3.2.2.2 DeviceUpdateList

|  Name | DeviceUpdateList  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Updates the internal device list. This feature should be readable if the execution cannot be performed immediately. The command then returns and the status can be polled. This feature interacts with the IFUpdateDeviceList function of the GenTL Producer. It is up to the GenTL Consumer to handle access in case both methods are used.

### 3.2.2.3 DeviceUpdateTimeout

|  Name | DeviceUpdateTimeout  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | ms  |
|  Visibility | Expert  |
|  Values | >0  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Specifies timeout for the DeviceUpdateList Command.

#### 3.2.2.4 DeviceSelector

|  Name | DeviceSelector  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Selector for the different devices on this interface. This value only changes on execution of "DeviceUpdateList". The selector is 0-based in order to match the index of the C interface.

#### 3.2.2.5 DeviceID

|  Name | DeviceID[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

Interface wide unique identifier of the selected device. This value only changes on execution of the DeviceUpdateList command.

Corresponds to the IFGetDeviceID function with the index corresponding to "DeviceSelector".

#### 3.2.2.6 DeviceVendorName

|  Name | DeviceVendorName[DeviceSelector]  |
| --- | --- |

|  GEN<i>CAM |   | ![img-16.jpeg](img-16.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Category | DeviceEnumeration  |
| --- | --- |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

Name of the device vendor. This value only changes on execution of the DeviceUpdateList command.

Corresponds to the "DeviceVendorName" feature of the remote device and is retrieved during device discovery.

##### 3.2.2.7 DeviceModelName

|  Name | DeviceModelName[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

Name of the device model. This value only changes on execution of the DeviceUpdateList command.

Corresponds to the "DeviceModelName" feature of the remote device and is retrieved during device discovery.

##### 3.2.2.8 DeviceAccessStatus

|  Name | DeviceAccessStatus[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |

|  GEN<i>CAM |   | ![img-17.jpeg](img-17.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Interface | IEnumeration  |
| --- | --- |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | UnknownReadWriteReadOnlyNoAccessBusyOpenReadWriteOpenReadOnly  |

Gives the device's access status at the moment of the last execution of the DeviceUpdateList command. This value only changes on execution of the DeviceUpdateList command.

- Unknown : Not known to producer.
- ReadWrite: Full access
- ReadOnly: Read-only access
- NoAccess: Not available to connect.
- Busy: The device is already opened by another entity.
- OpenReadWrite : Open in Read/Write mode by this GenTL host
- OpenReadOnly : Open in Read only mode by this GenTL host

3.2.2.9 DeviceSerialNumber

|  Name | DeviceSerialNumber[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Recommended  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Serial number of the remote device. This value only changes on execution of the DeviceUpdateList command.

Corresponds to the "DeviceSerialNumber" feature of the remote device and is retrieved during device discovery. Note that this feature was added in GenICam SFNC 2.0 and later, thus, for remote devices following an older GenICam SFNC version it corresponds to the "DeviceID" feature of the remote device.

#### 3.2.2.10 DeviceUserID

|  Name | DeviceUserID[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Optional  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

User-programmable device identifier of the remote device. This value only changes on execution of the DeviceUpdateList command.

Corresponds to the "DeviceUserID" feature of the remote device and it is usually retrieved during device discovery.

#### 3.2.2.11 DeviceTLVersionMajor

|  Name | DeviceTLVersionMajor[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | >0  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Major version number of the transport layer specification the remote device complies with.

### 3.2.2.12 DeviceTLVersionMinor

|  Name | DeviceTLVersionMinor[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Minor version number of the transport layer specification the remote device complies with.

### 3.2.2.13 GevDeviceIPAddress

|  Name | GevDeviceIPAddress[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Current IP address of the GVCP interface of the selected remote device. This value only changes on execution of the DeviceUpdateList command.

Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

### 3.2.2.14 GevDeviceSubnetMask

|  Name | GevDeviceSubnetMask[DeviceSelector]  |
| --- | --- |

|  GEN<i>CAM |   | ![img-18.jpeg](img-18.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Category | DeviceEnumeration  |
| --- | --- |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Current subnet mask of the GVCP interface of the selected remote device. This value only changes on execution of the DeviceUpdateList command.

Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

3.2.2.15 GevDeviceGateway

|  Name | GevDeviceGateway[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Current gateway IP address of the GVCP interface of the selected remote device. This value only changes on execution of the DeviceUpdateList command.

3.2.2.16 GevDeviceIPConfigurationStatus

|  Name | GevDeviceIPConfigurationStatus[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Recommended  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  TLType | GigEVision  |
| --- | --- |
|  Interface | IEnum  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | DHCP PersistentIP LinkLocal  |

Device IP configuration of the GVCP interface of the selected remote device. This value only changes on execution of the DeviceUpdateList command.

### 3.2.2.17 GevDeviceMACAddress

|  Name | GevDeviceMACAddress[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

48-bit MAC address of the GVCP interface of the selected remote device.

Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory, and that the Representation element should be used in the XML to facilitate understanding the data.

### 3.2.2.18 GevDeviceCurrentControlMode

|  Name | GevDeviceCurrentControlMode[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Optional  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  TLType | GigEVision  |
| --- | --- |
|  Interface | IEnum  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Open ControlAccess ExclusiveAccess NoAccess  |

The current control mode of the device. This value only changes on execution of the DeviceUpdateList command. See also DeviceAccessStatus, which gives a similar TL independent status. The values are.

- Open : The device is open for control or exclusive access.
- ControlAccess: The device is controlled by another host, but switchover or readonly access is possible.
- ExclusiveAccess: The device is under exclusive access by a host and cannot be accessed by another.
- NoAccess: The device cannot be accessed, for instance it may be a GigE Vision device on a subnet different from the interface.

### 3.2.2.19 GevApplicationSwitchoverKey

|  Name | GevApplicationSwitchoverKey[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Optional  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

|  GEN<i>CAM |   | ![img-19.jpeg](img-19.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Application switchover key to use when requesting ControlAccess switchover. Setup of the key for switchover is done via device features in the device by a host connected in ExclusiveAccess mode.

3.2.2.20 GevDeviceForceIP

|  Name | GevDeviceForceIP[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Apply the force IP settings (GevDeviceForceIPAddress, GevDeviceForceSubnetMask and GevDeviceForceGateway) in the Device using ForceIP command.

This command is only accepted by a device showing ReadWrite DeviceAccessStatus. The IP change is not persistent in the device.

3.2.2.21 GevDeviceForceIPAddress

|  Name | GevDeviceForceIPAddress[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Static IP address to set for the GVCP interface of the remote device.

|  GEN<i>CAM |   | ![img-20.jpeg](img-20.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

3.2.2.22 GevDeviceForceSubnetMask

|  Name | GevDeviceForceSubnetMask[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Static subnet mask to set for GVCP interface of the remote device.

3.2.2.23 GevDeviceForceGateway

|  Name | GevDeviceForceGateway[DeviceSelector]  |
| --- | --- |
|  Category | DeviceEnumeration  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Static gateway IP address to set for the GVCP interface of the remote device.

#### 3.2.3 Action Control

Features in this section provide give access to the Action Control features within the Interface Module.

|  GEN<i>CAM |   | ![img-21.jpeg](img-21.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

3.2.3.1 ActionControl

|  Name | ActionControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains all Action Control features of the Interface module.

3.2.3.2 ActionCommand

|  Name | ActionCommand  |
| --- | --- |
|  Category | ActionControl  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Send ActionCommand to device(s).

3.2.3.3 ActionDeviceKey

|  Name | ActionDeviceKey  |
| --- | --- |
|  Category | ActionControl  |
|  Level | Recommended  |
|  TLType | GigEVision  |

|  ![img-22.jpeg](img-22.jpeg) GEN<i>CAM |   | ![img-23.jpeg](img-23.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Interface | IInteger  |
| --- | --- |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

The Action Command Device Key to use in the Action Command.

##### 3.2.3.4 ActionGroupKey

|  Name | ActionGroupKey  |
| --- | --- |
|  Category | ActionControl  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

The Action Command Group Key to use in the Action Command.

##### 3.2.3.5 ActionGroupMask

|  Name | ActionGroupMask  |
| --- | --- |
|  Category | ActionControl  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

The Action Command Group Mask to use in the Action Command.

#### 3.2.3.6 ActionScheduledTimeEnable

|  Name | ActionScheduledTimeEnable  |
| --- | --- |
|  Category | ActionControl  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Specifies if a time enabled Action Command should be given.

#### 3.2.3.7 ActionScheduledTime

|  Name | ActionScheduledTime  |
| --- | --- |
|  Category | ActionControl  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Specifies the time a time enabled Action Command is scheduled.

#### 3.2.3.8 GevActionDestinationIPAddress

|  Name | GevActionDestinationIPAddress  |
| --- | --- |

|  GEN<i>CAM |   | ![img-24.jpeg](img-24.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Category | ActionControl  |
| --- | --- |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Specifies destination the IP address for the Action Command. This can be any valid destination address (thus including broadcast addresses for this interface).

#### 3.2.4 GenICam Control

This chapter provides the necessary features to use the GenICam feature tree of the Interface module.

Note: In case of discrepancy between the features described in this chapter and the “GenICam Standard text” the GenTL SFNC document prevails.

3.2.4.1 Root

|  Name | Root  |
| --- | --- |
|  Category | None  |
|  Level | Mandatory  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Provides the Root of the GenICam features tree.

3.2.4.2 InterfacePort

|  Name | InterfacePort  |
| --- | --- |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Category | None  |
| --- | --- |
|  Level | Mandatory  |
|  Interface | IPort  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | -  |

The GenICam port through which the Interface module is accessed.

Note that InterfacePort is a port node (not a feature node) and is generally not accessed by the end user directly.

### 3.2.5 Event Control

Controls the generation of events for an instance of the interface module. An Event is a message that is sent to the host application to notify it of the occurrence of an internal event.

See GenICam SFNC for more details on event control.

EventSelector selects which particular Event to control.

#### 3.2.5.1 EventControl

|  Name | EventControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains Event control features.

#### 3.2.5.2 EventSelector

|  Name | EventSelector  |
| --- | --- |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Category | EventControl  |
| --- | --- |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | InterfaceLost DeviceListChanged  |

Selects which Event to signal to the host application.

Possible values are:

- InterfaceLost: Raised when the interface connection is lost.
- DeviceListChanged: The list of devices is updated.

### 3.2.5.3 EventNotification

|  Name | EventNotification[EventSelector]  |
| --- | --- |
|  Category | EventControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Off On Once  |

Activate or deactivate the notification to the host application of the occurrence of the selected Event.

Possible values are:

- Off: The selected Event notification is disabled.
- On: The selected Event notification is enabled.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

- Once: The selected Event notification is enabled for one event then return to Off state

### 3.3 Device Module

Contains all features of the Device module that are independent from the underlying transport technology. Do not mistake the features of the Device module with the features of the remote device.

### 3.3.1 Device Information

Features in this section provide basic information about the Device module and its identity. Note that all features in this section are defined read-only.

#### 3.3.1.1 Device Information

|  Name | DeviceInformation  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category that contains all Device Information features of the Device module.

#### 3.3.1.2 DeviceID

|  Name | DeviceID  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Mandatory  |
|  Interface | ISString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Values | Any NULL-terminated string  |
| --- | --- |

Interface-wide unique identifier of this device.

Corresponds to the DEVICE_INFO_ID command of DevGetInfo function.

### 3.3.1.3 DeviceSerialNumber

|  Name | DeviceSerialNumber  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Recommended  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

Serial number of the remote device.

Corresponds to the "DeviceSerialNumber" feature of the remote device and usually is retrieved via the bootstrap register of the remote device. Note that this feature has been added in GenICam SFNC 2.0, thus, for remote devices following an older GenICam SFNC version it corresponds to the "DeviceID" feature of the remote device.

Corresponds to the DEVICE_INFO_SERIAL_NUMBER command of DevGetInfo function.

### 3.3.1.4 DeviceUserID

|  Name | DeviceUserID  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Optional  |
|  Interface | IString  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

User-programmable device identifier of the remote device.

Corresponds to the “DeviceUserID” feature of the remote device and usually it is retrieved via the bootstrap register of the remote device.

Corresponds to the DEVICE_INFO_USER_DEFINED_NAME command of DevGetInfo function.

### 3.3.1.5 DeviceVendorName

|  Name | DeviceVendorName  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Name of the remote device vendor.

Corresponds to the DEVICE_INFO_VENDOR command of DevGetInfo function.

### 3.3.1.6 DeviceModelName

|  Name | DeviceModelName  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Mandatory  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Name of the remote device model.

Corresponds to the DEVICE_INFO_MODEL command of DevGetInfo function.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.3.1.7 DeviceFamilyName

|  Name | DeviceFamilyName  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Recommended  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Name of the product family of the remote device model.

Corresponds to the "DeviceFamilyName" feature of the remote device and is usually retrieved via the bootstrap register of the remote device.

### 3.3.1.8 DeviceVersion

|  Name | DeviceVersion  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Recommended  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

The version of the remote device model.

Corresponds to the "DeviceVersion" feature of the remote device and is usually retrieved via the bootstrap register of the remote device.

Corresponds to the DEVICE_INFO_VERSION command of DevGetInfo function.

|  GEN<i>CAM |   | ![img-25.jpeg](img-25.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

3.3.1.9 DeviceManufacturerInfo

|  Name | DeviceManufacturerInfo  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Recommended  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Manufacturer information about the remote device.

Corresponds to the "DeviceManufacturerInfo" feature of the remote device and is usually retrieved via the bootstrap register of the remote device.

3.3.1.10 DeviceType

|  Name | DeviceType  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Mandatory  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | GigEVisionCameraLinkCameraLinkHSCoaXPressUSB3VisionCustom  |

Transport layer type of the device.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Note that these values already follow the updated value list of the "DeviceTLType" feature from GenICam SFNC 2.3. Depending on this value, the transport layer specific features for the chosen transport layer standard have to be considered.

- CameraLink: Camera Link
- CameraLinkHS: Camera Link High Speed
- CoaXPress: CoaXPress
- GigEVision: GigE Vision
- USB3Vision: USB3 Vision
- Custom: Custom transport layer

Corresponds to the DEVICE_INFO_TLTYPE command of DevGetInfo function.

##### 3.3.1.11 DeviceDisplayName

|  Name | DeviceDisplayName  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Recommended  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

User readable name of the device. If this is not defined in the device this should be "VENDOR MODEL (ID)".

Corresponds to the DEVICE_INFO_DISPLAYNAME command of DevGetInfo function.

##### 3.3.1.12 DeviceTimestampFrequency

|  Name | DeviceTimestampFrequency  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |

|  **GEN<i>CAM** |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Unit | -  |
| --- | --- |
|  Visibility | Beginner  |
|  Values | >0  |

The tick-frequency of the time stamp clock.

Corresponds to the DEVICE_INFO_TIMESTAMP_FREQUENCY command of DevGetInfo function.

### 3.3.1.13 DeviceAccessStatus

|  Name | DeviceAccessStatus  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Mandatory  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Unknown ReadWrite ReadOnly NoAccess Busy OpenReadWrite OpenReadOnly  |

Gives the device's access status at the moment of the last execution of the DeviceUpdateList command. This value only changes on execution of the DeviceUpdateList command.

- **Unknown**: Not known to producer.
- **ReadWrite**: Full access
- **ReadOnly**: Read-only access
- **NoAccess**: Not available to connect.
- **Busy**: The device is already opened by another entity.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

- OpenReadWrite: Open in Read/Write mode by this GenTL host
- OpenReadOnly: Open in Read access mode by this GenTL host

Corresponds to the DEVICE_INFO_ACCESS_STATUS command of DevGetInfo function.

3.3.1.14 DeviceChunkDataFormat

|  Name | DeviceChunkDataFormat  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | None GigEVision Custom  |

Chunk data format used by the device. This information allows devices based on other technologies or protocols than "standard" ones such as GigE Vision to inform the GenTL Consumer about the chunk data layout they use.

In contrast, one can assume that any generic GenTL Consumer will understand the GigE Vision chunk format because the GigE Vision chunk adapter is readily available.

Note that GenTL Consumers having access to a generic chunk adapter can use this adapter without caring about the actual data layout, provided that the GenTL Producer implements the DSGetBufferChunkData function. However, using the native chunk adapter might typically lead to slightly better performance.

- None: The device does not use chunk data at all.
- GigEVision: The device formats the chunk data using the chunk data format defined by GigE Vision specification version 1.x. The chunk data decoding algorithm (chunk adapter) common for the GigE Vision devices can be used.
- Custom: The device formats the chunk data using a custom, non-standard format. Without a-priori additional knowledge about the device and its implementation, the GenTL Consumer should always use the generic chunk adapter to decode the chunk data, not making any assumptions about the internal chunk data layout.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.3.1.15 DeviceEventDataFormat

|  Name | DeviceEventDataFormat  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | None GigEVision GigEVisionExtendedId Custom  |

Enumeration, informing about the event data format used by the device (meaning the "device events", see event type EVENT_REMOTE_DEVICE (named EVENT_FEATURE_DEVEVENT in GenTL up to version 1.4). This allows devices based on other technologies or protocols than "standard" ones such as GigE Vision to inform the GenTL Consumer about the event data layout they use.

In contrast, one can assume that any generic GenTL Consumer will understand the GigE Vision event format because the GigE Vision event adapter is readily available

Note that GenTL Consumers having access to a generic event adapter can use this adapter without caring about the actual data layout.

- None: The device does not use event data at all.
- GigEVision: The device formats the event data using the event data format defined by GigE Vision specification version 1.x. The event data decoding algorithm (event adapter) common for the GigE Vision devices can be used.
- GigEVisionExtendedId: The device formats the event data using the event data format defined by GigE Vision specification version 2.x. The event data decoding algorithm (event adapter) common for the GigE Vision devices can be used.
- Custom: The device formats the event data using a custom, non-standard format. Without a-priori additional knowledge about the device and its implementation, the GenTL Consumer should always use the generic event adapter to decode the event data, not making any assumptions about the internal event data layout.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.3.1.16 GevDeviceMACAddress

|  Name | GevDeviceMACAddress  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

48-bit MAC address of the GVCP interface of the remote device.

Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory, and that the Representation element should be used in the XML to facilitate understanding the data.

### 3.3.1.17 GevDeviceIPAddress

|  Name | GevDeviceIPAddress  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Current IP address of the GVCP interface of the remote device.

Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

3.3.1.18 GevDeviceSubnetMask

|  Name | GevDeviceSubnetMask  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Current subnet mask of the GVCP interface of the remote device.

Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

3.3.1.19 GevDeviceGateway

|  Name | GevDeviceGateway  |
| --- | --- |
|  Category | DeviceInformation  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Current gateway IP address of the GVCP interface of the remote device.

Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.3.2 Device Control

The Device Control section contains all features related to control specific properties of the Device module.

#### 3.3.2.1 DeviceControl

|  Name | DeviceControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains all Device Control features of the Device module.

#### 3.3.2.2 DeviceEndianessMechanism

|  Name | DeviceEndianessMechanism  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Legacy Standard  |

Identifies the endianess handling mode.

- Legacy: Handling the device endianess according to GenICam Schema 1.0
- Standard: Handling the device endianess according to GenICam Schema 1.1 and later

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Default value is “Legacy”.

Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory, otherwise recommended.

### 3.3.2.3 LinkCommandTimeout

|  Name | LinkCommandTimeout  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Recommended  |
|  Interface | IFloat  |
|  Access | Read/Write  |
|  Unit | us  |
|  Visibility | Guru  |
|  Values | >0  |

Specifies application timeout for the control channel communication. This feature defines the application timeout, and it is related to the device feature DeviceLinkCommandTimeout specifying the maximum time for handling a command in the device. Up to DeviceLinkCommandRetryCount attempts with this timeout are made before a command fails with a timeout error.

### 3.3.2.4 LinkCommandRetryCount

|  Name | LinkCommandRetryCount  |
| --- | --- |
|  Category | DeviceControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | >=0  |

Specifies maximum number of tries before failing the control channel commands.

|  GEN<i>CAM |   | ![img-26.jpeg](img-26.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

#### 3.3.3 Stream Enumeration

The Stream Enumeration section describes all features related to the enumeration of data streams belonging to the Device module.

##### 3.3.3.1 StreamEnumeration

|  Name | StreamEnumeration  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Category that contains all Stream Enumeration features of the Device module.

##### 3.3.3.2 StreamSelector

|  Name | StreamSelector  |
| --- | --- |
|  Category | StreamEnumeration  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | ≥0  |

Selector for the different stream channels. The selector is 0-based in order to match the index of the C interface.

##### 3.3.3.3 StreamID

|  Name | StreamID[StreamSelector]  |
| --- | --- |
|  Category | StreamEnumeration  |

|  GEN<i>CAM |   | ![img-27.jpeg](img-27.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Level | Mandatory  |
| --- | --- |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

Device unique ID for the stream. Not Mandator for non-streaming DeviceCorresponds to the DevGetDataStreamID function with the index corresponding to "StreamSelector".

#### 3.3.4 GenICam Control

This chapter provides the necessary features to use the GenICam feature tree of the Device module.

Note: In case of discrepancy between the features described in this chapter and the “GenICam Standard text” the GenTL SFNC document prevails.

##### 3.3.4.1 Root

|  Name | Root  |
| --- | --- |
|  Category | None  |
|  Level | Mandatory  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Provides the Root of the GenICam features tree.

##### 3.3.4.2 DevicePort

|  Name | DevicePort  |
| --- | --- |
|  Category | None  |
|  Level | Mandatory  |

|  GEN<i>CAM |   | ![img-28.jpeg](img-28.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Interface | IPort  |
| --- | --- |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | -  |

The GenICam port through which the Device module is accessed. Note that DevicePort is a port node (not a feature node) and is generally not accessed by the end user directly.

#### 3.3.5 Event Control

Controls the generation of events for an instance of the Device module. An Event is a message that is sent to the host application to notify it of the occurrence of an internal event.

See GenICam SFNC for more details on event control.

EventSelector selects which particular Event to control.

##### 3.3.5.1 EventControl

|  Name | EventControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains Event control features.

##### 3.3.5.2 EventSelector

|  Name | EventSelector  |
| --- | --- |
|  Category | EventControl  |
|  Level | Recommended  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Interface | IEnumeration  |
| --- | --- |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | DeviceLost  |

Selects which Event to signal to the host application.

Possible values are:

- DeviceLost: Raised when the local host looses connection to the physical (remote) device.

### 3.3.5.3 EventNotification

|  Name | EventNotification[EventSelector]  |
| --- | --- |
|  Category | EventControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Off On Once  |

Activate or deactivate the notification to the host application of the occurrence of the selected Event.

Possible values are:

- Off: The selected Event notification is disabled.
- On: The selected Event notification is enabled.
- Once: The selected Event notification is enabled for one event then return to Off state.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.4 Data Stream Module

Contains all features of the Data Stream module that are independent from the underlying transport technology.

#### 3.4.1 Stream Information

Features in this section provide basic information about the Data Stream module and its identity.

##### 3.4.1.1 Stream Information

|  Name | StreamInformation  |
| --- | --- |
|  Category | Root  |
|  Level | Mandatory  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains all Stream Information features of the Data Stream module.

##### 3.4.1.2 StreamID

|  Name | StreamID  |
| --- | --- |
|  Category | StreamInformation  |
|  Level | Mandatory  |
|  Interface | ISString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

Device unique ID for the data stream.

Corresponds to the STREAM_INFO_ID command of DSGetInfo function.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.4.1.3 StreamType

|  Name | StreamType  |
| --- | --- |
|  Category | StreamInformation  |
|  Level | Mandatory  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | GigEVision CameraLink CameraLinkHS CoaXPress USB3Vision Custom  |

Transport layer type of the Data Stream.

Note that these values already follow the updated value list of the "DeviceTLType" feature from GenICam SFNC 2.3. Depending on this value, the transport layer specific features for the chosen transport layer standard have to be considered.

- CameraLink: Camera Link
- CameraLinkHS: Camera Link High Speed
- CoaXPress: CoaXPress
- GigEVision: GigE Vision
- USB3Vision: USB3 Vision
- Custom: Custom transport layer

Corresponds to the STREAM_INFO_TLTYPE command of DSGetInfo function.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.4.2 Device Stream Channel Control

3.4.2.1 DeviceStreamChannelControl

|  Name | DeviceStreamChannelControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category containing features to control the stream channel shared between the remote device and the GenTL Producer's data stream module. Applicable for GigE Vision stream channels, and operating on the boot strap registers of the device since the nodemap for the device is not accessible to the GenTL producer.

3.4.2.2 DeviceStreamChannelPacketSize

|  Name | DeviceStreamChannelPacketSize  |
| --- | --- |
|  Category | DeviceStreamChannelControl  |
|  Level | Recommended  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | B  |
|  Visibility | Expert  |
|  Values | >0  |

Specifies the stream packet size, in bytes, to send on the selected channel for a transmitter or specifies the maximum packet size supported by a receiver. Controls the packet size configuration of the remote device and if needed the GenTL Producer.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.4.2.3 DeviceStreamChannelPacketSizeMin

|  Name | DeviceStreamChannelPacketSizeMin  |
| --- | --- |
|  Category | DeviceStreamChannelControl  |
|  Level | Optional  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | B  |
|  Visibility | Guru  |
|  Values | >0  |

Controls desired minimum of the packet size feature to be used for the stream channel. Affects both the direct control of the packet size as well as the negotiation algorithm. The GenTL Consumer can set the value in accordance with the known limits of the remote device or apply further restrictions e.g. based on additional knowledge of the system.

### 3.4.2.4 DeviceStreamChannelPacketSizeMax

|  Name | DeviceStreamChannelPacketSizeMax  |
| --- | --- |
|  Category | DeviceStreamChannelControl  |
|  Level | Optional  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | B  |
|  Visibility | Guru  |
|  Values | >0  |

Controls desired maximum of the packet size feature to be used for the stream channel. Affects both the direct control of the packet size as well as the negotiation algorithm. The GenTL Consumer can set the value in accordance with the known limits of the remote device or apply further restrictions e.g. based on additional knowledge of the system.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.4.2.5 DeviceStreamChannelPacketSizeInc

|  Name | DeviceStreamChannelPacketSizeInc  |
| --- | --- |
|  Category | DeviceStreamChannelControl  |
|  Level | Optional  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | B  |
|  Visibility | Guru  |
|  Values | >0  |

Controls desired increment of the packet size feature to be used for the stream channel. Affects both the direct control of the packet size as well as the negotiation algorithm. The GenTL Consumer can set the value in accordance with the known limits of the remote device or apply further restrictions e.g. based on additional knowledge of the system.

### 3.4.2.6 DeviceStreamChannelNegotiatePacketSize

|  Name | DeviceStreamChannelNegotiatePacketSize  |
| --- | --- |
|  Category | DeviceStreamChannelControl  |
|  Level | Optional  |
|  TLType | GigEVision  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Starts negotiation for the optimal packet size considering the remote device, host and their connection path. The negotiation result is applied on the device and reflected in DeviceStreamChannelPacketSize. If the negotiation fails, the algorithm attempts to revert the configuration to the initial packet size value.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.4.3 Buffer Handling Control

Features in this section provide control over the buffers within the acquisition engine of a data stream.

#### 3.4.3.1 BufferHandlingControl

|  Name | BufferHandlingControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Contains all features of the Data Stream module that control the used buffers.

#### 3.4.3.2 StreamAnnouncedBufferCount

|  Name | StreamAnnouncedBufferCount  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Number of announced (known) buffers on this stream. This value is volatile. It may change if additional buffers are announced and/or buffers are revoked by the GenTL Consumer.

Corresponds to the STREAM_INFO_NUM_ANNOUNCED command of DSGetInfo function.

#### 3.4.3.3 StreamBufferHandlingMode

|  Name | StreamBufferHandlingMode  |
| --- | --- |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Category | BufferHandlingControl  |
| --- | --- |
|  Level | Mandatory  |
|  Interface | IEnumeration  |
|  Access | Read(/Write)  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | OldestFirst (Mandatory) OldestFirstOverwrite NewestOnly Default (Deprecated)  |

Available buffer handling modes of this Data Stream.

![img-29.jpeg](img-29.jpeg)

Figure 3-3.1: Buffer handling mode "OldestFirst"

- OldestFirst (Mandatory): The application always gets the buffer from the head of the Output Buffer Queue (thus, the oldest available one). If the Output Buffer Queue is empty, the application waits for a newly acquired buffer until the timeout expires.

When data for a new buffer is available, the acquisition engine looks for any available buffer in the Input Buffer Pool, fills it, and appends it to the tail of the Output Buffer Queue. If the Input Buffer Pool is empty, the new data is dropped.

This buffer handling mode is typically used if every image frame is to be acquired and the mean processing time is lower than acquisition time. No buffer is discarded or overwritten in the Output Buffer Queue and all filled buffers are delivered in the order they were acquired.

Acquisition Engine

![img-30.jpeg](img-30.jpeg)

Buffer Delivery

![img-31.jpeg](img-31.jpeg)

Figure 3-3.2: Buffer handling mode “OldestFirstOverwrite”

- OldestFirstOverwrite (Recommended): The application always gets the buffer from the head of the Output Buffer Queue (thus, the oldest available one). If the Output Buffer Queue is empty, the application waits for a newly acquired buffer until the timeout expires.

When data for a new buffer is available, the acquisition engine looks for any available buffer in the Input Buffer Pool, fills it, and appends it to the tail of the Output Buffer Queue. If the Input Buffer Pool is empty and the Output Buffer Queue is not empty, it discards the head of the Output Buffer Queue (i.e., the oldest buffer), overwrites it with the new data, and appends it to the tail of the Output Buffer Queue. If the Input Buffer Pool and the Output Buffer Queue are empty, the new data is dropped.

This buffer handling mode is typically used if not every image frame is to be acquired and the application may not fall behind.

- NewestOnly (Recommended): The application always gets the latest completed buffer (the newest one). If the Output Buffer Queue is empty, the application waits for a newly acquired buffer until the timeout expires.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

This buffer handling mode is typically used in a live display GUI where it is important that there is no lag between camera and display.

![img-32.jpeg](img-32.jpeg)

![img-33.jpeg](img-33.jpeg)

Figure 3-3.3. Buffer Handling Mode "Newest Only".

- Default (Deprecated): Same behavior as "OldestFirst".

Note that depending on the actual payload not only pure images, but any kind of data can be acquired.

### 3.4.3.4 StreamAnnounceBufferMinimum

|  Name | StreamAnnounceBufferMinimum  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Mandatory  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Minimal number of buffers to announce to enable selected buffer handling mode.

Corresponds to the STREAM_INFO_BUF_ANNOUNCE_MIN command of DSGetInfo function.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.4.3.5 StreamDeliveredFrameCount

|  Name | StreamDeliveredFrameCount  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Number of delivered frames since last acquisition start. It is not reset until the stream is closed.

Corresponds to the STREAM_INFO_NUM_DELIVERED command of DSGetInfo function.

### 3.4.3.6 StreamLostFrameCount

|  Name | StreamLostFrameCount  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Number of lost frames due to queue underrun. This number is initialized with zero at the time the stream is opened and incremented every time the data could not be acquired because there was no buffer in the input buffer pool. It is not reset until the stream is closed.

Corresponds to the STREAM_INFO_NUM_UNDERRUN command of DSGetInfo function.

### 3.4.3.7 StreamInputBufferCount

|  Name | StreamInputBufferCount  |
| --- | --- |
|  Category | BufferHandlingControl  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Level | Optional  |
| --- | --- |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Number of buffers in the input buffer pool plus the buffers(s) currently being filled.

Corresponds to the STREAM_INFO_NUM_QUEUED command of DSGetInfo function.

##### 3.4.3.8 StreamOutputBufferCount

|  Name | StreamOutputBufferCount  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Number of buffers in the output buffer queue.

Corresponds to the STREAM_INFO_NUM_AWAIT_DELIVERY command of DSGetInfo function.

##### 3.4.3.9 StreamStartedFrameCount

|  Name | StreamStartedFrameCount  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |

|  GEN<I>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Visibility | Expert  |
| --- | --- |
|  Values | ≥0  |

Number of frames started in the acquisition engine. This number is incremented every time in case of a new buffer is started and then to be filled (data written to) regardless even if the buffer is later delivered to the user or discarded for any reason. This number is initialized with 0 at the time the stream is opened. It is not reset until the stream is closed.

Corresponds to the STREAM_INFO_NUM_STARTED command of DSGetInfo function.

### 3.4.3.10 PayloadSize

|  Name | PayloadSize  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Byte  |
|  Visibility | Expert  |
|  Values | >0  |

Size of the expected data in bytes. Note that this feature "overwrites" the PayloadSize of the remote device, see also sections "Data Payload Delivery" and "Allocate Memory" of the GenICam GenTL standard.

Corresponds to the STREAM_INFO_PAYLOAD_SIZE command of DSGetInfo function.

### 3.4.3.11 StreamIsGrabbing

|  Name | StreamIsGrabbing  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Recommended  |
|  Interface | IBoolean  |
|  Access | Read  |
|  Unit |   |
|  Visibility | Expert  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Values | True False  |
| --- | --- |

Flag indicating whether the acquisition engine is started or not. This is independent from the acquisition status of the remote device.

Corresponds to the STREAM_INFO_IS_GRABBING command of DSGetInfo function.

### 3.4.3.12 StreamChunkCountMaximum

|  Name | StreamChunkCountMaximum  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit |   |
|  Visibility | Expert  |
|  Values | ≥0  |

Maximum number of chunks to be expected in a buffer (can be used to allocate the array for the DSGetBufferChunkData function).

Corresponds to the STREAM_INFO_NUM_CHUNKS_MAX command of DSGetInfo function.

### 3.4.3.13 StreamBufferAlignment

|  Name | StreamBufferAlignment  |
| --- | --- |
|  Category | BufferHandlingControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Byte  |
|  Visibility | Expert  |
|  Values | ≥0  |

Alignment size in bytes of the buffers passed to DSAnnounceBuffer.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

If a buffer is passed to DSAnnounceBuffer which is not aligned according to the alignment size it is up to the Producer to either reject the buffer and return a GC_ERR_INVALID_BUFFER error code or to cope with a potential overhead and use the unaligned buffer as is.

Corresponds to the STREAM_INFO_BUF_ALIGNMENT command of DSGetInfo function.

#### 3.4.4 GenICam Control

This chapter provides the necessary features to use the GenICam feature tree of the Device module.

Note: In case of discrepancy between the features described in this chapter and the “GenICam Standard text” the GenTL SFNC document prevails.

3.4.4.1 Root

|  Name | Root  |
| --- | --- |
|  Category | None  |
|  Level | Mandatory  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Provides the Root of the GenICam features tree.

3.4.4.2 StreamPort

|  Name | StreamPort  |
| --- | --- |
|  Category | None  |
|  Level | Mandatory  |
|  Interface | IPort  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | -  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

The GenICam port through which the Data Stream module is accessed.

Note that StreamPort is a port node (not a feature node) and is generally not accessed by the end user directly.

### 3.4.5 Event Control

Controls the generation of events for an instance of the buffer module. An Event is a message that is sent to the host application to notify it of the occurrence of an internal event.

See GenICam SFNC for more details on event control.

EventSelector selects which particular Event to control

#### 3.4.5.1 EventControl

|  Name | EventControl  |
| --- | --- |
|  Category | Root  |
|  Level | Recommended  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains Event control features.

#### 3.4.5.2 EventSelector

|  Name | EventSelector  |
| --- | --- |
|  Category | EventControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | NewBufferData  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|   | TransferEnd BufferTooSmall BuffersDiscarded  |
| --- | --- |

Selects which Event to signal to the host application.

Possible values are:

- **NewBufferData**: A new buffer is available.
- **TransferEnd**: The transfer of a data for new buffer finished; this is not directly related with delivering the buffer, the data might be appended to end of Output Buffer Queue, dropped, etc., depending on the buffer handling mode and acquisition engine status.
- **BufferTooSmall**: The buffer was too small to receive the expected amount of data.
- **BuffersDiscarded**: Buffers discarded by GenTL or device. This event could optionally carry two numeric child data fields EventBuffersDiscardedDeviceCount and EventBuffersDiscardedProducerCount.

**EventBuffersDiscardedDeviceCount**: Number of buffers discarded by the device since last fired instance of this event (the producer would get to know about this for example by observing a gap in the block_id sequence)

**EventBuffersDiscardedProducerCount**: Number of buffers discarded by the producer since last fired instance of this event (this would happen e.g. if there are no free buffers available or if given buffer handling mode requires discarding old buffers etc.)

### 3.4.5.3 EventNotification

|  Name | EventNotification[EventSelector]  |
| --- | --- |
|  Category | EventControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Unit | -  |
| --- | --- |
|  Visibility | Expert  |
|  Values | Off On Once  |

Activate or deactivate the notification to the host application of the occurrence of the selected Event.

Possible values are:

- Off: The selected Event notification is disabled.
- On: The selected Event notification is enabled.
- Once: The selected Event notification is enabled for one event then return to Off state.

### 3.5 Buffer Module

Contains all features of the Buffer module that are independent from the underlying transport technology. Since for the Buffer module the GenTL Port is optional, all features listed in this chapter are optional.

#### 3.5.1 Buffer Information

Features in this section provide basic information about the Buffer module.

##### 3.5.1.1 BufferInformation

|  Name | BufferInformation  |
| --- | --- |
|  Category | Root  |
|  Level | Optional  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Category that contains all Buffer Information features of the Buffer module.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Note that this category depends whether a Port access is provided through the "BufferPort" feature.

### 3.5.1.2 BufferUserData

|  Name | BufferUserData  |
| --- | --- |
|  Category | BufferInformation  |
|  Level | Optional (but mandatory if Port access provided)  |
|  Interface | Integer  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Pointer to user data casted to an integer number referencing GenTL Consumer specific data. It is reflecting the pointer provided by the user data pointer at buffer announcement. This allows the GenTL Consumer to attach information to a buffer.

Note that according to the GenICam GenTL standard, this feature is mandatory if a Port access is provided through the “BufferPort” feature.

Corresponds to the BUFFER_INFO_USER_PTR command of DSGetBufferInfo function.

### 3.5.1.3 BufferType

|  Name | BufferType  |
| --- | --- |
|  Category | BufferInformation  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | GigEVision CameraLink CameraLinkHS CoaXPress USB3Vision  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|   | Custom  |
| --- | --- |

Transport layer type of the buffer.

Note that these values already follow the updated value list of the "DeviceTLType" feature from GenICam SFNC 2.3. Depending on this value, the transport layer specific features for the chosen transport layer standard have to be considered.

- CameraLink: Camera Link
- CameraLinkHS: Camera Link High Speed
- CoaXPress: CoaXPress
- GigEVision: GigE Vision
- USB3Vision: USB3 Vision
- Custom: Custom transport layer

Corresponds to the BUF_INFO_TLTYPE command of DSGetBufferInfo function.

### 3.5.1.4 BufferSize

|  Name | BufferSize  |
| --- | --- |
|  Category | BufferInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Byte  |
|  Visibility | Expert  |
|  Values | ≥0  |

Size of the buffer in bytes.

Corresponds to the BUF_INFO_SIZE command of DSGetBufferInfo function.

### 3.5.2 Buffer Data Information

Features in this section provide information about the currently filled data in the buffers. Note that for multipart buffers the BufferPartSelector is used to extract information for each part of the buffer.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.5.2.1 BufferDataInformation

|  Name | BufferDataInformation  |
| --- | --- |
|  Category | Root  |
|  Level | Optional  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | -  |

Contains all Buffer Data Information features of the Buffer module.

Note that this category depends whether a Port access is provided through the "BufferPort" feature.

### 3.5.2.2 BufferData

|  Name | BufferData  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional (but mandatory if Port access provided)  |
|  Interface | IRegister  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

Entire buffer data.

Note that according to the GenICam GenTL standard, this feature is mandatory if a Port access is provided through the "BufferPort" feature.

Corresponds to the BUFFER_INFO_BASE command of DSGetBufferInfo function.

### 3.5.2.3 BufferTimeStamp

|  Name | BufferTimeStamp  |
| --- | --- |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Category | BufferInformation  |
| --- | --- |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Timestamp the buffer was acquired. The unit is device/implementation dependent. In case the technology and/or the device does not support this for example under Windows a QueryPerformanceCounter can be used.

Corresponds to the BUF_INFO_TIMESTAMP command of DSGetBufferInfo function.

### 3.5.2.4 BufferNewData

|  Name | BufferNewData  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Flag to indicate that the buffer contains new data since the last delivery.

Corresponds to the BUFFER_INFO_NEW_DATA command of DSGetBufferInfo function.

### 3.5.2.5 BufferIsQueued

|  Name | BufferIsQueued  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IBoolean  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Access | Read  |
| --- | --- |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Flag to indicate if the buffer is in the input pool or output buffer queue.

Corresponds to the BUFFER_INFO_IS_QUEUED command of DSGetBufferInfo function.

### 3.5.2.6 BufferIsAcquiring

|  Name | BufferIsAcquiring  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Flag to indicate that the buffer is currently being filled with data.

Corresponds to the BUFFER_INFO_IS_ACQUIRING command of DSGetBufferInfo function.

### 3.5.2.7 BufferIsIncomplete

|  Name | BufferIsIncomplete  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Values | True False  |
| --- | --- |

Flag to indicate that a buffer was filled but an error occurred during that process.

Corresponds to the BUFFER_INFO_IS_INCOMPLETE command of DSGetBufferInfo function.

### 3.5.2.8 BufferPayloadType

|  Name | BufferPayloadType  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Unknown Image RawData File ChunkData JPEG JPEG2000 H264 ChunkOnly MultiPart GenDC  |

Payload type of the data.

- **Unknown**: The GenTL Producer is not aware of the payload type of the data in the provided buffer. For the GenTL Consumer perspective this can be handled as raw data.
- **Image**: The buffer payload contains pure image data. In particular, no chunk data is attached to the image.
- **RawData**: The buffer payload contains raw, unspecified data. For instance, this can be used to send acquisition statistics.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

- File: The buffer payload contains data of a file. It is used to transfer files, such as JPEG compressed images, which can be stored by the GenTL Producer directly to a hard disk. The user might get a hint how to interpret the buffer by the filename by the "BufferFileName" feature.
- ChunkData: The buffer payload contains chunk data which can be parsed. The chunk data type might be reported through SFNC or deduced from the technology the device is based on. Note that the chunk data can also contain an image. The GenTL Producer should report the presence, position (offset in the buffer) and properties of the image through corresponding BUFFER_INFO_CMD commands.
- JPEG: The buffer payload is a Jpeg formatted image.
- JPEG2000: The buffer payload is a JPEG2000 formatted image.
- H264: The buffer payload is H.264 formatted image data.
- ChunkOnly: The buffer only contains chunk data.
- MultiPart: The buffer payload has multiple parts.
- GenDC: The buffer payload contains a GenDC container.

Corresponds to the BUFFER_INFO_PAYLOADTYPE command of DSGetBufferInfo function.

### 3.5.2.9 BufferNumberOfParts

|  Name | BufferNumberOfParts  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

The number of parts in the current buffer as delivered by the transport mechanism. For non-multipart this is 0, giving that it is not a multipart buffer.

### 3.5.2.10 BufferPartSelector

|  Name | BufferPartSelector  |
| --- | --- |
|  Category | BufferDataInformation  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Level | Optional  |
| --- | --- |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

The buffer part to extract information from. For non-multipart the value is 0. The maximum value should be dynamic and reflect the number of parts possible to index.

##### 3.5.2.11 BufferSizeFilled

|  Name | BufferSizeFilled  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Byte  |
|  Visibility | Expert  |
|  Values | ≥0  |

Number of bytes written into the buffer last time it was filled. This value is reset to 0 when the buffer is placed into the Input Buffer Pool.

Corresponds to the BUFFER_INFO_SIZE_FILLED command of DSGetBufferInfo function.

##### 3.5.2.12 BufferPartDataType

|  Name | BufferPartDataType[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |

|  GEN<i>CAM |   | ![img-34.jpeg](img-34.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Visibility | Expert  |
| --- | --- |
|  Values | Image2DBiplanarImagePlane2DTriplanarImagePlane2DQuadPlanarImagePlane2DImage3DBiplanarImagePlane3DTriplanarImagePlane3DQuadPlanarImagePlane3DConfidenceMapChunkJpegJpeg2000Custom  |

The data type of the part.

3.5.2.13 BufferPartSourceIDValue

|  Name | BufferPartSourceIDValue[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

The Source ID type of the part.

3.5.2.14 BufferPartRegionIDValue

|  Name | BufferPartRegionIDValue[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |

|  GEN<i>CAM |   | ![img-35.jpeg](img-35.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Access | Read  |
| --- | --- |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

The Region ID type of the part.

3.5.2.15 BufferPartComponentIDValue

|  Name | BufferPartComponentIDValue[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

The Component ID type of the part.

3.5.2.16 BufferWidth

|  Name | BufferWidth[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Width of the data in the buffer in number of pixels.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

This information refers for example to the width entry in the GigE Vision image stream data leader. For other technologies, this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_WIDTH command of DSGetBufferInfo function.

### 3.5.2.17 BufferHeight

|  Name | BufferHeight[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Height of the data in the buffer in number of pixels as configured. For variable size images this is the max Height of the buffer.

For example this information refers to the height entry in the GigE Vision image stream data leader. For other technologies this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_HEIGHT command of DSGetBufferInfo function.

### 3.5.2.18 BufferXOffset

|  Name | BufferXOffset[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

XOffset of the data in the buffer in number of pixels from the image origin to handle areas of interest.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

This information refers for example to the information provided in the GigE Vision image stream data leader. For other technologies this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_XOFFSET command of DSGetBufferInfo function.

### 3.5.2.19 BufferYOffset

|  Name | BufferYOffset[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

YOffset of the data in the buffer in number of lines from the image origin to handle areas of interest.

This information refers for example to the information provided in the GigE Vision image stream data leader. For other technologies, this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_YOFFSET command of DSGetBufferInfo function.

### 3.5.2.20 BufferXPadding

|  Name | BufferXPadding[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Byte  |
|  Visibility | Expert  |
|  Values | ≥0  |

XPadding of the data in the buffer in number of bytes.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

This information refers for example to the information provided in the GigE Vision image stream data leader. For other technologies, this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_XPADDING command of DSGetBufferInfo function.

### 3.5.2.21 BufferYPadding

|  Name | BufferYPadding  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Byte  |
|  Visibility | Expert  |
|  Values | ≥0  |

YPadding of the data in the buffer in number of bytes.

This information refers for example to the information provided in the GigE Vision image stream data leader. For other thechnologies, this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_YPADDING command of DSGetBufferInfo function.

### 3.5.2.22 BufferFrameID

|  Name | BufferFrameID  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

A sequentially incremented number of the frame.

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

This information refers for example to the information provided in the GigE Vision image stream block id. For other technologies, this is to be implemented accordingly. The wrap around of this number is transportation technology dependent

Corresponds to the BUFFER_INFO_FRAMEID command of DSGetBufferInfo function.

### 3.5.2.23 BufferImagePresent

|  Name | BufferImagePresent  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Flag to indicate if the current data in the buffer contains image data.

This information refers for example to the information provided in the GigE Vision image stream data leader. For other technologies, this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_IMAGEPRESET command of DSGetBufferInfo function.

### 3.5.2.24 BufferImageOffset

|  Name | BufferImageOffset  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | Byte  |
|  Visibility | Expert  |
|  Values | ≥0  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Offset of the image data from the beginning of the delivered buffer in bytes. Applies for example when delivering the image as part of chunk data or on technologies requiring specific buffer alignment.

Corresponds to the BUFFER_INFO_IMAGEOFFSET command of DSGetBufferInfo function.

3.5.2.25 BufferPixelFormat

|  Name | BufferPixelFormat[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Mono1p Mono2p Mono4p Mono8 Mono8s Mono10 Mono10c3a64 Mono10c3p32 Mono10g12 Mono10msb Mono10p Mono10pmsb Mono10s Mono12 Mono12g Mono12msb Mono14 Mono16 R8 G8 B8 RGB8 RGB8_Planar  |

|  GEN<i>CAM |   | ![img-36.jpeg](img-36.jpeg) emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|   | RGB8a32RGBa8RGB10RGB10_PlanarRGB10g32RGB10g32msbRGB10p32RGB10p32msbRGB12RGB12_PlanarRGB16RGB16_PlanarRGB565pBGR10BGR12BGR16BGR565pBGR8BGRa8YUV411_8YUV422_8YUV8YCbCr411_8YCbCr422_8YCbCr601_411_8 YCbCr601_422_8YCbCr601_8YCbCr709_411_8 YCbCr709_422_8YCbCr709_8 YCbCr8BayerBG8BayerGB8BayerGR8BayerRG8BayerBG10BayerBG10g12BayerGB10BayerGB10g12BayerGR10BayerGR10g12  |
| --- | --- |

|   | BayerRG10 BayerRG10g12 BayerBG12 BayerBG12g BayerGB12 BayerGB12g BayerGR12 BayerGR12g BayerRG12 BayerRG12g BayerBG16 BayerGB16 BayerGR16 BayerRG16 Raw16 Raw8 Device-specific - GigE Vision Specific: Mono12Packed BayerGR10Packed BayerRG10Packed BayerGB10Packed BayerBG10Packed BayerGR12Packed BayerRG12Packed BayerGB12Packed BayerBG12Packed RGB10V1Packed BGR10V1Packed RGB12V1Packed  |
| --- | --- |

Format of the pixels provided by the buffer.

Note that the value list already follows the updated value list of the "PixelFormat" feature from GenICam SFNC 2.0, i.e., this feature does not exactly correspond to the

BUFFER_INFO_PIXELFORMAT command of DSGetBufferInfo function in the GenICam GenTL 1.3 standard. For multipart buffers this corresponds to

BUFFER_PART_INFO_DATA_FORMAT for PFNC formatted parts.

Note that only a subset of the possible pixel formats is listed here. The complete list of possible standard pixel formats and their detailed layout can be found in the "Pixel Format Naming

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Convention (PFNC)" specification hosted by the AIA organisation. Refer to the most recent version of that convention for additional information about the construction of a pixel format name.

### 3.5.2.26 BufferDeliveredImageHeight

|  Name | BufferDeliveredImageHeight[BufferPartSelector]  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

The number of lines in the current buffer part as delivered by the transport mechanism. For area scan type images, this is usually the number of lines configured in the device. For variable size linescan images, this number may be lower than the configured image height.

This information refers for example to the information provided in the GigE Vision image stream data trailer. For other technologies, this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_DELIVERED_IMAGEHEIGHT command of DSGetBufferInfo function and BUFFER_PART_INFO_DELIVEREDIMAGEHEIGHT in a DSGetPartInfo function

### 3.5.2.27 BufferDeliveredChunkPayloadSize

|  Name | BufferDeliveredChunkPayloadSize  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

Size of the valid chunk payload data delivered in the buffer.

This information refers for example to the information provided in the GigE Vision image stream data trailer. For other technologies, this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_DELIVERED_CHUNKPAYLOADSIZE command of DSGetBufferInfo function.

### 3.5.2.28 BufferChunkLayoutID

|  Name | BufferChunkLayoutID  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

ID of the chunk data layout delivered in the buffer. Can be used to track changes of the layout data among individual buffers.

This information refers for example to the information provided in the GigE Vision image stream data leader. The chunk layout id serves as an indicator that the chunk layout has changed and the application should re-parse the chunk layout in the buffer. When a chunk layout (availability or position of individual chunks) changes since the last buffer delivered by the device through the same stream, the device MUST change the chunk layout id. As long as the chunk layout remains stable, the camera MUST keep the chunk layout id intact. When switching back to a layout, which was already used before, the camera can use the same id again or use a new id. A chunk layout id value of 0 is invalid. It is reserved for use by cameras not supporting the layout id functionality. The algorithm used to compute the chunk layout id is left as quality of implementation. For other technologies this is to be implemented accordingly.

Corresponds to the BUFFER_INFO_CHUNKLAYOUTID command of DSGetBufferInfo function.

### 3.5.2.29 BufferFileName

|  Name | BufferFileName  |
| --- | --- |
|  Category | BufferDataInformation  |
|  Level | Optional  |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Interface | IString  |
| --- | --- |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Any NULL-terminated string  |

Filename for the file payload data delivered in the buffer.

This information refers for example to the information provided in the GigE Vision image stream data leader. For other technologies, this is to be implemented accordingly. Since this is GigE Vision related information and the filename in GigE Vision is UTF8 coded, this filename is also UTF8 coded.

Corresponds to the BUFFER_INFO_FILENAME command of DSGetBufferInfo function.

#### 3.5.3 GenICam Control

This chapter provides the necessary features to use the GenICam feature tree of the Buffer module.

Note: In case of discrepancy between the features described in this chapter and the “GenICam Standard text” the GenTL SFNC document prevails.

3.5.3.1 Root

|  Name | Root  |
| --- | --- |
|  Category | None  |
|  Level | Optional  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | -  |

Provides the Root of the GenICam features tree.

3.5.3.2 BufferPort

|  Name | BufferPort  |
| --- | --- |

|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Category | None  |
| --- | --- |
|  Level | Optional  |
|  Interface | IPort  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | -  |

The GenICam port through which the Buffer module is accessed.

Note that BufferPort is a port node (not a feature node) and is generally not accessed by the end user directly.

Note that according to the GenICam GenTL standard, this feature is not mandatory. However, if this feature is provided, also the features “BufferData” and “BufferUserData” are mandatory.

|  GEN<I>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

## 4 Acknowledgements

The following companies have participated in the elaboration of the GenICam GenTL Standard Features Naming Convention:

|  Company | Represented by  |
| --- | --- |
|  Active Silicon | Jean-Philippe Arnaud, Chris Beynon  |
|  Allied Vision | Holger Edelbüttel  |
|  Groget | Jan Becvar  |
|  STEMMER IMAGING | Rupert Stelz  |
|  MathWorks | Mark Jones  |
|  MATRIX VISION | Stefan Battmer  |
|  Matrox Imaging | Stephane Maurice  |
|  MVTec Software | Thomas Hopfner, Christoph Zierl  |
|  SICK | Mattias Johannesson  |