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