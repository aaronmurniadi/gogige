|  GENICAM |   | ![img-15.jpeg](img-15.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

The signature of the interfaces is given in more detail in section 2.9. The available node types are described in section 2.8. There might be multiple node types implementing the same interface type. The IInteger interface, for example, is (among others) implemented by the following node types:

- IntReg – extracts an integer lying byte-bounded in a register
- MaskedIntReg – extracts an integer packed into a register, e.g., from bit 8 to bit 12
- Integer – merges the integer's value, min, max, and increment properties from different nodes

Each node type extracts an integer from different sources in a different way. The output of all of these nodes, however, can be used as type-safe input for all links where an integer is required.

Abstract features are always described in terms of an interface type, a name, and a meaning. For example, the Gain (name) of a camera might be defined as an Integer (interface type) and might describe the amplification inside a camera (meaning). Note that other possible definitions exist, e.g., the Gain could be defined as an IEnumeration or as an IFloat.

### 2.4 Getting and Setting Values

When the user reads or writes the value of a node, this node will trigger a cascade of read and write operations within the node graph. To illustrate this, Figure 4 shows a more elaborate example for the Gain feature. The Gain feature is exposed via an Integer interface that lets the user get and set the feature's Value and lets her read (among other things) the feature's Min and Max value. The example in Figure 4 assumes that the camera has three registers, one for the Gain Value itself, one for its Min value, and one for its Max value. From each of these registers, the corresponding value is extracted using an IntReg node. The Integer node with the name Gain then collects the data and merges them, exposing the results with an Integer interface.