|  GENICAM |   | ![img-35.jpeg](img-35.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

<Node Name="Gain" NameSpace="Standard">
<Extension>
    <MyElement>Something vendor specific</MyElement>
</Extension>
<toolTip>The amplification of the camera</ToolTip>
<Description>A more elaborated description</Description>
<DisplayName>Gain</DisplayName>
<Visibility>Expert</Visibility>
<EventID>12fc</EventID>
<pIsImplemented>SomeNode1</pIsImplemented>
<pIsAvailable>SomeNode2</pIsAvailable>
<pIsLocked>SomeNode3</pIsLocked>
<pError>NodeIndicatingAnError</pError>
<ImposedAccessMode>RO</ImposedAccessMode>
<pAlias>SomeNode4</pAlias>

Each node has a Name attribute. The Name must be unique within the camera description file. Names can be composed of alphanumeric characters [A-Za-z0-9]. The schema also allows the use of the underscore ‘_’, but not as a leading character. This is because the reference implementation uses a leading underscore for implementation related names.

Each Name lives inside a name space which is identified by the NameSpace attribute of the node which can have two possible values: Custom or Standard. If it is Custom, any name can be used as long as it is unique within the camera description file.

If it is Standard, it must come from the standard feature name lists (SFNC). The SFNC is mostly interface agnostic and contains only very few interface specific features. In order to advertise which interface is used the floating node DeviceTLType should be implemented (for details see SFNC). Typical entries are

- IIDC : cameras following the 1394 IIDC standard (also called DCAM standard)
• GEV : cameras following the GigE Vision standard
- CL : cameras following the Camera Link standard
• U3V : cameras following the USB3 Vision standard
• CXP : cameras following the CoaXPRESS standard
- CLHS : cameras following the Camera Link HS standard
• None : no standard is used

The StandardNameSpace attribute of the enclosing <RegisterDescription> element (see section 2.7) is deprecated since it does not reflect new interface types like, e.g. CXP.

A Node can have a MergePriority attribute which can have the values +1, 0, or -1. It controls the way two XML files A and B are merged to a target file C. A is called the target file and B is called the inject file. Nodes are compared based on their Name attribute only.

- If a node is present only in A or B it is copied to C
- If a non-Category node is present in A and B the following rules apply (note that the MergePriority attribute of the target file A is ignored):

○ If the node from the injected file B has MergePriority = +1 it is copied to C.