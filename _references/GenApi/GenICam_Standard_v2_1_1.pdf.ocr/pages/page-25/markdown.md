|  GEN<ICAM |   | ![img-37.jpeg](img-37.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

points to another node which describes the same feature in a way that the original node and the cast alias node can be casted into each other.

<isdeprecated> denotes that the corresponding feature is deprecated and should not be used for new designs anymore.</isdeprecated>

<streamable> denotes that the corresponding feature is prepared to be stored to and loaded from a file via the GenApi node tree. The idea is to persist the state of a camera by storing the features marked as Streamable and restore the state by writing those features back to the node tree.</streamable>

<pError> points to an enumeration which is checked after setting the value of a node. The enumeration must have one entry with IntValue 0 which indicates no error. If another value is set an exception is thrown with the DisplayName and the ToolTip of the EnumEntry as error message.</pError>

<DocuURL> Provides a http URL pointing to a location were documentation for the node can be found. The notation can contain Variables in the form $(NAME) were NAME is either a Node name or one of the following special names:</DocuURL>

- Sys::NodeName : Name of the current node
- Sys::ModelName: content of the XML file's ModelName attribute
- Sys::VendorName: content of the XML file's VendorName attribute
- Sys::StandardNamespace: content of the XML file's StandardNamespace attribute
- Sys::GenApiVersion : version of the GenApi software (<major>.<minor>.<subminor>)
- Sys::DeviceVersion : version of the device (<major>.<minor>.<subminor>)
- Sys::SchemaVersion : version of the schema (<major>.<minor>.<subminor>)
- Sys::Application: Name of the executable file
- Sys::OperatingSystem: Name of the operating system. Format "Windows5.1_SP3.0"
- Sys::Language: Name of the operating system's locale ID. Format "German"

#### 2.8.2 Category

The Category node is used to group features that should be presented to the user. It implements the ICategory interface and inherits all Node elements. It also contains a list of  elements that point to the features contained in the category. Categories can contain other categories, thus forming a tree of arbitrary depth.

There is one special Category node with the standard name  \( Root^{6} \)  that is the basis of the category tree. Users may want to start browsing the features of a camera from here. The following example creates the node graph shown in Figure 12: