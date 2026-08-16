|  ![img-32.jpeg](img-32.jpeg) |   | ![img-33.jpeg](img-33.jpeg)  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Interface | IEnumeration  |
| --- | --- |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Group0 (if 0 based), Group1, Group2, ...  |

Selects a Group of component to control or inquire. The GroupSelector determines which components Group will be used for the selected features.

If only one group exists in the payload streamed by device (the common case), this selector can be omitted and all the components are considered member of the Group 0.

This Group notion is typically used by a transmitter to associate a subset of the Components member of a complex payload because they are related together.

An example could be a payload containing a group of Components related to the left view of an object and another group to the right view of that object.

Note that in order to simplify the standard text and feature descriptions, the optional GroupSelector is not explicitly propagated to all the features of the SFNC that it can potentially select.

4.15GroupIDValue

|  Name | GroupIDValue[GroupSelector]  |
| --- | --- |
|  Category | ImageFormatControl  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Returns a unique Identifier value corresponding to the selected Group of Components. If no grouping is required, this value should be set to 0.

This value is typically used by a transmitter to group a subset of the Components member of a complex payload when they are related together.

An example could be a payload containing a group of Components related to the left view of an object and another group to the right view of that object.