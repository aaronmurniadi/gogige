|  **GENICAM** |   | emva  |
| --- | --- | --- |
|  Version 2.4 | Pixel Format Naming Convention  |   |

# 1 Introduction

## 1.1 Purpose

The intention of this document is to define a generic convention to name the pixel formats used in machine vision. This covers 2D images as well as 3D imaging data. The aim is not to provide a unique definition for all theoretical possibilities, but to provide clear guidelines to follow when a new pixel format is introduced. As such, the pixel format designation is not sufficient to deduce all the pixel characteristics (that would be next to impossible anyway with the number of possible permutations!), but following those guidelines should provide a uniform way to name new pixel types so they fit well within the current set, even though the layout of each specific pixel format might need to be explicitly illustrated. When this convention is not sufficient, a camera interface-specific designator can be appended to remove any ambiguity.

Note: The main objective is to have clear guidelines in how to designate pixel format: a text string associated to a pixel format. The actual numerical value associated to each pixel format, the GenICam display name and the way pixel information is put into data packets is beyond the scope of this document.

This document covers the traditional 2D images, but starting with version 2.0 it also introduces support for 3D imaging data. For 3D, the formats are proposed as “abstract” with no defined mapping to actual real-world units and coordinate systems (such as Cartesian or spherical) or its properties (such as orientation). Such mapping should be defined by other means, in particular through the GenICam Standard Feature Naming Convention (SFNC) device description file. Letters A, B and C are used for the abstract coordinate names, where A-B-C can mean X-Y-Z for the Cartesian coordinate system, Theta-Phi-Rho for the spherical system, etc. For so-called 2.5D, the C always stands for the “depth/range” coordinate that can also be transferred standalone.

The Pixel Format Naming Convention supplements the GenICam Standard Feature Naming Convention (SFNC). As such, it is a child document of the SFNC. Request for clarifications or to add pixel formats not supported by the current syntax should be directed at the current editor of this document, as listed in the document history section.

Important: The PFNC defines pixels names and formats in an interoperable way so they can be shared across technologies. A given imaging standard might use internally a different low level bit encoding while keeping the PFNC name. This is acceptable as long as the low level bit encoding respects PFNC when it exits the boundary of that imaging standard. In that case, the imaging standard is required to define its actual bit coding.