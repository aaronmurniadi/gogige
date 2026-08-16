|  GEN<i>CAM |   | ![img-1.jpeg](img-1.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

## 1. Introduction

### 1.1. Motivation

Products, which rely on a serial link for communication, implement a wide variety of proprietary control protocols. Most of these protocols are based on ASCII command strings and ASCII responses or even binary protocols. Proprietary protocols can be integrated into GenICam through the GenICam CLProtocol module, assuming the device manufacturer provides a dynamic link library (DLL) for all supported platforms/operating systems. This DLL does the translation between the camera-specific proprietary control protocol and a GenICam compliant register map, which allows the integration of a device into GenICam.

Providing a manufacturer-specific and platform-specific DLL adds cost and effort:

- It has to be maintained for various platforms and OS versions.
• Device features must be added and updated
- The integration of embedded platforms must be taken into account

A more straightforward approach is to provide a read/write register protocol, which can also run on a serial link and do the register map integration in the camera. There would be only one place to change, the camera firmware, in order to introduce new features. There would be no platform-specific software needed, which would allow the use of embedded devices as the controlling host. This protocol can be packet based and therefore used on other packet-based technologies as well.

Some devices on the market implement serial protocols in a similar way already. The idea is to propose a common approach for implementing a protocol to give new implementers a hint and maybe to allow a de facto standard in the future.

The original idea was to simplify the CLProtocol implementation by providing a protocol description. Because a protocol can potentially be used on other technologies as well, the definition is kept more generic. It can be adjusted to other technologies however the serial link of Camera Link was the first approach.

### 1.2. Objective

The objective of this document is to describe

- a packet-based protocol to read and write registers in a register-based device
- a Bootstrap Register Map (BRM) to provide basic device information
- access to the device's GenICam file
• the technology specific communication configuration