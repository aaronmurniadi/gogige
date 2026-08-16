|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

## 27 Transport Layer Control

This chapter provides the Transport Layer control features.

In general the generic features are directly under the TransportLayerControl category and the other Transport Layer specific features are under their respective sub category.

### 27.1 TLParamsLocked Usage

TLParamsLocked related features are used to lock the state of certain features during acquisition.

By default (for example after a new Device connection), TLParamsLocked should be zero or unlocked in order to allow full device configuration. When the transmitter is configured properly, the receiver sets the TLParamsLocked feature of the transmitter to one which locks its transport layer related features. This allows starting the acquisition. The TLParamsLocked feature must be keep locked as long as the acquisition and transfer are active.

There are different group of features that can be locked. Those groups are selected using the TLParamsLockedSelector. A feature must follow the rules of the group in which it falls. Features that affect other features must follow the rules of the more restrictive group of feature it is affecting. For example binning features also affect image size features. When a feature does not affect any predefined group, the device is free to decide if the feature is locked by TLParamsLocked.

#### Streaming Channel Control

Features that are related to the configuration of the streaming channel must be locked by the TLParamsLocked feature. Streaming channel configuration features are used to change the streaming channel behavior, packet formatting or usage of the physical link.

#### Data Flow Control

When the transport layer features are unlocked, both the AcquisitionStart and TransferStart features should be locked to prevent a user from starting an acquisition in an invalid state.

#### Stream Content

Not all transmitters and receivers have the same capability when it comes to dynamic change in the stream content. By default, all features that can modify the image size or pixel format should be locked by the TLParamsLocked feature. A device can implement the TLParamsLockedSelector and TLParamsLockedState feature to change to allow more precise control on the features locking.

### 27.2 Transport Layer features

This section describes the general transport layer features.