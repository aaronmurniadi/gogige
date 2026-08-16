|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

# 1 Introduction

This document describes the Generic Data Container (GenDC) GenICam module. GenDC is Transport Layer neutral, self-described and used to represent, transmit or receive various kinds of data. GenDC targets especially machine vision related image data (such as 2D, 3D, multi-spectral) and metadata (like extra information, histograms and statistics). Besides the GenDC Container layout, this specification also describes the available data types.

A Transport Layer defines how to transport the GenDC Container with the concept of GenDC Flows but it does not know the content of the Container. This allows using and adding data types without touching a particular Transport Layer specification.

## 1.1 Objectives

The GenDC specification is intended to meet the following objectives:

1. Define a generic and self-described autonomous Data Container usable for representation, transmission and reception of arbitrary data Components.
2. Be Transport Layer agnostic, the Transport Layer is able to transport GenDC as an opaque Data Container without further knowledge.
3. Separate the notion of "What" is the data from the way "How" the data is generated or transported as far as feasible.
4. Be useable in hardware and software implementations.
5. Favor generality, future flexibility and expandability over Transport Layer or media specific definitions.
6. Be able to use the same Container layout anywhere in the data manipulation chain from the sensor data encoding to data delivery.
7. Define a GenDC Container that is also usable for general data storage.
8. Support complex and arbitrary data content (1D, 2D, 3D images, processing results, image sequences/bursts, multispectral, metadata, Mixed content ...).
9. Support heterogeneous and independent data Component's members.
10. Support multi-plane Components made of individual Parts of various and mixed size data format.
11. Support information metadata (such as GenICam chunk data).
12. Allow separate transfer and storage of Container Descriptor and data sections.
13. Permit reuse of GenDC Container encoder and decoder independent of the Transport Layer protocol.
14. Support simple and efficient implementation of encoding and decoding.
15. Permit the addition of new data types without requiring any Transport Layer protocol specification update.
16. Define a Container structure that supports early processing of data during the reception.