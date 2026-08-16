|  GENICAM |   | ![img-34.jpeg](img-34.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

Note that compatibility refers only to the feature nodes and the underlying register layout. However it does not refer to implementation nodes (for details see section 2.8.2).

#### 2.7.3 Identifying and Caching the Camera Description File

Loading a camera description file may involve one or more pre-processing steps. To speed things up, the pre-processed XML file can be cached. For caching, a key is required that uniquely identifies the camera description file. A combination of the <RegisterDescription> element's VendorName, ModelName, MajorVersion, MinorVersion, and SubMinorVersion attributes would be sufficient, but is a bit clumsy to use.</RegisterDescription>

The VersionGuid and ProductGuid are no longer used; they are kept in the schema for backward compatibility. GenICam description file must however have a valid GUID in those fields. Note that for caching purposes a hash over the XML file's content should be used.

### 2.8 Available Node Types

This section gives a brief description of each available node type, of their behavior, usage, and most interesting parameters. In addition, there is a formal description for the XML layout of each node in an XML schema file included with the GenICam standard. This schema file can be read by most XML editors and will greatly simplify creating camera description files by providing a syntax check and context sensitive fill-in helpers.

This document refers to the GenApi schema version 1.1 found in the file GenApiSchema_Version_1_1.xsd. Note that in subsequent versions of the standard, additional node types, elements, and attributes may be added, however, backward compatibility will be maintained if at all possible.

Some node types have elements or attributes which are only an optional part of a valid XML device description file. The default values for this elements and attributes are listed in the Appendix “Default values of optional Node elements and attributes” (section 3.2).

#### 2.8.1 Node

The Node type contains those elements and attributes common to all other node types. A stand-alone Node node is pretty useless, but is possible for testing purposes. Here is an example: