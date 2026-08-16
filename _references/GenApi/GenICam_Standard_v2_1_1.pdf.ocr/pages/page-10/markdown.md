|  GEN<ICAM |   | ![img-14.jpeg](img-14.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

“Example01” from vendor “Test” in this case). The other attributes are explained later in section 2.7.

Inside the <registerdescription> element, the nodes are lying in a flat order. Each node has a unique Name attribute and can be linked by sub-elements named pRole containing the Name of some other node.</registerdescription>

Each node has an optional <tooltip> element that contains a short description. The Gain node has additional elements that depend on its actual IntReg type and tells us, for example, the Address of the register or its Length. The default value is an empty string.</tooltip>

Typically, an implementation will create one software object per node and will link these objects together according to the logical links described in the XML file. \( ^{2} \)  The nodes can either be retrieved by their (unique) name or can be found by traversing the node graph starting with the root node. Once the user has a pointer to the node, he can access that feature through the node object's programming interface.

The syntax of the XML file is defined in the XML schema given by the schemaLocation-attribute. The schema is part of the standard. This document explains the ideas and overall structure of GenICam. The schema and its embedded reference documentation describe the formal details. In case of doubt, the schema's content overrides the content of this text.

The file location http://www.genicam.org/GenApi/GenApiSchema_Version_1_1.xsd is mandatory for the camera configuration file but can be overridden at runtime.

### 2.3 Nodes, Interfaces, and Abstract Features

Each node in the camera description file describes a single item. Depending on the item's nature, the node is of a specific node type and has a specific interface. The following interfaces are currently available \( ^{3} \) (each one is given with the typical widget used to map it on a graphical user interface):

- IInteger – maps to a slider with value, min, max, and increment
- IFloat – maps to a slider with value, min, and max plus a physical unit.
- IString – maps to an edit box showing a string
- IEnumeration – maps to a drop down box
- ICommand – maps to a command button
- IBoolean – maps to a check box
- IRegister – maps to an edit box showing a hex string
- ICategory – maps to an entry in a tree structuring the camera's features
- IPort – maps to the camera port and is typically not shown graphically