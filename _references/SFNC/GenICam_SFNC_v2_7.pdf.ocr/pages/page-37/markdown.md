## 1 Introduction

This document contains the GenICam "Standard Features Naming Convention (SFNC)" that provides a standard features naming convention and a standard behavioral model for the devices based on the GenICam standard. The latest release version of all the GenICam standard documents can be found on the GenICam download page on the EMVA web site (see in particular the section "SFNC (Standard Features Naming Convention)").

In general, the GenICam technology allows exposing arbitrary features of a camera through a unified API and GUI. Each feature can be defined in an abstract manner by its name, interface type, unit of measurement and behavior. The GenApi module of the GenICam standard defines how to write a camera XML description file that describes the features of a device.

The usage of GenApi alone could be sufficient to make all the features of a camera or a device accessible through the GenICam API. However if the user wants to write generic and portable software for a whole class of cameras or devices and be interoperable, then GenApi alone is not sufficient and the software and the device vendors have to agree on a common naming convention for the standard features. This is the role of the GenICam "Standard Features Naming Convention (SFNC)" to provide a common set of features, their name, and to define a standard behavior for them.

The Standard Features Naming Convention of GenICam is targeting maximum usability and interoperability by existing and future transport layer technologies. In order to achieve this, it provides the definition of standard and transport layer agnostic features and their expected behavior. The goal is to cover and to standardize the naming convention used in all the basic use cases of the devices where the implementation by different vendors would be very similar anyway.

**To be GenICam compliant a product must provide a GenICam XML file.**

- The GenICam XML file must be compatible with the latest GenApi and schema.
- The GenICam XML file must include all the public features of the product it describes.
- The GenICam XML features must follow the Standard Features Naming Convention whenever applicable or possible.

Those requirements ensure that the users can rely on a complete, consistent and portable feature set for its device and that those features are always accessible in a standard way.