|  GEN<ICAM |   | ![img-36.jpeg](img-36.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

○ If the node from B has MergePriority = -1 the corresponding node from A it is copied to C.
○ If the node from B has MergePriority = 0 (default) or the MergeAttribute is missing an error occurs.

- If a Category node is present in A and B the Category node from the target file A is copied to C with the <pFeature> entries from file B added while avoiding duplicates.</pFeature>

A Node can have an ExposeStatic attribute which can have the values Yes or No to be missing. It controls which node is exposed in the static use case to the customer. The following rules apply:

- Features are exposed if they don't have ExposeStatic=No set
• Non-Features are exposed only if they have ExposeStatic=Yes set

An <Extension> element can be used to add custom specific data to a camera description file. All elements placed inside the <Extension> element are ignored.</Extension>

The <ToolTip> element gives a short description of the node. It may also be used as a brief description for reference documentation automatically generated from the camera description file.</ToolTip>

The <Description> element gives a more detailed description of the node. It may also be used as a long description for reference documentation automatically generated from the camera description file.</Description>

The <DisplayName> element lets you define feature captions that might be used instead of the feature's Name.</DisplayName>

The <Visibility> element defines the user level that should get access to the feature. Possible values are: Beginner, Expert, Guru, and Invisible. The latter is required to make a feature show up in the API, but not in the GUI (see section 2.8.2). </Visibility>

The <EventID> element is used for delivering asynchronous events. A camera might send an event package to indicate that one or more data item in the camera has changed its value. GenICam handles the event by invalidating the nodes corresponding to the data items. The nodes are found by the EventID which is a hexadecimal number which comes with the event package from the camera. Each node can have one (optional) EventID element.</EventID>

The <pIsImplemented>, <pIsAvailable>, and <pIsLocked> elements contains the names of nodes implementing an Integer interface. If these elements are present, they influence the access mode of this node as described in section 2.5.</pIsLocked>

The <pBlockPolling> element bocks the polling on a node with a PollingTime entry if the target of the element is !=0.</pBlockPolling>

An <ImposedAccessMode> element can be used to narrow the access mode resulting from other nodes.</ImposedAccessMode>

<pAlias> points to another node which describes the same feature in a different manner. This feature will be mainly used in a GUI: a Category might be replaced by its alias if not all members are shown; an integer and a flat node might be aliases of each other if they show the raw and the abs value of a feature.</pAlias>